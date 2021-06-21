package v1auth

import (
	"context"

	"github.com/influxdata/influx-cli/v2/api"
	"github.com/influxdata/influx-cli/v2/clients"
)

type Client struct {
	clients.CLI
	api.AuthorizationsApi
}

type CreateParams struct {
	clients.OrgParams
	Desc        string
	Username    string
	Password    string
	NoPassword  bool
	ReadBucket  []string
	WriteBucket []string
}

func (c Client) Create(ctx context.Context, params *CreateParams) error {
	return nil
}

type RemoveParams struct {
	clients.AuthLookupParams
}

func (c Client) Remove(ctx context.Context, params *RemoveParams) error {
	return nil
}

type ListParams struct {
	clients.OrgParams
	clients.AuthLookupParams
	User     string
	UserID   string
}

func (c Client) List(ctx context.Context, params *ListParams) error {
	return nil
}

type SetActiveParams struct {
	clients.AuthLookupParams
}

func (c Client) SetActive(ctx context.Context, params *SetActiveParams) error {
	return nil
}

type SetInactiveParams struct {
	clients.AuthLookupParams
}

func (c Client) SetInactive(ctx context.Context, params *SetInactiveParams) error {
	return nil
}

type SetPasswordParams struct {
	clients.AuthLookupParams
	Password string
}

func (c Client) SetPassword(ctx context.Context, params *SetPasswordParams) error {
	return nil
}
