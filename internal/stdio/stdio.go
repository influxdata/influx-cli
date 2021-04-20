package stdio

import "io"

type StdIO interface {
	io.Writer
	Banner(message string) error
	Error(message string) error
	GetStringInput(prompt, defaultValue string) (string, error)
	GetPassword(prompt string, minLen int) (string, error)
	GetConfirm(prompt string) bool
}
