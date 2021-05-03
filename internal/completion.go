package internal

import (
	"fmt"
)

type CompletionShell int

const (
	CompletionShellBash CompletionShell = iota
	CompletionShellZsh
	CompletionShellPowershell
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

func (c *CLI) Complete(shell CompletionShell, prog string, bashCompleteFlag string) error {
	var script string
	switch shell {
	case CompletionShellBash:
		script = bashPattern
	case CompletionShellZsh:
		script = zshPattern
	case CompletionShellPowershell:
		script = ps1Pattern
	}

	_, err := c.StdIO.Write([]byte(fmt.Sprintf(script, prog, bashCompleteFlag)))
	return err
}
