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

func (p *Params) Matches(bkt api.BucketMetadataManifest) bool {
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
	baseName := time.Now().UTC().Format(backupFilenamePattern)

	log.Println("INFO: Downloading metadata snapshot")
	rawResp, err := c.GetBackupMetadata(ctx).AcceptEncoding("gzip").Execute()
	if err != nil {
		return fmt.Errorf("failed to download metadata snapshot: %w", err)
	}

	kvName := fmt.Sprintf("%s.bolt", baseName)
	sqlName := fmt.Sprintf("%s.sqlite", baseName)

	writeFile := func(from io.Reader, to string) (os.FileInfo, error) {
		toPath := filepath.Join(params.Path, to)
		if params.Compression == GzipCompression {
			toPath = toPath + ".gz"
		}

		out, err := os.OpenFile(toPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
		if err != nil {
			return nil, err
		}
		defer out.Close()

		var outW io.Writer = out
		flush := func() {}
		if params.Compression == GzipCompression {
			gw := gzip.NewWriter(out)
			outW = gw
			flush = func() {
				gw.Close()
			}
		}
		_, err = io.Copy(outW, from)
		flush()
		if err != nil {
			return nil, err
		}
		return out.Stat()
	}

	_, contentParams, err := mime.ParseMediaType(rawResp.Header.Get("Content-Type"))
	if err != nil {
		return err
	}
	body, err := api.GunzipIfNeeded(rawResp)
	if err != nil {
		return err
	}
	defer body.Close()

	var m Manifest
	var buckets []api.BucketMetadataManifest
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
			m.KV = ManifestFileEntry{
				FileName: kvName,
				Size:     fi.Size(),
			}
		case "sql":
			fi, err := writeFile(part, sqlName)
			if err != nil {
				return fmt.Errorf("failed to save local copy of SQL backup to %q: %w", sqlName, err)
			}
			m.SQL = ManifestFileEntry{
				FileName: sqlName,
				Size:     fi.Size(),
			}
		case "buckets":
			if err := json.NewDecoder(part).Decode(&buckets); err != nil {
				return fmt.Errorf("failed to decode bucket manifest from backup: %w", err)
			}
		default:
			return fmt.Errorf("response contained unexpected part %q", name)
		}
	}

	m.Buckets = make([]ManifestBucketEntry, 0, len(buckets))
	for _, b := range buckets {
		if !params.Matches(b) {
			continue
		}
		bktManifest, err := ConvertBucketManifest(b, func(shardId int64) (os.FileInfo, error) {
			log.Printf("INFO: Backing up TSM for shard %d", shardId)
			res, err := c.GetBackupShardId(ctx, shardId).AcceptEncoding("gzip").Execute()
			if err != nil {
				if oapiErr, ok := err.(*api.GenericOpenAPIError); ok {
					if coreErr, ok := oapiErr.Model().(*api.Error); ok && coreErr.Code == api.ERRORCODE_NOT_FOUND {
						log.Printf("WARN: Shard %d removed during backup", shardId)
						return nil, nil
					}
				}
				return nil, err
			}
			defer res.Body.Close()

			fileName := fmt.Sprintf("%s.%d.tar", baseName, shardId)
			if params.Compression == GzipCompression {
				fileName = fileName + ".gz"
			}

			path := filepath.Join(params.Path, fileName)
			f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
			if err != nil {
				return nil, err
			}
			defer f.Close()

			var inR io.Reader = res.Body
			var outW io.Writer = f
			flush := func() {}

			if params.Compression == GzipCompression && res.Header.Get("Content-Encoding") != "gzip" {
				gzw := gzip.NewWriter(outW)
				flush = func() {
					gzw.Close()
				}
				outW = gzw
			}
			if params.Compression == NoCompression && res.Header.Get("Content-Encoding") == "gzip" {
				gzr, err := gzip.NewReader(inR)
				if err != nil {
					return nil, err
				}
				defer gzr.Close()
				inR = gzr
			}

			_, err = io.Copy(outW, inR)
			flush()
			if err != nil {
				return nil, err
			}

			return f.Stat()
		})
		if err != nil {
			return err
		}
		m.Buckets = append(m.Buckets, bktManifest)
	}

	manifestPath := filepath.Join(params.Path, fmt.Sprintf("%s.manifest", baseName))
	buf, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		return err
	}
	buf = append(buf, '\n')
	if err := os.WriteFile(manifestPath, buf, 0600); err != nil {
		return fmt.Errorf("failed to write manifest to %q: %w", manifestPath, err)
	}
	return nil
}
