package secret

import (
	"context"
	"fmt"
	"os"

	"github.com/influxdata/influx-cli/v2/api"
	"github.com/influxdata/influx-cli/v2/clients"
	"github.com/tcnksm/go-input"
)

type Client struct {
	clients.CLI
	api.SecretsApi
	api.OrganizationsApi
}

type secretPrintOpt struct {
	deleted bool
	secret  secret
	secrets []secret
}

type secret struct {
	key   string
	orgID string
}

type DeleteParams struct {
	clients.OrgParams
	Key string
}

func (c Client) Delete(ctx context.Context, params *DeleteParams) error {
	orgID, err := c.getOrgID(ctx, params.OrgParams)
	if err != nil {
		return err
	}

	if err := c.PostOrgsIDSecrets(ctx, orgID).Execute(); err != nil {
		return fmt.Errorf("failed to delete secret with key %q: %v", params.Key, err)
	}

	return c.printSecrets(secretPrintOpt{
		deleted: true,
		secret: secret{
			key:   params.Key,
			orgID: orgID,
		},
	})
}

type ListParams struct {
	clients.OrgParams
}

func (c Client) List(ctx context.Context, params *ListParams) error {

	orgID, err := c.getOrgID(ctx, params.OrgParams)
	if err != nil {
		return err
	}

	response, err := c.GetOrgsIDSecrets(ctx, orgID).Execute()
	if err != nil {
		return fmt.Errorf("failed to retrieve secret keys: %s", err)
	}

	secrets := response.GetSecrets()
	opts := make([]secret, 0, len(secrets))
	for _, entry := range secrets {
		opts = append(opts, secret{
			key:   entry,
			orgID: orgID,
		})
	}
	return c.printSecrets(secretPrintOpt{secrets: opts})
}

type UpdateParams struct {
	clients.OrgParams
	Key   string
	Value string
}

func (c Client) Update(ctx context.Context, params *UpdateParams) error {
	orgID, err := c.getOrgID(ctx, params.OrgParams)
	if err != nil {
		return err
	}

	ui := &input.UI{
		Writer: c.StdIO,
		Reader: os.Stdin,
	}
	var secretVal string
	if params.Value != "" {
		secretVal = params.Value
	} else {
		secretVal = getSecret(ui)
	}

	err = c.PatchOrgsIDSecrets(ctx, orgID).RequestBody(map[string]string{
		params.Key: secretVal,
	}).Execute()
	if err != nil {
		return fmt.Errorf("failed to update secret with key %q: %v", params.Key, err)
	}

	return c.printSecrets(secretPrintOpt{
		secret: secret{
			key:   params.Key,
			orgID: orgID,
		},
	})
}

func (c Client) printSecrets(opts secretPrintOpt) error {
	if c.PrintAsJSON {
		var v interface{} = opts.secrets
		if opts.secrets == nil {
			v = opts.secret
		}
		return c.PrintJSON(v)
	}

	headers := []string{"Key", "Organization ID"}
	if opts.deleted {
		headers = append(headers, "Deleted")
	}
	if opts.secrets == nil {
		opts.secrets = append(opts.secrets, opts.secret)
	}

	var rows []map[string]interface{}
	for _, u := range opts.secrets {
		row := map[string]interface{}{
			"Key":             u.key,
			"Organization ID": u.orgID,
		}
		if opts.deleted {
			row["Deleted"] = true
		}
		rows = append(rows, row)
	}
	return c.PrintTable(headers, rows...)
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
