package stdio

import (
	"errors"
	"io"

	"github.com/AlecAivazis/survey/v2"
	"github.com/AlecAivazis/survey/v2/terminal"
)

// interactiveStdio interacts with the user via an interactive terminal.
type interactiveStdio struct {
	in  terminal.FileReader
	out terminal.FileWriter
	err io.Writer
}

func (t *interactiveStdio) Write(p []byte) (int, error) {
	return t.out.Write(p)
}

func (t *interactiveStdio) WriteErr(p []byte) (int, error) {
	return t.err.Write(p)
}

type bannerTemplateData struct {
	Message string
}

var bannerTemplate = `{{color "cyan+hb"}}> {{ .Message }}{{color "reset"}}
`

func (t *interactiveStdio) Banner(message string) error {
	r := survey.Renderer{}
	r.WithStdio(terminal.Stdio{In: t.in, Out: t.out, Err: t.err})
	return r.Render(bannerTemplate, &bannerTemplateData{Message: message})
}

func (t *interactiveStdio) Error(message string) error {
	r := survey.Renderer{}
	r.WithStdio(terminal.Stdio{In: t.in, Out: t.out, Err: t.err})
	cfg := survey.PromptConfig{Icons: survey.IconSet{Error: survey.Icon{Text: "X", Format: "red"}}}
	return r.Error(&cfg, errors.New(message))
}

func (t *interactiveStdio) IsInteractive() bool {
	return true
}

func (t *interactiveStdio) GetStringInput(prompt, defaultValue string) (input string, err error) {
	question := survey.Input{
		Message: prompt,
		Default: defaultValue,
	}
	err = survey.AskOne(&question, &input,
		survey.WithStdio(t.in, t.out, t.err),
		survey.WithValidator(survey.Required))
	return
}

func (t *interactiveStdio) GetSecret(prompt string, minLen int) (password string, err error) {
	question := survey.Password{Message: prompt}
	opts := []survey.AskOpt{survey.WithStdio(t.in, t.out, t.err)}
	if minLen > 0 {
		opts = append(opts, survey.WithValidator(survey.MinLength(minLen)))
	}
	err = survey.AskOne(&question, &password, opts...)
	question.NewCursor().HorizontalAbsolute(0)
	return
}

func (t *interactiveStdio) GetPassword(prompt string) (string, error) {
	for {
		pass1, err := t.GetSecret(prompt, MinPasswordLen)
		if err != nil {
			return "", err
		}
		// Don't bother with the length check the 2nd time, since we check equality to pass1.
		pass2, err := t.GetSecret(prompt+" again", 0)
		if err != nil {
			return "", err
		}
		if pass1 == pass2 {
			return pass1, nil
		}
		if err := t.Error("Passwords do not match"); err != nil {
			return "", err
		}
	}
}

func (t *interactiveStdio) GetConfirm(prompt string) (answer bool) {
	question := survey.Confirm{
		Message: prompt,
	}
	if err := survey.AskOne(&question, &answer, survey.WithStdio(t.in, t.out, t.err)); err != nil {
		answer = false
	}
	return
}
