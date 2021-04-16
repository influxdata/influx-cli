package internal

import (
	"errors"
	"io"

	"github.com/AlecAivazis/survey/v2"
	"github.com/AlecAivazis/survey/v2/terminal"
	"github.com/influxdata/influx-cli/v2/internal/config"
)

// CLI is a container for common functionality used to execute commands.
type CLI struct {
	Stdin  terminal.FileReader
	Stdout terminal.FileWriter
	Stderr io.Writer

	TraceId string

	ActiveConfig  config.Config
	ConfigService config.Service
}

// GetStringInput prompts the user for arbitrary input.
func (c *CLI) GetStringInput(prompt, defaultValue string) (input string, err error) {
	question := survey.Input{
		Message: prompt,
		Default: defaultValue,
	}
	err = survey.AskOne(&question, &input,
		survey.WithStdio(c.Stdin, c.Stdout, c.Stderr),
		survey.WithValidator(survey.Required))
	return
}

const MinPasswordLen = 8

type bannerTemplateData struct {
	Message string
}

var bannerTemplate = `{{color "cyan+hb"}}> {{ .Message }}{{color "reset"}}
`

func (c *CLI) Banner(message string) error {
	r := survey.Renderer{}
	r.WithStdio(terminal.Stdio{In: c.Stdin, Out: c.Stdout, Err: c.Stderr})
	return r.Render(bannerTemplate, &bannerTemplateData{Message: message})
}

func (c *CLI) Error(message string) error {
	r := survey.Renderer{}
	r.WithStdio(terminal.Stdio{In: c.Stdin, Out: c.Stdout, Err: c.Stderr})
	cfg := survey.PromptConfig{Icons: survey.IconSet{Error: survey.Icon{Text: "X", Format: "red"}}}
	return r.Error(&cfg, errors.New(message))
}

// GetPassword prompts the user for a password.
func (c *CLI) GetPassword(prompt string, checkLen bool) (password string, err error) {
	question := survey.Password{Message: prompt}
	opts := []survey.AskOpt{survey.WithStdio(c.Stdin, c.Stdout, c.Stderr)}
	if checkLen {
		opts = append(opts, survey.WithValidator(survey.MinLength(MinPasswordLen)))
	}
	err = survey.AskOne(&question, &password, opts...)
	question.NewCursor().HorizontalAbsolute(0)
	return
}

// GetConfirm asks the user for a y/n answer to a prompt.
func (c *CLI) GetConfirm(prompt string) (answer bool) {
	question := survey.Confirm{
		Message: prompt,
	}
	if err := survey.AskOne(&question, &answer, survey.WithStdio(c.Stdin, c.Stdout, c.Stderr)); err != nil {
		answer = false
	}
	return
}
