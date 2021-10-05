package replication

import (
	"context"
	"fmt"

	"github.com/influxdata/influx-cli/v2/api"
	"github.com/influxdata/influx-cli/v2/clients"
)

type Client struct {
	clients.CLI
	api.ReplicationsApi
	api.OrganizationsApi
}

type CreateParams struct {
	Name           string
	Description    string
	OrgID          string
	OrgName        string
	RemoteID       string
	LocalBucketID  string
	RemoteBucketID string
	MaxQueueSize   int64
}

func (c Client) Create(ctx context.Context, params *CreateParams) error {
	orgID, err := c.GetOrgId(ctx, params.OrgID, params.OrgName, c.OrganizationsApi)
	if err != nil {
		return err
	}

	// set up a struct with required params
	body := api.ReplicationCreationRequest{
		Name:              params.Name,
		OrgID:             orgID,
		RemoteID:          params.RemoteID,
		LocalBucketID:     params.LocalBucketID,
		RemoteBucketID:    params.RemoteBucketID,
		MaxQueueSizeBytes: params.MaxQueueSize,
	}

	// set optional params if specified
	if params.Description != "" {
		body.Description = &params.Description
	}

	// send post request
	res, err := c.PostReplication(ctx).ReplicationCreationRequest(body).Execute()
	if err != nil {
		return fmt.Errorf("failed to create replication stream %q: %w", params.Name, err)
	}

	// print confirmation of new replication stream
	return c.printReplication(printReplicationOpts{replication: &res})
}

type printReplicationOpts struct {
	replication  *api.Replication
	replications []api.Replication
	deleted      bool
}

func (c Client) printReplication(opts printReplicationOpts) error {
	if c.PrintAsJSON {
		var v interface{}
		if opts.replication != nil {
			v = opts.replication
		} else {
			v = opts.replications
		}
		return c.PrintJSON(v)
	}

	headers := []string{"ID", "Name", "Description", "Org ID", "Remote Connection ID", "Local Bucket", "Remote Bucket", "Max Queue Size Bytes"}
	if opts.deleted {
		headers = append(headers, "Deleted")
	}

	if opts.replication != nil {
		opts.replications = append(opts.replications, *opts.replication)
	}

	var rows []map[string]interface{}
	for _, r := range opts.replications {
		row := map[string]interface{}{
			"ID":                   r.GetId(),
			"Name":                 r.GetName(),
			"Description":          r.GetDescription(),
			"Org ID":               r.GetOrgID(),
			"Remote Connection ID": r.GetRemoteID(),
			"Local Bucket":         r.GetLocalBucketID(),
			"Remote Bucket":        r.GetRemoteBucketID(),
			"Max Queue Size Bytes": r.GetMaxQueueSizeBytes(),
		}
		if opts.deleted {
			row["Deleted"] = true
		}
		rows = append(rows, row)
	}

	return c.PrintTable(headers, rows...)
}
