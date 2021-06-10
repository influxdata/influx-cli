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
)

type Client struct {
	clients.CLI
	api.BackupApi

	// Local state tracked across steps in the backup process.
	baseName       string
	bucketMetadata []api.BucketMetadataManifest
	manifest       Manifest
}

type Params struct {
	// Organization containing TSM data to back up.
	// If not set, all orgs will be included.
	OrgID string
	Org   string

	// Bucket containing TSM data to back up.
	// If not set, all buckets within the org filter will be included.
	BucketID string
	Bucket   string

	// Path to the directory where backup files should be written.
	Path string

	// Compression to use for local copies of snapshot files.
	Compression FileCompression
}

func (p *Params) matches(bkt api.BucketMetadataManifest) bool {
	if p.OrgID != "" && bkt.OrganizationID != p.OrgID {
		return false
	}
	if p.Org != "" && bkt.OrganizationName != p.Org {
		return false
	}
	if p.BucketID != "" && bkt.BucketID != p.BucketID {
		return false
	}
	if p.Bucket != "" && bkt.BucketName != p.Bucket {
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

	if err := c.downloadMetadata(ctx, params); err != nil {
		return fmt.Errorf("failed to backup metadata: %w", err)
	}
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

	writeFile := func(from io.Reader, to string) (ManifestFileEntry, error) {
		toPath := filepath.Join(params.Path, to)
		if params.Compression == GzipCompression {
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
			if params.Compression == GzipCompression {
				gw := gzip.NewWriter(out)
				defer gw.Close()
				outW = gw
			}

			_, err = io.Copy(outW, from)
			return err
		}(); err != nil {
			return ManifestFileEntry{}, err
		}

		fi, err := os.Stat(toPath)
		if err != nil {
			return ManifestFileEntry{}, err
		}
		return ManifestFileEntry{
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
			c.manifest.SQL = fi
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

// downloadBucketData downloads TSM snapshots for each shard in the buckets matching
// the filter parameters provided over the CLI. Snapshots are written to local files.
//
// Bucket metadata must be pre-seeded via downloadMetadata before this method is called.
func (c *Client) downloadBucketData(ctx context.Context, params *Params) error {
	c.manifest.Buckets = make([]ManifestBucketEntry, 0, len(c.bucketMetadata))
	for _, b := range c.bucketMetadata {
		if !params.matches(b) {
			continue
		}
		bktManifest, err := ConvertBucketManifest(b, func(shardId int64) (*ManifestFileEntry, error) {
			log.Printf("INFO: Backing up TSM for shard %d", shardId)
			res, err := c.GetBackupShardId(ctx, shardId).AcceptEncoding("gzip").Execute()
			if err != nil {
				if oapiErr, ok := err.(*api.GenericOpenAPIError); ok {
					if oapiErr.Model() != nil && oapiErr.Model().ErrorCode() == api.ERRORCODE_NOT_FOUND {
						log.Printf("WARN: Shard %d removed during backup", shardId)
						return nil, nil
					}
				}
				return nil, err
			}
			defer res.Body.Close()

			fileName := fmt.Sprintf("%s.%d.tar", c.baseName, shardId)
			if params.Compression == GzipCompression {
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
				if params.Compression == GzipCompression && res.Header.Get("Content-Encoding") != "gzip" {
					gzw := gzip.NewWriter(outW)
					defer gzw.Close()
					outW = gzw
				}
				if params.Compression == NoCompression && res.Header.Get("Content-Encoding") == "gzip" {
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
			return &ManifestFileEntry{
				FileName:    fi.Name(),
				Size:        fi.Size(),
				Compression: params.Compression,
			}, nil
		})
		if err != nil {
			return err
		}
		c.manifest.Buckets = append(c.manifest.Buckets, bktManifest)
	}
	return nil
}

// writeManifest writes a description of all files downloaded as part of the backup process
// to the backup folder, encoded as JSON.
func (c Client) writeManifest(params *Params) error {
	manifestPath := filepath.Join(params.Path, fmt.Sprintf("%s.manifest", c.baseName))
	f, err := os.OpenFile(manifestPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return err
	}
	defer f.Close()

	enc := json.NewEncoder(f)
	enc.SetIndent("", "  ")
	return enc.Encode(c.manifest)
}
