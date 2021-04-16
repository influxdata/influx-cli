package internal

import (
	"io"
)

type CLI struct {
	Stdin  io.Reader
	Stdout io.Writer
	Stderr io.Writer

	TraceId string
}
