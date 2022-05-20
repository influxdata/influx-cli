package v1repl

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/influxdata/influx-cli/v2/api"
	"github.com/influxdata/influx-cli/v2/clients"
)

type Client struct {
	clients.CLI
	PersistentQueryParams
	api.LegacyQueryApi
	api.OrganizationsApi
}

type PersistentQueryParams struct {
	clients.OrgParams
	Db     string // bucketID
	P      string // password OR token
	U      string
	Rp     string
	Epoch  string
	Format FormatType
}

func DefaultPersistentQueryParams() PersistentQueryParams {
	return PersistentQueryParams{
		Format: CsvFormat,
		Epoch:  "n",
	}

}

type FormatType string

var (
	CsvFormat  FormatType = "csv"
	JsonFormat FormatType = "json"
)

func (c Client) Create(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			return nil
		default:
			query := c.Prompt()
			response, err := c.Query(ctx, query)
			if err != nil {
				color.HiRed("Query failed.")
				color.Red("%v", err)
				continue
			}
			if c.Format == CsvFormat {
				fmt.Fprintf(c.CLI.StdIO, "%s", *response)
			} else {
				return fmt.Errorf("unimplemented format")
			}
			fmt.Fprintf(c.CLI.StdIO, "\n")
		}
	}
}

func (c Client) Prompt() string {
	sb := strings.Builder{}
	scanner := bufio.NewScanner(os.Stdin)
	// oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	// if err != nil {
	// 	os.Exit(0)
	// }
	// defer term.Restore(int(os.Stdin.Fd()), oldState)
	fmt.Fprintf(c.StdIO, "%s", color.GreenString("> "))
	for {

		// reads user input until \n by default
		scanner.Scan()
		// Holds the string that was scanned
		text := scanner.Text()
		if len(text) != 0 {
			sb.WriteString(text + "\n")
		} else {
			// exit if user entered an empty string
			break
		}
		fmt.Fprintf(c.StdIO, "%s", color.HiBlackString("| "))
	}

	// handle error
	if scanner.Err() != nil {
		fmt.Println("Error: ", scanner.Err())
	}
	s := sb.String()
	if strings.TrimSpace(s) == "quit" {
		fmt.Println("Bye!")
		os.Exit(0)
	}
	return s
}

func (c Client) Query(ctx context.Context, query string) (*string, error) {
	// if params.OrgID == "" && params.OrgName == "" && c.ActiveConfig.Org == "" {
	// 	return clients.ErrMustSpecifyOrg
	// }
	// orgId, err := params.GetOrgID(ctx, c.ActiveConfig, c.OrganizationsApi)
	// if err != nil {
	// 	return err
	// }
	var resContentType string
	switch c.Format {
	case CsvFormat:
		resContentType = "application/csv"
	case JsonFormat:
		resContentType = "application/json"
	default:
		return nil, fmt.Errorf("unexpected format: %s", c.Format)
	}

	resBody, err := c.GetLegacyQuery(ctx).
		U(c.U).
		P(c.P).
		Db(c.Db).
		Q(query).
		Rp(c.Rp).
		Epoch(c.Epoch).
		Accept(resContentType).
		Execute()
	if err != nil {
		return nil, err
	}
	return &resBody, nil
}

func (c Client) OutputCsv(csvBody *string) {
	fmt.Fprintf(c.CLI.StdIO, *csvBody)
}

func (c Client) OutputJson(jsonBody *string) {
	fmt.Fprintf(c.CLI.StdIO, *jsonBody)
}
