package backup

import (
	"compress/gzip"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"

	"github.com/influxdata/influx-cli/v2/api"
	"github.com/influxdata/influx-cli/v2/clients"
	br "github.com/influxdata/influx-cli/v2/internal/backup_restore"
)

type Client struct {
	clients.CLI
	api.HealthApi
	api.BackupApi

	// Local state tracked across steps in the backup process.
	baseName       string
	bucketMetadata []api.BucketMetadataManifest
	manifest       br.Manifest
}

type Params struct {
	// Organization containing TSM data to back up.
	// If not set, all orgs will be included.
	clients.OrgParams

	// Bucket containing TSM data to back up.
	// If not set, all buckets within the org filter will be included.
	clients.BucketParams

	// Path to the directory where backup files should be written.
	Path string

	// Compression to use for local copies of snapshot files.
	Compression br.FileCompression

	// Amount of Workers to create to copy bucket data in parallel.
	Workers int
}

func (p *Params) matches(bkt api.BucketMetadataManifest) bool {
	if p.OrgID != "" && bkt.OrganizationID != p.OrgID {
		return false
	}
	if p.OrgName != "" && bkt.OrganizationName != p.OrgName {
		return false
	}
	if p.BucketID != "" && bkt.BucketID != p.BucketID {
		return false
	}
	if p.BucketName != "" && bkt.BucketName != p.BucketName {
		return false
	}
	return true
}

const backupFilenamePattern = "20060102T150405Z"

func (c *Client) Backup(ctx context.Context, params *Params) error {
	if err := os.MkdirAll(params.Path, 0777); err != nil {
		return err
	}
	c.baseName = time.Now().UTC().Format(backupFilenamePattern)
	c.manifest.Version = br.ManifestVersion

	// The APIs we use to back up metadata depends on the server's version.
	legacyServer, err := br.ServerIsLegacy(ctx, c.HealthApi)
	if err != nil {
		return err
	}
	backupMetadata := c.downloadMetadata
	if legacyServer {
		backupMetadata = c.downloadMetadataLegacy
	}
	if err := backupMetadata(ctx, params); err != nil {
		return fmt.Errorf("failed to backup metadata: %w", err)
	}

	// Once metadata has been fetched, things are consistent across versions.
	if err := c.downloadBucketData(ctx, params); err != nil {
		return fmt.Errorf("failed to backup bucket data: %w", err)
	}
	if err := c.writeManifest(params); err != nil {
		return fmt.Errorf("failed to write backup manifest: %w", err)
	}
	return nil
}

// downloadMetadata downloads a snapshot of the KV store, SQL DB, and bucket
// manifests from the server. KV and SQL are written to local files. Bucket manifests
// are parsed into a slice for additional processing.
func (c *Client) downloadMetadata(ctx context.Context, params *Params) error {
	log.Println("INFO: Downloading metadata snapshot")
	rawResp, err := c.GetBackupMetadata(ctx).AcceptEncoding("gzip").Execute()
	if err != nil {
		return fmt.Errorf("failed to download metadata snapshot: %w", err)
	}
	defer rawResp.Body.Close()

	kvName := fmt.Sprintf("%s.bolt", c.baseName)
	sqlName := fmt.Sprintf("%s.sqlite", c.baseName)

	_, contentParams, err := mime.ParseMediaType(rawResp.Header.Get("Content-Type"))
	if err != nil {
		return err
	}
	body, err := api.GunzipIfNeeded(rawResp)
	if err != nil {
		return err
	}
	defer body.Close()

	writeFile := func(from io.Reader, to string) (br.ManifestFileEntry, error) {
		toPath := filepath.Join(params.Path, to)
		if params.Compression == br.GzipCompression {
			toPath = toPath + ".gz"
		}

		// Closure here so we can clean up file resources via `defer` without
		// returning from the whole function.
		if err := func() error {
			out, err := os.OpenFile(toPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
			if err != nil {
				return err
			}
			defer out.Close()

			var outW io.Writer = out
			if params.Compression == br.GzipCompression {
				gw := gzip.NewWriter(out)
				defer gw.Close()
				outW = gw
			}

			_, err = io.Copy(outW, from)
			return err
		}(); err != nil {
			return br.ManifestFileEntry{}, err
		}

		fi, err := os.Stat(toPath)
		if err != nil {
			return br.ManifestFileEntry{}, err
		}
		return br.ManifestFileEntry{
			FileName:    fi.Name(),
			Size:        fi.Size(),
			Compression: params.Compression,
		}, nil
	}

	mr := multipart.NewReader(body, contentParams["boundary"])
	for {
		part, err := mr.NextPart()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}

		_, partParams, err := mime.ParseMediaType(part.Header.Get("Content-Disposition"))
		if err != nil {
			return err
		}
		switch name := partParams["name"]; name {
		case "kv":
			fi, err := writeFile(part, kvName)
			if err != nil {
				return fmt.Errorf("failed to save local copy of KV backup to %q: %w", kvName, err)
			}
			c.manifest.KV = fi
		case "sql":
			fi, err := writeFile(part, sqlName)
			if err != nil {
				return fmt.Errorf("failed to save local copy of SQL backup to %q: %w", sqlName, err)
			}
			c.manifest.SQL = &fi
		case "buckets":
			if err := json.NewDecoder(part).Decode(&c.bucketMetadata); err != nil {
				return fmt.Errorf("failed to decode bucket manifest from backup: %w", err)
			}
		default:
			return fmt.Errorf("response contained unexpected part %q", name)
		}
	}
	return nil
}

// downloadMetadataLegacy downloads a snapshot of the KV store from the server, and extracts
// a bucket manifest from the result. KV is written to a local file; the extracted manifest
// is tracked as a slice for additional processing.
//
// NOTE: This should _not_ be used against an InfluxDB instance running v2.1.0 or later, as
// it will fail to capture metadata stored in SQL.
func (c *Client) downloadMetadataLegacy(ctx context.Context, params *Params) error {
	log.Println("INFO: Downloading legacy KV snapshot")
	rawResp, err := c.GetBackupKV(ctx).Execute()
	if err != nil {
		return fmt.Errorf("failed to download KV snapshot: %w", err)
	}
	defer rawResp.Body.Close()

	kvName := filepath.Join(params.Path, fmt.Sprintf("%s.bolt", c.baseName))
	tmpKv := fmt.Sprintf("%s.tmp", kvName)
	defer os.RemoveAll(tmpKv)

	// Since we need to read the bolt DB to extract a manifest, always save it uncompressed locally.
	if err := func() error {
		f, err := os.Create(tmpKv)
		if err != nil {
			return err
		}
		defer f.Close()

		_, err = io.Copy(f, rawResp.Body)
		return err
	}(); err != nil {
		return fmt.Errorf("failed to save downloaded KV snapshot: %w", err)
	}

	// Extract the metadata we need from the downloaded KV store, and convert it to a new-style manifest.
	log.Println("INFO: Extracting bucket manifest from legacy KV snapshot")
	c.bucketMetadata, err = br.ExtractBucketMetadata(tmpKv)
	if err != nil {
		return fmt.Errorf("failed to extract bucket metadata from downloaded KV snapshot: %w", err)
	}

	// Move/compress the bolt DB into its final location.
	if err := func() error {
		if params.Compression == br.NoCompression {
			return os.Rename(tmpKv, kvName)
		}

		tmpIn, err := os.Open(tmpKv)
		if err != nil {
			return err
		}
		defer tmpIn.Close()

		kvName = kvName + ".gz"
		out, err := os.Create(kvName)
		if err != nil {
			return err
		}
		defer out.Close()

		gzw := gzip.NewWriter(out)
		defer gzw.Close()

		_, err = io.Copy(gzw, tmpIn)
		return err
	}(); err != nil {
		return fmt.Errorf("failed to rename downloaded KV snapshot: %w", err)
	}

	fi, err := os.Stat(kvName)
	if err != nil {
		return fmt.Errorf("failed to inspect local KV snapshot: %w", err)
	}
	c.manifest.KV = br.ManifestFileEntry{
		FileName:    fi.Name(),
		Size:        fi.Size(),
		Compression: params.Compression,
	}
	return nil
}

// downloadBucketData downloads TSM snapshots for each shard in the buckets matching
// the filter parameters provided over the CLI. Snapshots are written to local files.
//
// Bucket metadata must be pre-seeded via downloadMetadata before this method is called.
func (c *Client) downloadBucketData(ctx context.Context, params *Params) error {
	c.manifest.Buckets = make([]br.ManifestBucketEntry, 0, len(c.bucketMetadata))
	for _, b := range c.bucketMetadata {
		if !params.matches(b) {
			continue
		}
		bktManifest, err := br.ConvertBucketManifest(b, params.Workers, func(shardId int64) (*br.ManifestFileEntry, error) {
			return c.downloadShardData(ctx, params, shardId)
		})
		if err != nil {
			return err
		}
		c.manifest.Buckets = append(c.manifest.Buckets, bktManifest)
	}
	return nil
}

// downloadShardData downloads the TSM snapshot for a single shard. The snapshot is written
// to a local file, and its metadata is returned for aggregation.
func (c Client) downloadShardData(ctx context.Context, params *Params, shardId int64) (*br.ManifestFileEntry, error) {
	log.Printf("INFO: Backing up TSM for shard %d", shardId)
	res, err := c.GetBackupShardId(ctx, shardId).AcceptEncoding("gzip").Execute()
	if err != nil {
		if apiError, ok := err.(api.ApiError); ok {
			if apiError.ErrorCode() == api.ERRORCODE_NOT_FOUND {
				log.Printf("WARN: Shard %d removed during backup", shardId)
				return nil, nil
			}
		}
		return nil, err
	}
	defer res.Body.Close()

	fileName := fmt.Sprintf("%s.%d.tar", c.baseName, shardId)
	if params.Compression == br.GzipCompression {
		fileName = fileName + ".gz"
	}
	path := filepath.Join(params.Path, fileName)

	// Closure here so we can clean up file resources via `defer` without
	// returning from the whole function.
	if err := func() error {
		f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
		if err != nil {
			return err
		}
		defer f.Close()

		var inR io.Reader = res.Body
		var outW io.Writer = f

		// Make sure the locally-written data is compressed according to the user's request.
		if params.Compression == br.GzipCompression && res.Header.Get("Content-Encoding") != "gzip" {
			gzw := gzip.NewWriter(outW)
			defer gzw.Close()
			outW = gzw
		}
		if params.Compression == br.NoCompression && res.Header.Get("Content-Encoding") == "gzip" {
			gzr, err := gzip.NewReader(inR)
			if err != nil {
				return err
			}
			defer gzr.Close()
			inR = gzr
		}

		_, err = io.Copy(outW, inR)
		return err
	}(); err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		return nil, err
	}
	return &br.ManifestFileEntry{
		FileName:    fi.Name(),
		Size:        fi.Size(),
		Compression: params.Compression,
	}, nil
}

// writeManifest writes a description of all files downloaded as part of the backup process
// to the backup folder, encoded as JSON.
func (c Client) writeManifest(params *Params) error {
	manifestPath := filepath.Join(params.Path, fmt.Sprintf("%s.%s", c.baseName, br.ManifestExtension))
	f, err := os.OpenFile(manifestPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return err
	}
	defer f.Close()

	enc := json.NewEncoder(f)
	enc.SetIndent("", "  ")
	return enc.Encode(c.manifest)
}
