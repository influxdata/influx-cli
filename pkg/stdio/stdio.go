package stdio

import "io"

const MinPasswordLen = 8

type StdIO interface {
	io.Writer
	WriteErr(p []byte) (n int, err error)
	Banner(message string) error
	Error(message string) error
	GetStringInput(prompt, defaultValue string) (string, error)
	GetSecret(prompt string, minLen int) (string, error)
	GetPassword(prompt string) (string, error)
	GetConfirm(prompt string) bool
}
