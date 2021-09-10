package middleware_test

import (
	"bytes"
	"fmt"
	"io"
	"testing"

	"github.com/influxdata/influx-cli/v2/pkg/cli/middleware"
	"github.com/stretchr/testify/require"
	"github.com/urfave/cli"
)

func Test_NoArgs(t *testing.T) {
	setup := func(out io.Writer, withCommand bool, withSubcommand bool) *cli.App {
		app := cli.App{
			Name:   "test",
			Before: middleware.NoArgs,
		}
		flags := []cli.Flag{
			cli.StringFlag{
				Name:     "hello",
				Required: true,
			},
		}
		action := func(ctx *cli.Context) error {
			_, err := fmt.Fprintf(out, "Hello, %s!", ctx.String("hello"))
			return err
		}
		if withCommand {
			app.Commands = []cli.Command{
				{
					Name:   "command",
					Before: middleware.NoArgs,
				},
			}
			if withSubcommand {
				app.Commands[0].Subcommands = []cli.Command{
					{
						Name:   "subcommand",
						Before: middleware.NoArgs,
						Flags:  flags,
						Action: action,
					},
				}
			} else {
				app.Commands[0].Flags = flags
				app.Commands[0].Action = action
			}
		} else {
			app.Flags = flags
			app.Action = action
		}
		return &app
	}

	testCases := []struct {
		name       string
		cli        []string
		command    bool
		subcommand bool
		expectOut  string
		expectErr  string
	}{
		{
			name:      "top-level without args",
			cli:       []string{"--hello", "world"},
			expectOut: "Hello, world!",
		},
		{
			name:      "top-level with args",
			cli:       []string{"--hello", "world", "etc"},
			expectErr: `unknown command "etc" for "test"`,
		},
		{
			name:      "command without args",
			cli:       []string{"command", "--hello", "world"},
			command:   true,
			expectOut: "Hello, world!",
		},
		{
			name:      "command with args",
			cli:       []string{"command", "--hello", "world", "etc"},
			command:   true,
			expectErr: `unknown command "etc" for "command"`,
		},
		{
			name:       "subcommand without args",
			cli:        []string{"command", "subcommand", "--hello", "world"},
			subcommand: true,
			expectOut:  "Hello, world!",
		},
		{
			name:       "subcommand with args",
			cli:        []string{"command", "subcommand", "--hello", "world", "etc"},
			subcommand: true,
			expectErr:  `unknown command "etc" for "subcommand"`,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			var out bytes.Buffer
			cmd := setup(&out, tc.command || tc.subcommand, tc.subcommand)
			err := cmd.Run(append([]string{"test"}, tc.cli...))
			if tc.expectErr != "" {
				require.Error(t, err)
				require.EqualError(t, err, tc.expectErr)
				return
			}
			require.NoError(t, err)
			require.Equal(t, tc.expectOut, out.String())
		})
	}
}
