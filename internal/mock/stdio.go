package mock

import (
	"bytes"
	"errors"
)

type Stdio struct {
	answers map[string]string
	confirm bool
	out     bytes.Buffer
}

func NewMockStdio(promptAnswers map[string]string, confirm bool) *Stdio {
	return &Stdio{answers: promptAnswers, confirm: confirm, out: bytes.Buffer{}}
}

func (m *Stdio) Write(p []byte) (int, error) {
	return m.out.Write(p)
}

func (m *Stdio) Banner(string) error {
	return nil
}

func (m *Stdio) Error(string) error {
	return nil
}

func (m *Stdio) GetStringInput(prompt, defaultValue string) (string, error) {
	v, ok := m.answers[prompt]
	if !ok {
		return defaultValue, nil
	}
	return v, nil
}

func (m *Stdio) GetPassword(prompt string, minLen int) (string, error) {
	v, ok := m.answers[prompt]
	if !ok {
		return "", errors.New("no password given")
	}
	if len(v) < minLen {
		return "", errors.New("password too short")
	}
	return v, nil
}

func (m *Stdio) GetConfirm(string) bool {
	return m.confirm
}

func (m *Stdio) Stdout() string {
	return m.out.String()
}
