package replication

import (
	"context"
	"errors"
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

type ListParams struct {
	Name          string
	OrgID         string
	OrgName       string
	RemoteID      string
	LocalBucketID string
}

func (c Client) List(ctx context.Context, params *ListParams) error {

	orgID, err := c.GetOrgId(ctx, params.OrgID, params.OrgName, c.OrganizationsApi)
	if err != nil {
		return err
	}

	// set up params
	req := c.GetReplications(ctx).OrgID(orgID)

	if params.Name != "" {
		req = req.Name(params.Name)
	}

	if params.RemoteID != "" {
		req = req.RemoteID(params.RemoteID)
	}

	if params.LocalBucketID != "" {
		req = req.LocalBucketID(params.LocalBucketID)
	}

	// send get request
	res, err := req.Execute()
	if err != nil {
		return fmt.Errorf("failed to get replication streams: %w", err)
	}

	// print replication stream info
	printOpts := printReplicationOpts{}
	if res.Replications != nil {
		printOpts.replications = *res.Replications
	} else {
		return errors.New("no replication streams found for specified parameters")
	}

	return c.printReplication(printOpts)
}

type UpdateParams struct {
	ReplicationID  string
	Name           string
	Description    string
	RemoteID       string
	RemoteBucketID string
	MaxQueueSize   int64
}

func (c Client) Update(ctx context.Context, params *UpdateParams) error {
	// build request
	body := api.ReplicationUpdateRequest{}

	if params.Name != "" {
		body.SetName(params.Name)
	}

	if params.Description != "" {
		body.SetDescription(params.Description)
	}

	if params.RemoteID != "" {
		body.SetRemoteID(params.RemoteID)
	}

	if params.RemoteBucketID != "" {
		body.SetRemoteBucketID(params.RemoteBucketID)
	}

	if params.MaxQueueSize != 0 {
		body.SetMaxQueueSizeBytes(params.MaxQueueSize)
	}

	// send patch request
	res, err := c.PatchReplicationByID(ctx, params.ReplicationID).ReplicationUpdateRequest(body).Execute()
	if err != nil {
		return fmt.Errorf("failed to update replication stream %q: %w", params.ReplicationID, err)
	}
	// print updated replication stream info
	return c.printReplication(printReplicationOpts{replication: &res})
}

func (c Client) Delete(ctx context.Context, replicationID string) error {
	// get replication stream via ID
	connection, err := c.GetReplicationByID(ctx, replicationID).Execute()
	if err != nil {
		return fmt.Errorf("could not find replication stream with ID %q: %w", replicationID, err)
	}

	// send delete request
	if err := c.DeleteReplicationByID(ctx, replicationID).Execute(); err != nil {
		return fmt.Errorf("failed to delete replication stream %q: %w", replicationID, err)
	}

	// print deleted replication stream info
	printOpts := printReplicationOpts{
		replication: &connection,
		deleted:     true,
	}

	return c.printReplication(printOpts)
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

	headers := []string{"ID", "Name", "Org ID", "Remote ID", "Local Bucket ID", "Remote Bucket ID",
		"Current Queue Bytes", "Max Queue Bytes", "Latest Status Code"}
	if opts.deleted {
		headers = append(headers, "Deleted")
	}

	if opts.replication != nil {
		opts.replications = append(opts.replications, *opts.replication)
	}

	var rows []map[string]interface{}
	for _, r := range opts.replications {
		row := map[string]interface{}{
			"ID":                  r.GetId(),
			"Name":                r.GetName(),
			"Org ID":              r.GetOrgID(),
			"Remote ID":           r.GetRemoteID(),
			"Local Bucket ID":     r.GetLocalBucketID(),
			"Remote Bucket ID":    r.GetRemoteBucketID(),
			"Current Queue Bytes": r.GetCurrentQueueSizeBytes(),
			"Max Queue Bytes":     r.GetMaxQueueSizeBytes(),
			"Latest Status Code":  r.GetLatestResponseCode(),
		}
		if opts.deleted {
			row["Deleted"] = true
		}
		rows = append(rows, row)
	}

	return c.PrintTable(headers, rows...)
}
