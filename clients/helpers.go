package clients

import (
	"fmt"
	"io"
	"os"
	"strings"
)

// ReadQuery reads a Flux query into memory from a --file argument, args, or stdin
func ReadQuery(filepath string, args []string) (string, error) {
	nargs := len(args)

	if nargs > 1 {
		return "", fmt.Errorf("at most 1 query string can be specified as an argument, got %d", nargs)
	}
	if nargs == 1 && filepath != "" {
		return "", fmt.Errorf("query can be specified as a CLI arg or passed in a file via flag, not both")
	}

	readFile := func(path string) (string, error) {
		queryBytes, err := os.ReadFile(path)
		if err != nil {
			return "", fmt.Errorf("failed to read query from %q: %w", path, err)
		}
		return string(queryBytes), nil
	}

	readStdin := func() (string, error) {
		queryBytes, err := io.ReadAll(os.Stdin)
		if err != nil {
			return "", fmt.Errorf("failed to read query from stdin: %w", err)
		}
		return string(queryBytes), err
	}

	if filepath != "" {
		return readFile(filepath)
	}
	if nargs == 0 {
		return readStdin()
	}

	arg := args[0]
	// Backwards compatibility.
	if strings.HasPrefix(arg, "@") {
		return readFile(arg[1:])
	} else if arg == "-" {
		return readStdin()
	} else {
		return arg, nil
	}
}
