package stdio

import (
	"bufio"
	"io"
	"os"

	"github.com/AlecAivazis/survey/v2/terminal"
	"github.com/mattn/go-isatty"
)

// Disable password length checking to let influxdb handle it
const MinPasswordLen = 0

type StdIO interface {
	// Write prints some bytes to stdout.
	Write(p []byte) (n int, err error)
	// WriteErr prints some bytes to stderr.
	WriteErr(p []byte) (n int, err error)
	// Banner displays informational text to the user.
	Banner(message string) error
	// Error displays an error message to the user.
	Error(message string) error
	// IsInteractive signals whether interactive I/O is supported.
	IsInteractive() bool
	// GetStringInput prompts the user for arbitrary input.
	GetStringInput(prompt, defaultValue string) (string, error)
	// GetSecret prompts the user for a secret.
	GetSecret(prompt string, minLen int) (string, error)
	// GetPassword prompts the user for a secret twice, and inputs must match.
	// Uses stdio.MinPasswordLen as the minimum input length
	GetPassword(prompt string) (string, error)
	// GetConfirm asks the user for a y/n answer to a prompt.
	GetConfirm(prompt string) bool
}

func newTerminalStdio(in terminal.FileReader, out terminal.FileWriter, err io.Writer) StdIO {
	interactiveIn := isatty.IsTerminal(in.Fd()) || isatty.IsCygwinTerminal(in.Fd())
	interactiveOut := isatty.IsTerminal(out.Fd()) || isatty.IsCygwinTerminal(out.Fd())

	if interactiveIn && interactiveOut {
		return &interactiveStdio{in: in, out: out, err: err}
	}

	return &noninteractiveStdio{
		in:  bufio.NewScanner(in),
		out: out,
		err: err,
	}
}

// TerminalStdio interacts with users over stdin/stdout/stderr.
var TerminalStdio = newTerminalStdio(os.Stdin, os.Stdout, os.Stderr)
