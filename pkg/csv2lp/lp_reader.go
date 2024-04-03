package csv2lp

import (
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
	// line buffer
	line []byte
	// wrapparound buffer
	wrapBuf []byte
}

// LineProtocolFilter creates a reader wrapper that parses points, skipping if invalid
func LineProtocolFilter(reader io.Reader) *LineProtocolFilterReader {
	lineReader := NewLineReader(reader)
	lineReader.LineNumber = 1 // start counting from 1
	return &LineProtocolFilterReader{
		lineReader: lineReader,
		line:       make([]byte, 4096),
		wrapBuf:    make([]byte, 4096),
	}
}

func (state *LineProtocolFilterReader) Read(b []byte) (int, error) {
	bytesRead, err := state.lineReader.Read(state.line)
	if err != nil && bytesRead == 0 {
		return 0, err
	}

	//read again when we read a partial line at the end of the line reader buffer
	if bytesRead > 0 && state.line[bytesRead-1] != '\n' {
		wrapBytesRead, _ := state.lineReader.Read(state.wrapBuf[0:])
		if wrapBytesRead > 0 {
			copy(state.line[bytesRead:], state.wrapBuf[0:wrapBytesRead])
			bytesRead = bytesRead + wrapBytesRead
		}
	}

	state.LineNumber = state.lineReader.LastLineNumber
	pts, err := models.ParsePoints(state.line[0:bytesRead]) // any time precision because we won't actually use this point
	if err != nil {
		log.Printf("invalid point on line %d: %v\n", state.LineNumber, err)
		return 0, nil
	} else if len(pts) == 0 {
		return 0, nil
	}

	copy(b, state.line[0:bytesRead])
	return bytesRead, nil
}
