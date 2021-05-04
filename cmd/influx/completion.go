package main

import (
	"fmt"
	"os"
	"path"

	"github.com/urfave/cli/v2"
)

var (
	// Copied from urfave/cli, see:
	// https://github.com/urfave/cli/blob/master/autocomplete/bash_autocomplete
	bashPattern = `#!/usr/bin/env bash

_cli_bash_autocomplete () {
    if [[ "${COMP_WORDS[0]}" != "source" ]]; then
        local cur opts base
        COMPREPLY=()
        cur="${COMP_WORDS[COMP_CWORD]}"
        if [[ "$cur" == "-"* ]]; then
            opts=$( ${COMP_WORDS[@]:0:$COMP_CWORD} ${cur} --%[2]s )
        else
            opts=$( ${COMP_WORDS[@]:0:$COMP_CWORD} --%[2]s )
        fi
        COMPREPLY=( $(compgen -W "${opts}" -- ${cur}) )
        return 0
    fi
}

complete -o bashdefault -o default -o nospace -F _cli_bash_autocomplete %[1]s
`
	// Copied from urfave/cli, see:
	// https://github.com/urfave/cli/blob/master/autocomplete/zsh_autocomplete
	zshPattern = `#compdef %[1]s

_cli_zsh_autocomplete() {
    local -a opts
    local cur
    cur=${words[-1]}
    if [[ "$cur" == "-"* ]]; then
        opts=("${(@f)$(_CLI_ZSH_AUTOCOMPLETE_HACK=1 ${words[@]:0:#words[@]-1} ${cur} --%[2]s)}")
    else
        opts=("${(@f)$(_CLI_ZSH_AUTOCOMPLETE_HACK=1 ${words[@]:0:#words[@]-1} --%[2]s)}")
    fi

    if [[ "${opts[1]}" != "" ]]; then
        _describe 'values' opts
    else
        _files
    fi

    return
}

compdef _cli_zsh_autocomplete %[1]s
`
	// Copied from urfave/cli, see:
	// https://github.com/urfave/cli/blob/master/autocomplete/powershell_autocomplete.ps1
	ps1Pattern = `Register-ArgumentCompleter -Native -CommandName %[1]s -ScriptBlock {
param($commandName, $wordToComplete, $cursorPosition)
     $other = "$wordToComplete --%[2]s"
         Invoke-Expression $other | ForEach-Object {
            [System.Management.Automation.CompletionResult]::new($_, $_, 'ParameterValue', $_)
         }
 }
`
)

func newCompletionCmd() *cli.Command {
	return &cli.Command{
		Name:      "completion",
		Usage:     "Generates completion scripts",
		ArgsUsage: "[bash|zsh|powershell]",
		Action: func(ctx *cli.Context) error {
			prog := path.Base(os.Args[0])
			completeFlag := cli.BashCompletionFlag.Names()[0]

			if ctx.NArg() != 1 {
				return fmt.Errorf("usage: %s completion [bash|zsh|powershell]", prog)
			}
			shellString := ctx.Args().Get(0)
			var script string
			switch shellString {
			case "bash":
				script = bashPattern
			case "zsh":
				script = zshPattern
			case "powershell":
				script = ps1Pattern
			default:
				return fmt.Errorf("usage: %s completion [bash|zsh|powershell]", prog)
			}

			_, err := fmt.Printf(script, prog, completeFlag)
			return err
		},
	}
}
