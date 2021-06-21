package stdio

import "io"

type StdIO interface {
	io.Writer
	WriteErr(p []byte) (n int, err error)
	Banner(message string) error
	Error(message string) error
	GetStringInput(prompt, defaultValue string) (string, error)
	GetSecret(prompt string, minLen int) (string, error)
	GetConfirm(prompt string) bool
}
