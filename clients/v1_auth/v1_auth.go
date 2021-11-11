package v1_auth

import (
	"context"
	"fmt"

	"github.com/influxdata/influx-cli/v2/api"
	"github.com/influxdata/influx-cli/v2/clients"
	"github.com/influxdata/influx-cli/v2/pkg/influxid"
)

type Client struct {
	clients.CLI
	api.LegacyAuthorizationsApi
	api.UsersApi
	api.OrganizationsApi
}

type AuthLookupParams struct {
	ID       string
	Username string
}

func (p AuthLookupParams) Validate() (err error) {
	if p.Username == "" && p.ID == "" {
		err = fmt.Errorf("id or username required")
	} else if p.Username != "" && p.ID != "" {
		err = fmt.Errorf("specify id or username, not both")
	}
	return
}

func (p AuthLookupParams) IsSet() bool {
	return p.ID != "" || p.Username != ""
}

type v1PrintOpts struct {
	deleted bool
	token   *v1Token
	tokens  []v1Token
}

type v1Token struct {
	ID          string   `json:"id"`
	Description string   `json:"description"`
	Token       string   `json:"token"`
	Status      string   `json:"status"`
	UserName    string   `json:"userName"`
	UserID      string   `json:"userID"`
	Permissions []string `json:"permissions"`
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
	if params.Password != "" && params.NoPassword {
		return fmt.Errorf("only one of --password and --no-password may be specified")
	}

	orgID, err := params.GetOrgID(ctx, c.ActiveConfig, c.OrganizationsApi)
	if err != nil {
		return err
	}

	// verify an existing token with the same username doesn't already exist
	auths, err := c.LegacyAuthorizationsApi.GetLegacyAuthorizations(ctx).Token(params.Username).Execute()
	if apiError, ok := err.(api.ApiError); ok {
		if apiError.ErrorCode() != api.ERRORCODE_NOT_FOUND {
			return fmt.Errorf("failed to verify username %q has no auth: %w", params.Username, err)
		} else if len(auths.GetAuthorizations()) != 0 {
			return fmt.Errorf("authorization with username %q already exists", params.Username)
		}
	}

	password := params.Password
	if password == "" && !params.NoPassword {
		pass, err := c.StdIO.GetPassword("Please type your password")
		if err != nil {
			return err
		}
		password = pass
	}

	bucketPerms := []struct {
		action string
		perms  []string
	}{
		{action: "read", perms: params.ReadBucket},
		{action: "write", perms: params.WriteBucket},
	}

	var permissions []api.Permission
	for _, bp := range bucketPerms {
		for _, p := range bp.perms {
			p := p
			// verify the input ID
			if err := influxid.Validate(p); err != nil {
				return fmt.Errorf("invalid bucket ID '%s': %w (did you pass a bucket name instead of an ID?)", p, err)
			}

			newPerm := api.Permission{
				Action: bp.action,
				Resource: api.PermissionResource{
					Type:  "buckets",
					Id:    &p,
					OrgID: &orgID,
				},
			}
			permissions = append(permissions, newPerm)
		}
	}
	authReq := api.LegacyAuthorizationPostRequest{
		Description: &params.Desc,
		OrgID:       orgID,
		Permissions: permissions,
		Token:       &params.Username,
	}

	newAuth, err := c.LegacyAuthorizationsApi.PostLegacyAuthorizations(ctx).LegacyAuthorizationPostRequest(authReq).Execute()
	if err != nil {
		return fmt.Errorf("failed to create new authorization: %w", err)
	}

	if password != "" {
		err := c.LegacyAuthorizationsApi.PostLegacyAuthorizationsIDPassword(ctx, newAuth.GetId()).
			PasswordResetBody(api.PasswordResetBody{Password: password}).
			Execute()
		if err != nil {
			_ = c.LegacyAuthorizationsApi.DeleteLegacyAuthorizationsID(ctx, newAuth.GetId()).Execute()
			return fmt.Errorf("failed to set password for %q: %w", params.Username, err)
		}
	}

	usr, err := c.UsersApi.GetUsersID(ctx, newAuth.GetUserID()).Execute()
	if err != nil {
		return err
	}

	ps := make([]string, 0, len(newAuth.Permissions))
	for _, p := range newAuth.Permissions {
		ps = append(ps, p.String())
	}

	return c.printV1Tokens(&v1PrintOpts{
		token: &v1Token{
			ID:          newAuth.GetId(),
			Description: newAuth.GetDescription(),
			Token:       newAuth.GetToken(),
			Status:      newAuth.GetStatus(),
			UserName:    usr.GetName(),
			UserID:      usr.GetId(),
			Permissions: ps,
		},
	})
}

type RemoveParams struct {
	AuthLookupParams
}

func (c Client) Remove(ctx context.Context, params *RemoveParams) error {
	id, err := c.getAuthReqID(ctx, &params.AuthLookupParams)
	if err != nil {
		return err
	}

	auth, err := c.LegacyAuthorizationsApi.GetLegacyAuthorizationsID(ctx, id).Execute()
	if err != nil {
		return fmt.Errorf("could not find Authorization from ID %q: %w", id, err)
	}
	err = c.LegacyAuthorizationsApi.DeleteLegacyAuthorizationsID(ctx, id).Execute()
	if err != nil {
		return fmt.Errorf("could not remove Authorization with ID %q: %w", id, err)
	}

	usr, err := c.UsersApi.GetUsersID(ctx, auth.GetUserID()).Execute()
	if err != nil {
		return fmt.Errorf("could not find user from user ID %q: %w", auth.GetUserID(), err)
	}

	ps := make([]string, 0, len(auth.GetPermissions()))
	for _, p := range auth.GetPermissions() {
		ps = append(ps, p.String())
	}

	return c.printV1Tokens(&v1PrintOpts{
		deleted: true,
		token: &v1Token{
			ID:          auth.GetId(),
			Description: auth.GetDescription(),
			Token:       auth.GetToken(),
			Status:      auth.GetStatus(),
			UserName:    usr.GetName(),
			UserID:      usr.GetId(),
			Permissions: ps,
		},
	})
}

type ListParams struct {
	clients.OrgParams
	AuthLookupParams
	User   string
	UserID string
}

func (c Client) List(ctx context.Context, params *ListParams) error {
	req := c.LegacyAuthorizationsApi.GetLegacyAuthorizations(ctx)

	if params.User != "" {
		req = req.User(params.User)
	}
	if params.UserID != "" {
		req = req.UserID(params.UserID)
	}
	if params.OrgID != "" {
		req = req.OrgID(params.OrgID)
	}
	if params.OrgName != "" {
		req = req.Org(params.OrgName)
	}
	if params.Username != "" {
		req = req.Token(params.Username)
	}
	if params.ID != "" {
		req = req.AuthID(params.ID)
	}

	auths, err := req.Execute()
	if err != nil {
		return fmt.Errorf("could not find Authorizations for specified arguments: %w", err)
	}

	var tokens []v1Token
	for _, a := range auths.GetAuthorizations() {
		var permissions []string
		for _, p := range a.GetPermissions() {
			permissions = append(permissions, p.String())
		}

		usr, err := c.UsersApi.GetUsersID(ctx, a.GetUserID()).Execute()
		if err != nil {
			return fmt.Errorf("could not find user with ID %q: %w", a.GetUserID(), err)
		}

		tokens = append(tokens, v1Token{
			ID:          a.GetId(),
			Description: a.GetDescription(),
			Token:       a.GetToken(),
			Status:      a.GetStatus(),
			UserName:    usr.GetName(),
			UserID:      usr.GetId(),
			Permissions: permissions,
		})
	}

	return c.printV1Tokens(&v1PrintOpts{tokens: tokens})
}

type ActiveParams struct {
	AuthLookupParams
}

func (c Client) SetActive(ctx context.Context, params *ActiveParams, active bool) error {
	id, err := c.getAuthReqID(ctx, &params.AuthLookupParams)
	if err != nil {
		return err
	}

	req := c.LegacyAuthorizationsApi.PatchLegacyAuthorizationsID(ctx, id)
	var status string
	if active {
		status = "active"
	} else {
		status = "inactive"
	}
	req = req.AuthorizationUpdateRequest(api.AuthorizationUpdateRequest{
		Status: &status,
	})

	auth, err := req.Execute()
	if err != nil {
		return fmt.Errorf("failed to update status of authorization with ID %q: %w", id, err)
	}

	usr, err := c.UsersApi.GetUsersID(ctx, auth.GetUserID()).Execute()
	if err != nil {
		return fmt.Errorf("could not find user from User ID %q: %w", auth.GetUserID(), err)
	}

	ps := make([]string, 0, len(auth.GetPermissions()))
	for _, p := range auth.GetPermissions() {
		ps = append(ps, p.String())
	}

	return c.printV1Tokens(&v1PrintOpts{
		token: &v1Token{
			ID:          auth.GetId(),
			Description: auth.GetDescription(),
			Token:       auth.GetToken(),
			Status:      auth.GetStatus(),
			UserName:    usr.GetName(),
			UserID:      usr.GetId(),
			Permissions: ps,
		},
	})
}

type SetPasswordParams struct {
	AuthLookupParams
	Password string
}

func (c Client) SetPassword(ctx context.Context, params *SetPasswordParams) error {
	id, err := c.getAuthReqID(ctx, &params.AuthLookupParams)
	if err != nil {
		return err
	}

	password := params.Password
	if password == "" {
		pass, err := c.StdIO.GetPassword("Please type new password")
		if err != nil {
			return err
		}
		password = pass
	}

	err = c.LegacyAuthorizationsApi.PostLegacyAuthorizationsIDPassword(ctx, id).
		PasswordResetBody(api.PasswordResetBody{Password: password}).
		Execute()
	if err != nil {
		return fmt.Errorf("error setting password: %w", err)
	}

	return nil
}

func (c Client) printV1Tokens(params *v1PrintOpts) error {
	if c.PrintAsJSON {
		var v interface{}
		if params.token != nil {
			v = params.token
		} else {
			v = params.tokens
		}
		return c.PrintJSON(v)
	}

	headers := []string{
		"ID",
		"Description",
		"Username",
		"v2 User Name",
		"v2 User ID",
		"Permissions",
	}
	if params.deleted {
		headers = append(headers, "Deleted")
	}
	if params.token != nil {
		params.tokens = append(params.tokens, *params.token)
	}

	var rows []map[string]interface{}
	for _, u := range params.tokens {
		row := map[string]interface{}{
			"ID":           u.ID,
			"Description":  u.Description,
			"Username":     u.Token,
			"v2 User Name": u.UserName,
			"v2 User ID":   u.UserID,
			"Permissions":  u.Permissions,
		}
		if params.deleted {
			row["Deleted"] = true
		}
		rows = append(rows, row)
	}
	return c.PrintTable(headers, rows...)
}

func (c Client) getAuthReqID(ctx context.Context, params *AuthLookupParams) (id string, err error) {
	if params.ID != "" {
		id = params.ID
	} else {
		var auths api.Authorizations
		auths, err = c.LegacyAuthorizationsApi.GetLegacyAuthorizations(ctx).Token(params.Username).Execute()
		if err != nil || len(auths.GetAuthorizations()) == 0 {
			err = fmt.Errorf("could not find v1 auth with token (username) %q: %w", params.Username, err)
		} else {
			id = auths.GetAuthorizations()[0].GetId()
		}
	}
	return
}
