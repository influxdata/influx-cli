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
	clients.OrgParams
	Name                   string
	Description            string
	RemoteID               string
	LocalBucketID          string
	RemoteBucketID         string
	RemoteBucketName       string
	MaxQueueSize           int64
	DropNonRetryableData   bool
	NoDropNonRetryableData bool
	MaxAge                 int64
}

func (c Client) Create(ctx context.Context, params *CreateParams) error {
	if params.RemoteBucketID == "" && params.RemoteBucketName == "" {
		return fmt.Errorf("please supply one of: remote-bucket-id, remote-bucket-name")
	}
	orgID, err := params.GetOrgID(ctx, c.ActiveConfig, c.OrganizationsApi)
	if err != nil {
		return err
	}

	// set up a struct with required params
	body := api.ReplicationCreationRequest{
		Name:              params.Name,
		OrgID:             orgID,
		RemoteID:          params.RemoteID,
		LocalBucketID:     params.LocalBucketID,
		MaxQueueSizeBytes: params.MaxQueueSize,
		MaxAgeSeconds:     params.MaxAge,
	}

	if params.RemoteBucketID != "" {
		body.RemoteBucketID = &params.RemoteBucketID
	} else {
		body.RemoteBucketName = &params.RemoteBucketName
	}

	// set optional params if specified
	if params.Description != "" {
		body.Description = &params.Description
	}

	dropNonRetryableDataBoolPtr, err := dropNonRetryableDataBoolPtrFromFlags(params.DropNonRetryableData, params.NoDropNonRetryableData)
	if err != nil {
		return err
	}
	body.DropNonRetryableData = dropNonRetryableDataBoolPtr

	// send post request
	res, err := c.PostReplication(ctx).ReplicationCreationRequest(body).Execute()
	if err != nil {
		return fmt.Errorf("failed to create replication stream %q: %w", params.Name, err)
	}

	// print confirmation of new replication stream
	return c.printReplication(printReplicationOpts{replication: &res})
}

type ListParams struct {
	clients.OrgParams
	Name          string
	RemoteID      string
	LocalBucketID string
}

func (c Client) List(ctx context.Context, params *ListParams) error {
	orgID, err := params.GetOrgID(ctx, c.ActiveConfig, c.OrganizationsApi)
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
	ReplicationID          string
	Name                   string
	Description            string
	RemoteID               string
	RemoteBucketID         string
	MaxQueueSize           int64
	DropNonRetryableData   bool
	NoDropNonRetryableData bool
	MaxAge                 int64
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

	dropNonRetryableDataBoolPtr, err := dropNonRetryableDataBoolPtrFromFlags(params.DropNonRetryableData, params.NoDropNonRetryableData)
	if err != nil {
		return err
	}

	if dropNonRetryableDataBoolPtr != nil {
		body.SetDropNonRetryableData(*dropNonRetryableDataBoolPtr)
	}

	if params.MaxAge != 0 {
		body.SetMaxAgeSeconds(params.MaxAge)
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

func (c Client) DeleteWithRemoteID(ctx context.Context, conn api.RemoteConnection) error {
	reps, err := c.GetReplications(ctx).OrgID(conn.OrgID).RemoteID(conn.Id).Execute()
	if err != nil {
		return fmt.Errorf("failed to find replication streams with remote ID %q: %w", conn.Id, err)
	}

	if reps.Replications != nil {
		for _, rep := range reps.GetReplications() {
			if err := c.DeleteReplicationByID(ctx, rep.Id).Execute(); err != nil {
				return fmt.Errorf("failed to delete replication with ID %q: %w", rep.Id, err)
			}
		}
	} else {
		return fmt.Errorf("no replications found for remote ID %q", conn.Id)
	}

	printOpts := printReplicationOpts{
		replications: reps.GetReplications(),
		deleted:      true,
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

	headers := []string{"ID", "Name", "Org ID", "Remote ID", "Local Bucket ID", "Remote Bucket ID", "Remote Bucket Name",
		"Remaining Bytes to be Synced", "Current Queue Bytes on Disk", "Max Queue Bytes", "Latest Status Code", "Drop Non-Retryable Data"}
	if opts.deleted {
		headers = append(headers, "Deleted")
	}

	if opts.replication != nil {
		opts.replications = append(opts.replications, *opts.replication)
	}

	var rows []map[string]interface{}
	for _, r := range opts.replications {
		bucketID := r.GetRemoteBucketID()
		if r.GetRemoteBucketName() != "" {
			// This hides the default id that is required due to platform.ID implementation details
			bucketID = ""
		}
		row := map[string]interface{}{
			"ID":                           r.GetId(),
			"Name":                         r.GetName(),
			"Org ID":                       r.GetOrgID(),
			"Remote ID":                    r.GetRemoteID(),
			"Local Bucket ID":              r.GetLocalBucketID(),
			"Remote Bucket ID":             bucketID,
			"Remote Bucket Name":           r.GetRemoteBucketName(),
			"Remaining Bytes to be Synced": r.GetRemainingBytesToBeSynced(),
			"Current Queue Bytes on Disk":  r.GetCurrentQueueSizeBytes(),
			"Max Queue Bytes":              r.GetMaxQueueSizeBytes(),
			"Latest Status Code":           r.GetLatestResponseCode(),
			"Drop Non-Retryable Data":      r.GetDropNonRetryableData(),
		}
		if opts.deleted {
			row["Deleted"] = true
		}
		rows = append(rows, row)
	}

	return c.PrintTable(headers, rows...)
}

func dropNonRetryableDataBoolPtrFromFlags(dropNonRetryableData, noDropNonRetryableData bool) (*bool, error) {
	if dropNonRetryableData && noDropNonRetryableData {
		return nil, errors.New("cannot specify both --drop-non-retryable-data and --no-drop-non-retryable-data at the same time")
	}

	if dropNonRetryableData {
		return api.PtrBool(true), nil
	}

	if noDropNonRetryableData {
		return api.PtrBool(false), nil
	}

	return nil, nil
}
