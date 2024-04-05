package csv2lp

import (
	"errors"
	"io"
	"log"

	"github.com/influxdata/influxdb/v2/models"
)

// LineProtocolFilterReader wraps a line reader and parses points, skipping if invalid
type LineProtocolFilterReader struct {
	// lineReader is used to report line number of the last read CSV line
	lineReader *LineReader
	// LineNumber represents line number of csv.Reader, 1 is the first
	LineNumber int
	buffer     []byte
	// A slice into buffer that contains the current line
	curLine []byte
}

// LineProtocolFilter creates a reader wrapper that parses points, skipping if invalid
func LineProtocolFilter(reader io.Reader) *LineProtocolFilterReader {
	lineReader := NewLineReader(reader)
	lineReader.LineNumber = 1 // start counting from 1
	return &LineProtocolFilterReader{
		lineReader: lineReader,
		buffer:     make([]byte, defaultBufSize),
	}
}

func (state *LineProtocolFilterReader) Read(b []byte) (int, error) {
	var err error
	var bytesRead int

	state.buffer = state.buffer[:0]
	for {
		if len(state.curLine) > 0 {
			n := copy(b, state.curLine)
			// if we have more data than we can put in b, save it for next read
			state.curLine = state.curLine[n:]
			state.buffer = state.buffer[:0]
			return n, err
		}
		for {
			bytesRead, err = state.lineReader.Read(state.buffer[len(state.buffer):cap(state.buffer)])
			if bytesRead == 0 {
				// real error
				if err != nil && !errors.Is(err, io.EOF) {
					state.curLine = nil
					state.buffer = state.buffer[:0]
					return 0, err
				}
				// EOF
				break
			} else {
				//  bytesRead > 0
				state.buffer = state.buffer[:len(state.buffer)+bytesRead]
				if state.buffer[len(state.buffer)-1] != '\n' && !errors.Is(err, io.EOF) {
					// We have run out of buffer, but we need to read more.
					if cap(state.buffer) == len(state.buffer) {
						state.buffer = append(state.buffer, 0)[:len(state.buffer)]
					}
					// read until we get a full line
					continue
				} else {
					// Got a whole line or an EOF on this read
					break
				}
			}
		}
		state.LineNumber = state.lineReader.LastLineNumber
		state.curLine = state.buffer
		pts, parseErr := models.ParsePoints(state.curLine) // any time precision because we won't actually use this point
		if parseErr != nil {
			log.Printf("invalid point on line %d: %v\n", state.LineNumber, err)
			state.buffer = state.buffer[:0]
			state.curLine = nil
			return 0, nil
		} else if len(pts) == 0 {
			state.curLine = nil
			state.buffer = state.buffer[:0]
			// err should only be nil or io.EOF here
			return 0, err
		}
	}
}
