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
	api.AuthorizationsApi
	api.UsersApi
	api.OrganizationsApi
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

	orgID, err := c.getOrgID(ctx, params.OrgParams)
	if err != nil {
		return err
	}

	ids, err := c.UsersApi.GetUsers(ctx).Name(params.Username).Execute()
	if err != nil || len(ids.GetUsers()) == 0 {
		return fmt.Errorf("could not get user from Username %q: %w", params.Username, err)
	}

	// verify an existing token with the same username doesn't already exist
	auth, err := c.AuthorizationsApi.GetAuthorizationsID(ctx, ids.GetUsers()[0].GetId()).Execute()
	authPtr := &auth
	if &authPtr != nil {
		return fmt.Errorf("authorization with username %q already exists", params.Username)
	}
	if err != nil {
		return fmt.Errorf("failed to verify username %q has no auth: %w", params.Username, err)
	}

	password := params.Password
	if password == "" && !params.NoPassword {
		password, err = c.CLI.StdIO.GetSecret("Please enter your password", 0)
		if err != nil {
			return err
		}
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
			// verify the input ID
			if _, err := influxid.IDFromString(p); err != nil {
				return fmt.Errorf("invalid bucket ID '%s': %w (did you pass a bucket name instead of an ID?)", p, err)
			}

			newPerm := api.Permission{
				Action: bp.action,
				Resource: api.PermissionResource{
					Type:  "buckets",
					Id:    *api.NewNullableString(&p),
					OrgID: *api.NewNullableString(&orgID),
				},
			}
			permissions = append(permissions, newPerm)
		}
	}

	authReq := &api.Authorization{
		Description: &params.Desc,
		OrgID:       orgID,
		Permissions: permissions,
		Token:       &params.Username,
	}

	// TODO Something wrong here
	newAuth, err := c.AuthorizationsApi.PostAuthorizations(ctx).AuthorizationPostRequest(authReq).Execute()
	if err != nil {
		return fmt.Errorf("failed to create new authorization: %w", err)
	}

	if password != "" {
		if err := c.UsersApi.PostUsersIDPassword(ctx, newAuth.GetUserID()).
			PasswordResetBody(api.PasswordResetBody{
				Password: password,
			}).Execute(); err != nil {
			_ = c.AuthorizationsApi.DeleteAuthorizationsID(ctx, newAuth.GetId()).Execute()
			return fmt.Errorf("failed to update password for user ID %q: %w", newAuth.GetUserID(), err)
		}
	}

	usr, err := c.UsersApi.GetUsersID(ctx, newAuth.GetUserID()).Execute()
	if err != nil {
		return err
	}

	ps := make([]string, 0, len(auth.Permissions))
	for _, p := range auth.Permissions {
		ps = append(ps, permString(p))
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
	clients.AuthLookupParams
}

func (c Client) Remove(ctx context.Context, params *RemoveParams) error {
	id, err := c.getAuthReqID(ctx, &params.AuthLookupParams)
	if err != nil {
		return err
	}

	auth, err1 := c.AuthorizationsApi.GetAuthorizationsID(ctx, id).Execute()
	err2 := c.AuthorizationsApi.DeleteAuthorizationsID(ctx, id).Execute()
	if err1 != nil || err2 != nil {
		return fmt.Errorf("could not find Authorization from ID %q: %w", id, err1)
	}

	usr, err := c.UsersApi.GetUsersID(ctx, auth.GetUserID()).Execute()
	if err != nil {
		return fmt.Errorf("could not find user from user ID %q: %w", auth.GetUserID(), err)
	}

	ps := make([]string, 0, len(auth.GetPermissions()))
	for _, p := range auth.GetPermissions() {
		ps = append(ps, permString(p))
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
	clients.AuthLookupParams
	User   string
	UserID string
}

func (c Client) List(ctx context.Context, params *ListParams) error {
	req := c.AuthorizationsApi.GetAuthorizations(ctx)

	if params.User != "" {
		req = req.User(params.User)
	}
	if params.OrgID.Valid() {
		req = req.OrgID(params.OrgID.String())
	}
	if params.OrgName != "" {
		req = req.Org(params.OrgName)
	}
	var id string
	if params.AuthLookupParams.IsSet() {
		newId, err := c.getAuthReqID(ctx, &params.AuthLookupParams)
		if err != nil {
			return fmt.Errorf("could not get User from Username %q: %w", params.Username, err)
		}
		id = newId
	}
	if params.UserID != "" {
		id = params.UserID
	}
	if id != "" {
		req = req.UserID(id)
	}

	auths, err := req.Execute()
	if err != nil {
		return fmt.Errorf("could not find Authorizations for specified arguments %w", err)
	}

	var tokens []v1Token
	for _, a := range auths.GetAuthorizations() {
		var permissions []string
		for _, p := range a.GetPermissions() {
			permissions = append(permissions, permString(p))
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
	clients.AuthLookupParams
}

func (c Client) SetActive(ctx context.Context, params *ActiveParams, active bool) error {
	id, err := c.getAuthReqID(ctx, &params.AuthLookupParams)
	if err != nil {
		return err
	}

	req := c.AuthorizationsApi.PatchAuthorizationsID(ctx, id)
	var status string
	if active {
		status = "active"
	} else {
		status = "inactive"
	}
	req.AuthorizationUpdateRequest(api.AuthorizationUpdateRequest{
		Status: &status,
	})

	auth, err := req.Execute()
	if err != nil {
		return fmt.Errorf("could not find Authorization with ID %q: %w", id, err)
	}

	usr, err := c.UsersApi.GetUsersID(ctx, auth.GetUserID()).Execute()
	if err != nil {
		return fmt.Errorf("could not find user from User ID %q: %w", auth.GetUserID(), err)
	}

	ps := make([]string, 0, len(auth.GetPermissions()))
	for _, p := range auth.GetPermissions() {
		ps = append(ps, permString(p))
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
	clients.AuthLookupParams
	Password string
}

func (c Client) SetPassword(ctx context.Context, params *SetPasswordParams) error {
	id, err := c.getAuthReqID(ctx, &params.AuthLookupParams)
	if err != nil {
		return err
	}

	auth, err := c.AuthorizationsApi.GetAuthorizationsID(ctx, id).Execute()
	if err != nil {
		return fmt.Errorf("could not find authorization with User ID %q: %w", id, err)
	}

	password := params.Password
	if password == "" {
		password, err = c.StdIO.GetSecret("Please enter new password", 0)
		if err != nil {
			return err
		}
	}

	if err := c.UsersApi.PostUsersIDPassword(ctx, auth.GetUserID()).Execute(); err != nil {
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
		"Token",
		"User Name",
		"User ID",
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
			"ID":          u.ID,
			"Description": u.Description,
			"Token":       u.Token,
			"User Name":   u.UserName,
			"User ID":     u.UserID,
			"Permissions": u.Permissions,
		}
		if params.deleted {
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
		ret += "orgs" + r.GetOrgID()
	}
	ret += r.GetType()
	if r.GetId() != "" {
		ret += r.GetId()
	}
	return ret
}

func (c Client) getAuthReqID(ctx context.Context, params *clients.AuthLookupParams) (id string, err error) {
	if params.ID.Valid() {
		id = params.ID.String()
	} else {
		userid, err := c.UsersApi.GetUsers(ctx).Name(params.Username).Execute()
		if err != nil || len(userid.GetUsers()) == 0 {
			err = fmt.Errorf("could not find User ID from username %q: %w", params.Username, err)
		} else {
			id = userid.GetUsers()[0].GetId()
		}
	}
	return
}

func (c Client) getOrgID(ctx context.Context, params clients.OrgParams) (string, error) {
	if params.OrgID.Valid() || params.OrgName != "" || c.ActiveConfig.Org != "" {
		if params.OrgID.Valid() {
			return params.OrgID.String(), nil
		}
		for _, name := range []string{params.OrgName, c.ActiveConfig.Org} {
			if name != "" {
				org, err := c.GetOrgs(ctx).Org(name).Execute()
				if err != nil {
					return "", fmt.Errorf("failed to lookup org with name %q: %w", name, err)
				}
				if len(org.GetOrgs()) == 0 {
					return "", fmt.Errorf("no organization with name %q: %w", name, err)
				}
				return org.GetOrgs()[0].GetId(), nil
			}
		}
	}
	return "", fmt.Errorf("org or org-id must be provided")
}
