package stdio

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

// noninteractiveStdio interacts with the user via an interactive terminal.
type noninteractiveStdio struct {
	in  *bufio.Scanner
	out io.Writer
	err io.Writer
}

func (t *noninteractiveStdio) Write(p []byte) (int, error) {
	return t.out.Write(p)
}

func (t *noninteractiveStdio) WriteErr(p []byte) (int, error) {
	return t.err.Write(p)
}

func (t *noninteractiveStdio) Banner(message string) error {
	_, err := fmt.Fprintf(t.out, "> %s\n", message)
	return err
}

func (t *noninteractiveStdio) Error(message string) error {
	_, err := fmt.Fprintf(t.err, "X %s\n", message)
	return err
}

func (t *noninteractiveStdio) IsInteractive() bool {
	return false
}

func (t *noninteractiveStdio) GetStringInput(prompt, defaultValue string) (string, error) {
	if inLine := t.readStdinLine(); inLine != "" {
		return inLine, nil
	}
	if defaultValue != "" {
		return defaultValue, nil
	}
	return "", fmt.Errorf("couldn't get input for prompt %q: no data on stdin", prompt)
}

func (t *noninteractiveStdio) GetSecret(prompt string, minLen int) (string, error) {
	inLine := t.readStdinLine()
	if len(inLine) >= minLen {
		return inLine, nil
	} else if minLen > 0 {
		return "", fmt.Errorf("value for prompt %q is too short: min length is %d", prompt, minLen)
	}
	return "", nil
}

func (t *noninteractiveStdio) GetPassword(prompt string) (string, error) {
	return t.GetSecret(prompt, MinPasswordLen)
}

func (t *noninteractiveStdio) GetConfirm(prompt string) (answer bool) {
	return strings.HasPrefix(t.readStdinLine(), "y")
}

// readStdinLine returns the first line of text on stdin, or empty string if stdin is at EOF.
func (t *noninteractiveStdio) readStdinLine() string {
	if !t.in.Scan() {
		return ""
	}
	return t.in.Text()
}
