package main

import (
	"fmt"
	"os"

	"github.com/influxdata/influx-cli/v2/internal"
	"github.com/urfave/cli/v2"
)

func newCompletionCmd() *cli.Command {
	return &cli.Command{
		Name:      "completion",
		Usage:     "Generates completion scripts",
		ArgsUsage: "[bash|zsh|powershell]",
		Before:    withCli(),
		Action: func(ctx *cli.Context) error {
			prog := os.Args[0]
			completeFlag := cli.BashCompletionFlag.Names()[0]

			if ctx.NArg() != 1 {
				return fmt.Errorf("usage: %s [bash|zsh|powershell]", prog)
			}
			shellString := ctx.Args().Get(0)
			var shell internal.CompletionShell
			switch shellString {
			case "bash":
				shell = internal.CompletionShellBash
			case "zsh":
				shell = internal.CompletionShellZsh
			case "powershell":
				shell = internal.CompletionShellPowershell
			default:
				return fmt.Errorf("usage: %s [bash|zsh|powershell]", prog)
			}

			return getCLI(ctx).Complete(shell, prog, completeFlag)
		},
	}
}
