package auth

import (
	"context"
	"fmt"

	"github.com/influxdata/influx-cli/v2/api"
	"github.com/influxdata/influx-cli/v2/clients"
	"github.com/influxdata/influx-cli/v2/pkg/influxid"
)

type Client struct {
	clients.CLI
	api.AuthorizationsApi
	api.UsersApi
}

const (
	ReadAction  = "read"
	WriteAction = "write"
)

type token struct {
	ID          string   `json:"id"`
	Description string   `json:"description"`
	Token       string   `json:"token"`
	Status      string   `json:"status"`
	UserName    string   `json:"userName"`
	UserID      string   `json:"userID"`
	Permissions []string `json:"permissions"`
}

type printParams struct {
	deleted bool
	token   *token
	tokens  []token
}

type CreateParams struct {
	clients.OrgParams
	User        string
	Description string

	WriteUserPermission bool
	ReadUserPermission  bool

	WriteBucketsPermission bool
	ReadBucketsPermission  bool

	WriteBucketPermissions []string
	ReadBucketPermissions  []string

	WriteTasksPermission bool
	ReadTasksPermission  bool

	WriteTelegrafsPermission bool
	ReadTelegrafsPermission  bool

	WriteOrganizationsPermission bool
	ReadOrganizationsPermission  bool

	WriteDashboardsPermission bool
	ReadDashboardsPermission  bool

	WriteCheckPermission bool
	ReadCheckPermission  bool

	WriteNotificationRulePermission bool
	ReadNotificationRulePermission  bool

	WriteNotificationEndpointPermission bool
	ReadNotificationEndpointPermission  bool

	WriteDBRPPermission bool
	ReadDBRPPermission  bool
}

func (c Client) Create(ctx context.Context, params *CreateParams) error {

	orgID := params.OrgID.String()
	orgName := params.OrgName
	bucketPerms := []struct {
		action string
		perms  []string
	}{
		{action: ReadAction, perms: params.ReadBucketPermissions},
		{action: WriteAction, perms: params.WriteBucketPermissions},
	}

	var permissions []api.Permission
	for _, bp := range bucketPerms {
		for _, p := range bp.perms {
			// verify the input ID
			if _, err := influxid.IDFromString(p); err != nil {
				return fmt.Errorf("invalid bucket ID '%s': %w (did you pass a bucket name instead of an ID?)", p, err)
			}

			newPerm := api.Permission{
				Action:   bp.action,
				Resource: makePermResource("buckets", p, orgID, orgName),
			}
			permissions = append(permissions, newPerm)
		}
	}

	providedPerm := []struct {
		readPerm, writePerm bool
		ResourceType        string
	}{
		{
			readPerm:     params.ReadBucketsPermission,
			writePerm:    params.WriteBucketsPermission,
			ResourceType: "buckets",
		},
		{
			readPerm:     params.ReadCheckPermission,
			writePerm:    params.WriteCheckPermission,
			ResourceType: "checks",
		},
		{
			readPerm:     params.ReadDashboardsPermission,
			writePerm:    params.WriteDashboardsPermission,
			ResourceType: "dashboards",
		},
		{
			readPerm:     params.ReadNotificationEndpointPermission,
			writePerm:    params.WriteNotificationEndpointPermission,
			ResourceType: "notificationEndpoints",
		},
		{
			readPerm:     params.ReadNotificationRulePermission,
			writePerm:    params.WriteNotificationRulePermission,
			ResourceType: "notificationRules",
		},
		{
			readPerm:     params.ReadOrganizationsPermission,
			writePerm:    params.WriteOrganizationsPermission,
			ResourceType: "orgs",
		},
		{
			readPerm:     params.ReadTasksPermission,
			writePerm:    params.WriteTasksPermission,
			ResourceType: "tasks",
		},
		{
			readPerm:     params.ReadTelegrafsPermission,
			writePerm:    params.WriteTelegrafsPermission,
			ResourceType: "telegrafs",
		},
		{
			readPerm:     params.ReadUserPermission,
			writePerm:    params.WriteUserPermission,
			ResourceType: "users",
		},
		{
			readPerm:     params.ReadDBRPPermission,
			writePerm:    params.WriteDBRPPermission,
			ResourceType: "dbrp",
		},
	}

	for _, provided := range providedPerm {
		var actions []string
		if provided.readPerm {
			actions = append(actions, ReadAction)
		}
		if provided.writePerm {
			actions = append(actions, WriteAction)
		}

		for _, action := range actions {
			p := api.Permission{
				Action:   action,
				Resource: makePermResource(provided.ResourceType, "", orgID, orgName),
			}
			permissions = append(permissions, p)
		}
	}

	// Get the user ID because the command only takes a username, not ID
	users, err := c.UsersApi.GetUsers(ctx).Name(params.User).Execute()
	if err != nil || len(users.GetUsers()) == 0 {
		return fmt.Errorf("could not find user with name %q: %q", params.User, err)
	}
	userID := users.GetUsers()[0].GetId()

	authReq := api.AuthorizationPostRequest{
		Description: &params.Description,
		OrgID:       orgID,
		UserID:      &userID,
		Permissions: permissions,
	}

	auth, err := c.PostAuthorizations(ctx).AuthorizationPostRequest(authReq).Execute()
	if err != nil {
		return fmt.Errorf("could not write auth with provided arguments: %w", err)
	}

	ps := make([]string, 0, len(auth.GetPermissions()))
	for _, p := range auth.GetPermissions() {
		ps = append(ps, permString(p))
	}

	return c.printAuth(printParams{
		token: &token{
			ID:          auth.GetId(),
			Description: auth.GetDescription(),
			Token:       auth.GetToken(),
			Status:      auth.GetStatus(),
			UserName:    auth.GetUser(),
			UserID:      auth.GetUserID(),
			Permissions: ps,
		},
	})
}

func (c Client) Remove(ctx context.Context, authID string) error {
	// check if auth exists first for better error logging, and to
	// acquire the auth that was deleted since the delete
	// request does not return the Authorization object.
	a, err := c.GetAuthorizationsID(ctx, authID).Execute()
	if err != nil {
		return fmt.Errorf("could not find auth with ID %q: %w", authID, err)
	}

	if err := c.DeleteAuthorizationsID(ctx, authID).Execute(); err != nil {
		return fmt.Errorf("could not remove auth with ID %q: %w", authID, err)
	}

	ps := make([]string, 0, len(a.Permissions))
	for _, p := range a.Permissions {
		ps = append(ps, permString(p))
	}

	return c.printAuth(printParams{
		deleted: true,
		token: &token{
			ID:          a.GetId(),
			Description: a.GetDescription(),
			Token:       a.GetToken(),
			Status:      a.GetStatus(),
			UserName:    a.GetUser(),
			UserID:      a.GetUserID(),
			Permissions: ps,
		},
	})
}

type ListParams struct {
	clients.OrgParams
	Id     string
	User   string
	UserID string
}

func (c Client) List(ctx context.Context, params *ListParams) error {

	// If ID parameter is set, search by that over other filters
	if params.Id != "" {
		return c.findAuthorization(ctx, params.Id)
	}

	req := c.GetAuthorizations(ctx)
	if params.User != "" {
		req.User(params.User)
	}
	if params.UserID != "" {
		req.UserID(params.UserID)
	}
	if params.OrgName != "" {
		req.Org(params.OrgName)
	}
	if params.OrgID.Valid() {
		req.OrgID(params.OrgID.String())
	}

	auths, err := req.Execute()
	if err != nil {
		return fmt.Errorf("could not find authorization with given parameters: %w", err)
	}

	var tokens []token
	for _, a := range auths.GetAuthorizations() {
		var ps []string
		for _, p := range a.GetPermissions() {
			ps = append(ps, permString(p))
		}

		tokens = append(tokens, token{
			ID:          a.GetId(),
			Description: a.GetDescription(),
			Token:       a.GetToken(),
			Status:      a.GetStatus(),
			UserName:    a.GetUser(),
			UserID:      a.GetUserID(),
			Permissions: ps,
		})
	}

	return c.printAuth(printParams{tokens: tokens})
}

func (c Client) findAuthorization(ctx context.Context, authID string) error {
	a, err := c.GetAuthorizationsID(ctx, authID).Execute()
	if err != nil {
		return fmt.Errorf("could not find authorization with ID %q: %w", authID, err)
	}

	ps := make([]string, 0, len(a.GetPermissions()))
	for _, p := range a.GetPermissions() {
		ps = append(ps, permString(p))
	}

	return c.printAuth(printParams{
		token: &token{
			ID:          a.GetId(),
			Description: a.GetDescription(),
			Token:       a.GetToken(),
			Status:      a.GetStatus(),
			UserName:    a.GetUser(),
			UserID:      a.GetUserID(),
			Permissions: ps,
		},
	})
}

func (c Client) SetActive(ctx context.Context, authID string, active bool) error {

	// check if auth exists first for better error logging
	if _, err := c.GetAuthorizationsID(ctx, authID).Execute(); err != nil {
		return fmt.Errorf("could not find auth with ID %q: %w", authID, err)
	}

	var status string
	if active {
		status = "active"
	} else {
		status = "inactive"
	}
	a, err := c.PatchAuthorizationsID(ctx, authID).
		AuthorizationUpdateRequest(api.AuthorizationUpdateRequest{Status: &status}).
		Execute()
	if err != nil {
		return fmt.Errorf("could not update status of auth with ID %q: %w", authID, err)
	}

	ps := make([]string, 0, len(a.GetPermissions()))
	for _, p := range a.GetPermissions() {
		ps = append(ps, permString(p))
	}

	return c.printAuth(printParams{
		token: &token{
			ID:          a.GetId(),
			Description: a.GetDescription(),
			Token:       a.GetToken(),
			Status:      a.GetStatus(),
			UserName:    a.GetUser(),
			UserID:      a.GetUserID(),
			Permissions: ps,
		},
	})
}

func (c Client) printAuth(opts printParams) error {
	if c.PrintAsJSON {
		var v interface{}
		if opts.token != nil {
			v = opts.token
		} else {
			v = opts.tokens
		}
		return c.PrintJSON(v)
	}

	headers := []string{
		"ID",
		"Description",
		"Token",
		"User Name",
		"User ID",
		"Permissions",
	}
	if opts.deleted {
		headers = append(headers, "Deleted")
	}

	if opts.token != nil {
		opts.tokens = append(opts.tokens, *opts.token)
	}

	var rows []map[string]interface{}
	for _, t := range opts.tokens {
		row := map[string]interface{}{
			"ID":          t.ID,
			"Description": t.Description,
			"Token":       t.Token,
			"User Name":   t.UserName,
			"User ID":     t.UserID,
			"Permissions": t.Permissions,
		}
		if opts.deleted {
			row["Deleted"] = true
		}
		rows = append(rows, row)
	}
	return c.PrintTable(headers, rows...)
}

func permString(p api.Permission) string {
	ret := p.GetAction() + ":"
	r := p.GetResource()

	if r.GetOrgID() != "" {
		ret += "orgs/" + r.GetOrgID()
	}
	ret += "/" + r.GetType()
	if r.GetId() != "" {
		ret += "/" + r.GetId()
	}
	return ret
}

func makePermResource(permType string, id string, orgId string, org string) api.PermissionResource {
	return api.PermissionResource{
		Type:  permType,
		Id:    &id,
		OrgID: &orgId,
		Org:   &org,
	}
}
