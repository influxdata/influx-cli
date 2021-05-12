package query

import (
	"io"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/influxdata/influx-cli/v2/pkg/fluxcsv"
)

// formattingPrinter formats query results into a structured table before printing.
type formattingPrinter struct {
	widths    []int
	maxWidth  int
	newWidths []int
	pad       []byte
	dash      []byte
	// fmtBuf is used to format values
	fmtBuf [64]byte

	cols       []fluxcsv.FluxColumn
	lastColIdx int
}

func NewFormattingPrinter() *formattingPrinter {
	return &formattingPrinter{}
}

func (f *formattingPrinter) PrintQueryResults(resultStream io.ReadCloser, out io.Writer) error {
	res := fluxcsv.NewQueryTableResult(resultStream)
	defer res.Close()
	return f.write(res, out)
}

const fixedWidthTimeFmt = "2006-01-02T15:04:05.000000000Z"

var eol = []byte{'\n'}

type writeHelper struct {
	w   io.Writer
	err error
}

func (w *writeHelper) write(data []byte) {
	if w.err != nil {
		return
	}
	_, err := w.w.Write(data)
	w.err = err
}

var minWidthsByType = map[fluxcsv.ColType]int{
	fluxcsv.BoolDatatype:        12,
	fluxcsv.LongDatatype:        26,
	fluxcsv.ULongDatatype:       27,
	fluxcsv.DoubleDatatype:      28,
	fluxcsv.StringDatatype:      22,
	fluxcsv.TimeDatatypeRFC:     len(fixedWidthTimeFmt),
	fluxcsv.TimeDatatypeRFCNano: len(fixedWidthTimeFmt),
}

// write writes the formatted table data to w.
func (f *formattingPrinter) write(res *fluxcsv.QueryTableResult, out io.Writer) error {
	w := &writeHelper{w: out}

	r := 0
	for res.Next() {
		record := res.Record()

		if res.AnnotationsChanged() {
			// Reset and sort cols
			f.cols = res.Metadata().Columns()
			f.lastColIdx = len(f.cols) - 1
			groupKeys := make(map[string]int, len(res.Metadata().GroupKeyCols()))
			for i, k := range res.Metadata().GroupKeyCols() {
				groupKeys[k] = i
			}
			sort.Slice(f.cols, func(i, j int) bool {
				iCol, jCol := f.cols[i], f.cols[j]
				iGroupIdx, iIsGroup := groupKeys[iCol.Name()]
				jGroupIdx, jIsGroup := groupKeys[jCol.Name()]

				if iIsGroup && jIsGroup {
					return iGroupIdx < jGroupIdx
				}
				if !iIsGroup && !jIsGroup {
					return i < j
				}
				return iIsGroup && !jIsGroup
			})

			// Compute header widths
			f.widths = make([]int, len(f.cols))
			for i, c := range f.cols {
				// Column header is "<label>:<type>"
				l := len(c.Name()) + len(display(c.DataType())) + 1
				min := minWidthsByType[c.DataType()]
				if min > l {
					l = min
				}
				f.widths[i] = l
				if l > f.maxWidth {
					f.maxWidth = l
				}
			}
		}

		if res.ResultChanged() {
			w.write([]byte("Result: "))
			w.write([]byte(record.Result()))
			w.write(eol)
		}
		if res.TableIdChanged() || res.AnnotationsChanged() {
			w.write([]byte("Table: keys: ["))
			labels := make([]string, len(res.Metadata().GroupKeyCols()))
			for i, c := range res.Metadata().GroupKeyCols() {
				labels[i] = c
			}
			w.write([]byte(strings.Join(labels, ", ")))
			w.write([]byte("]"))
			w.write(eol)

			// Check err and return early
			if w.err != nil {
				return w.err
			}

			r = 0
		}

		if r == 0 {
			for i, c := range f.cols {
				buf := f.valueBuf(c.DataType(), record.ValueByKey(c.Name()))
				l := len(buf)
				if l > f.widths[i] {
					f.widths[i] = l
				}
				if l > f.maxWidth {
					f.maxWidth = l
				}
			}
			f.makePaddingBuffers()
			f.writeHeader(w)
			f.writeHeaderSeparator(w)
			f.newWidths = make([]int, len(f.widths))
			copy(f.newWidths, f.widths)
		}
		for i, c := range f.cols {
			buf := f.valueBuf(c.DataType(), record.ValueByKey(c.Name()))
			l := len(buf)
			padding := f.widths[i] - l
			if padding >= 0 {
				w.write(f.pad[:padding])
				w.write(buf)
			} else {
				//TODO make unicode friendly
				w.write(buf[:f.widths[i]-3])
				w.write([]byte{'.', '.', '.'})
			}
			if i != f.lastColIdx {
				w.write(f.pad[:2])
			}
			if l > f.newWidths[i] {
				f.newWidths[i] = l
			}
			if l > f.maxWidth {
				f.maxWidth = l
			}
		}
		w.write(eol)
		r++
	}
	return w.err
}

func (f *formattingPrinter) makePaddingBuffers() {
	if len(f.pad) != f.maxWidth {
		f.pad = make([]byte, f.maxWidth)
		for i := range f.pad {
			f.pad[i] = ' '
		}
	}
	if len(f.dash) != f.maxWidth {
		f.dash = make([]byte, f.maxWidth)
		for i := range f.dash {
			f.dash[i] = '-'
		}
	}
}

func (f *formattingPrinter) writeHeader(w *writeHelper) {
	for i, c := range f.cols {
		buf := append(append([]byte(c.Name()), ':'), []byte(display(c.DataType()))...)
		w.write(f.pad[:f.widths[i]-len(buf)])
		w.write(buf)
		if i != f.lastColIdx {
			w.write(f.pad[:2])
		}
	}
	w.write(eol)
}

func (f *formattingPrinter) writeHeaderSeparator(w *writeHelper) {
	for i := range f.cols {
		w.write(f.dash[:f.widths[i]])
		if i != f.lastColIdx {
			w.write(f.pad[:2])
		}
	}
	w.write(eol)
}

func display(t fluxcsv.ColType) string {
	switch t {
	case fluxcsv.StringDatatype:
		return "string"
	case fluxcsv.DoubleDatatype:
		return "float"
	case fluxcsv.BoolDatatype:
		return "boolean"
	case fluxcsv.LongDatatype:
		return "int"
	case fluxcsv.Base64BinaryDataType:
		return "base64Binary"
	case fluxcsv.TimeDatatypeRFC:
		fallthrough
	case fluxcsv.TimeDatatypeRFCNano:
		return "time"
	// TODO: These weren't implemented in the influxdb CLI.
	//case fluxcsv.ULongDatatype:
	//case fluxcsv.DurationDatatype:
	default:
		panic("shouldn't happen")
	}
}

func (f *formattingPrinter) valueBuf(typ fluxcsv.ColType, v interface{}) []byte {
	var buf []byte
	if v == nil {
		return buf
	}
	switch typ {
	case fluxcsv.StringDatatype:
		buf = []byte(v.(string))
	case fluxcsv.DoubleDatatype:
		buf = strconv.AppendFloat(f.fmtBuf[0:0], v.(float64), 'f', -1, 64)
	case fluxcsv.BoolDatatype:
		buf = strconv.AppendBool(f.fmtBuf[0:0], v.(bool))
	case fluxcsv.LongDatatype:
		buf = strconv.AppendInt(f.fmtBuf[0:0], v.(int64), 10)
	case fluxcsv.ULongDatatype:
		buf = strconv.AppendUint(f.fmtBuf[0:0], v.(uint64), 10)
	case fluxcsv.TimeDatatypeRFC:
		fallthrough
	case fluxcsv.TimeDatatypeRFCNano:
		buf = []byte(v.(time.Time).Format(fixedWidthTimeFmt))
		// TODO: These weren't implemented in the influxdb CLI.
		//case fluxcsv.DurationDatatype:
		//case fluxcsv.Base64BinaryDataType:
	}
	return buf
}
