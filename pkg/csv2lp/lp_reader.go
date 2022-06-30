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
}

// LineProtocolFilter creates a reader wrapper that parses points, skipping if invalid
func LineProtocolFilter(reader io.Reader) *LineProtocolFilterReader {
	lineReader := NewLineReader(reader)
	lineReader.LineNumber = 1 // start counting from 1
	return &LineProtocolFilterReader{
		lineReader: lineReader,
	}
}

func (state *LineProtocolFilterReader) Read(b []byte) (int, error) {
	for {
		// buf := make([]byte, len(b))
		bytesRead, err := state.lineReader.Read(b)
		if err != nil {
			return bytesRead, err
		}
		state.LineNumber = state.lineReader.LastLineNumber
		buf := b[0:bytesRead]
		pts, err := models.ParsePoints(buf) // any time precision because we won't actually use this point
		if err != nil {
			log.Printf("invalid point on line %d: %v\n", state.LineNumber, err)
			continue
		} else if len(pts) == 0 { // no points on this line
			continue
		}
		return bytesRead, nil
	}
}
