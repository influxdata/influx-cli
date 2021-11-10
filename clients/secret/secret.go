package secret

import (
	"context"
	"fmt"

	"github.com/influxdata/influx-cli/v2/api"
	"github.com/influxdata/influx-cli/v2/clients"
)

type Client struct {
	clients.CLI
	api.SecretsApi
	api.OrganizationsApi
}

type secretPrintOpt struct {
	deleted bool
	secret  *secret
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
	orgID, err := params.GetOrgID(ctx, c.ActiveConfig, c.OrganizationsApi)
	if err != nil {
		return err
	}

	// PostOrgsIDSecrets is used to remove secrets from an organization.
	// The name is generated from the operationId in the
	// orgs_orgsID_secrets_delete.yml path.
	err = c.PostOrgsIDSecrets(ctx, orgID).
		SecretKeys(api.SecretKeys{Secrets: &[]string{params.Key}}).
		Execute()
	if err != nil {
		return fmt.Errorf("failed to delete secret with key %q: %w", params.Key, err)
	}

	return c.printSecrets(secretPrintOpt{
		deleted: true,
		secret: &secret{
			key:   params.Key,
			orgID: orgID,
		},
	})
}

type ListParams struct {
	clients.OrgParams
}

func (c Client) List(ctx context.Context, params *ListParams) error {
	orgID, err := params.GetOrgID(ctx, c.ActiveConfig, c.OrganizationsApi)
	if err != nil {
		return err
	}

	response, err := c.GetOrgsIDSecrets(ctx, orgID).Execute()
	if err != nil {
		return fmt.Errorf("failed to retrieve secret keys: %w", err)
	}

	secrets := response.GetSecrets()
	opts := make([]secret, len(secrets))
	for i, entry := range secrets {
		opts[i] = secret{
			key:   entry,
			orgID: orgID,
		}
	}
	return c.printSecrets(secretPrintOpt{secrets: opts})
}

type UpdateParams struct {
	clients.OrgParams
	Key   string
	Value string
}

func (c Client) Update(ctx context.Context, params *UpdateParams) error {
	orgID, err := params.GetOrgID(ctx, c.ActiveConfig, c.OrganizationsApi)
	if err != nil {
		return err
	}

	var secretVal string
	if params.Value != "" {
		secretVal = params.Value
	} else {
		secretVal, err = c.StdIO.GetSecret("Please type your secret", 0)
		if err != nil {
			return err
		}
	}

	err = c.PatchOrgsIDSecrets(ctx, orgID).
		RequestBody(map[string]string{params.Key: secretVal}).
		Execute()
	if err != nil {
		return fmt.Errorf("failed to update secret with key %q: %w", params.Key, err)
	}

	return c.printSecrets(secretPrintOpt{
		secret: &secret{
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
	if opts.secret != nil {
		opts.secrets = append(opts.secrets, *opts.secret)
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
