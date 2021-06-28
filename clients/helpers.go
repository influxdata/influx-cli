package clients

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/urfave/cli"
)

// ReadQuery reads a Flux query into memory from a --file argument, args, or stdin
func ReadQuery(ctx *cli.Context) (string, error) {
	nargs := ctx.NArg()
	file := ctx.String("file")

	if nargs > 1 {
		return "", fmt.Errorf("at most 1 query string can be specified over the CLI, got %d", ctx.NArg())
	}
	if nargs == 1 && file != "" {
		return "", fmt.Errorf("query can be specified via --file or over the CLI, not both")
	}

	readFile := func(path string) (string, error) {
		queryBytes, err := ioutil.ReadFile(path)
		if err != nil {
			return "", fmt.Errorf("failed to read query from %q: %w", path, err)
		}
		return string(queryBytes), nil
	}

	readStdin := func() (string, error) {
		queryBytes, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			return "", fmt.Errorf("failed to read query from stdin: %w", err)
		}
		return string(queryBytes), err
	}

	if file != "" {
		return readFile(file)
	}
	if nargs == 0 {
		return readStdin()
	}

	arg := ctx.Args().Get(0)
	// Backwards compatibility.
	if strings.HasPrefix(arg, "@") {
		return readFile(arg[1:])
	} else if arg == "-" {
		return readStdin()
	} else {
		return arg, nil
	}
}
