package middleware

import (
	"fmt"

	"github.com/urfave/cli"
)

var NoArgs cli.BeforeFunc = func(ctx *cli.Context) error {
	// `Before` funcs get run prior to resolving subcommands from args
	if ctx.NArg() > 0 && ctx.App.Command(ctx.Args()[0]) == nil {
		cmdName := ctx.Command.Name
		if cmdName == "" && ctx.App.Name != "" {
			cmdName = ctx.App.Name
		}
		// Use the same error format as `cobra.NoArgs` for consistency with the old CLI.
		return fmt.Errorf("unknown command %q for %q", ctx.Args()[0], cmdName)
	}
	return nil
}
