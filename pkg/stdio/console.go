package stdio

import (
	"errors"
	"io"
	"os"

	"github.com/AlecAivazis/survey/v2"
	"github.com/AlecAivazis/survey/v2/terminal"
)

// terminalStdio interacts with the user via an interactive terminal.
type terminalStdio struct {
	Stdin  terminal.FileReader
	Stdout terminal.FileWriter
	Stderr io.Writer
}

// TerminalStdio interacts with users over stdin/stdout/stderr.
var TerminalStdio StdIO = &terminalStdio{
	Stdin:  os.Stdin,
	Stdout: os.Stdout,
	Stderr: os.Stderr,
}

// Write prints some bytes to stdout.
func (t *terminalStdio) Write(p []byte) (int, error) {
	return t.Stdout.Write(p)
}

// WriteErr prints some bytes to stderr.
func (t *terminalStdio) WriteErr(p []byte) (int, error) {
	return t.Stderr.Write(p)
}

type bannerTemplateData struct {
	Message string
}

var bannerTemplate = `{{color "cyan+hb"}}> {{ .Message }}{{color "reset"}}
`

// Banner displays informational text to the user.
func (t *terminalStdio) Banner(message string) error {
	r := survey.Renderer{}
	r.WithStdio(terminal.Stdio{In: t.Stdin, Out: t.Stdout, Err: t.Stderr})
	return r.Render(bannerTemplate, &bannerTemplateData{Message: message})
}

// Error displays an error message to the user.
func (t *terminalStdio) Error(message string) error {
	r := survey.Renderer{}
	r.WithStdio(terminal.Stdio{In: t.Stdin, Out: t.Stdout, Err: t.Stderr})
	cfg := survey.PromptConfig{Icons: survey.IconSet{Error: survey.Icon{Text: "X", Format: "red"}}}
	return r.Error(&cfg, errors.New(message))
}

// GetStringInput prompts the user for arbitrary input.
func (t *terminalStdio) GetStringInput(prompt, defaultValue string) (input string, err error) {
	question := survey.Input{
		Message: prompt,
		Default: defaultValue,
	}
	err = survey.AskOne(&question, &input,
		survey.WithStdio(t.Stdin, t.Stdout, t.Stderr),
		survey.WithValidator(survey.Required))
	return
}

// GetSecret prompts the user for a password.
func (t *terminalStdio) GetSecret(prompt string, minLen int) (password string, err error) {
	question := survey.Password{Message: prompt}
	opts := []survey.AskOpt{survey.WithStdio(t.Stdin, t.Stdout, t.Stderr)}
	if minLen > 0 {
		opts = append(opts, survey.WithValidator(survey.MinLength(minLen)))
	}
	err = survey.AskOne(&question, &password, opts...)
	question.NewCursor().HorizontalAbsolute(0)
	return
}

// GetConfirm asks the user for a y/n answer to a prompt.
func (t *terminalStdio) GetConfirm(prompt string) (answer bool) {
	question := survey.Confirm{
		Message: prompt,
	}
	if err := survey.AskOne(&question, &answer, survey.WithStdio(t.Stdin, t.Stdout, t.Stderr)); err != nil {
		answer = false
	}
	return
}
