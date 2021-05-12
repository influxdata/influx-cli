package fluxcsv

import (
	"encoding/base64"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"
	"time"

	"github.com/influxdata/influx-cli/v2/internal/duration"
)

type ColType int

const (
	StringDatatype ColType = iota
	DoubleDatatype
	BoolDatatype
	LongDatatype
	ULongDatatype
	DurationDatatype
	Base64BinaryDataType
	TimeDatatypeRFC
	TimeDatatypeRFCNano
	InvalidDatatype
)

func ParseType(s string) (ColType, error) {
	switch s {
	case "string":
		return StringDatatype, nil
	case "double":
		return DoubleDatatype, nil
	case "boolean":
		return BoolDatatype, nil
	case "long":
		return LongDatatype, nil
	case "unsignedLong":
		return ULongDatatype, nil
	case "duration":
		return DurationDatatype, nil
	case "base64Binary":
		return Base64BinaryDataType, nil
	case "dateTime:RFC3339":
		return TimeDatatypeRFC, nil
	case "dateTime:RFC3339Nano":
		return TimeDatatypeRFCNano, nil
	default:
		return InvalidDatatype, fmt.Errorf("unknown data type %s", s)
	}
}

// QueryTableResult parses streamed flux query response into structures representing flux table parts
// Walking though the result is done by repeatedly calling Next() until returns false.
// Actual flux table info (columns with names, data types, etc) is returned by TableMetadata() method.
// Data are acquired by Record() method.
// Preliminary end can be caused by an error, so when Next() return false, check Err() for an error
type QueryTableResult struct {
	io.Closer
	csvReader *csv.Reader

	resultChanged      bool
	tableIdChanged     bool
	annotationsChanged bool

	record *FluxRecord

	baseColumns []*FluxColumn
	metadata    *FluxTableMetadata
	err         error
}

func NewQueryTableResult(rawResponse io.ReadCloser) *QueryTableResult {
	csvReader := csv.NewReader(rawResponse)
	csvReader.FieldsPerRecord = -1
	return &QueryTableResult{Closer: rawResponse, csvReader: csvReader}
}

// ResultChanged returns true if the last call of Next() found a new query result
func (q *QueryTableResult) ResultChanged() bool {
	return q.resultChanged
}

// TableIdChanged returns true if the last call of Next() found a new table within the query result
func (q *QueryTableResult) TableIdChanged() bool {
	return q.tableIdChanged
}

// AnnotationsChanged returns true if last call of Next() found new CSV annotations
func (q *QueryTableResult) AnnotationsChanged() bool {
	return q.annotationsChanged
}

// Record returns last parsed flux table data row
// Use Record methods to access value and row properties
func (q *QueryTableResult) Record() *FluxRecord {
	return q.record
}

// Metadata returns table-level info for last parsed flux table data row
func (q *QueryTableResult) Metadata() *FluxTableMetadata {
	return q.metadata
}

type parsingState int

const (
	parsingStateNormal parsingState = iota
	parsingStateAnnotation
	parsingStateNameRow
	parsingStateError
)

// Next advances to next row in query result.
// During the first time it is called, Next creates also table metadata
// Actual parsed row is available through Record() function
// Returns false in case of end or an error, otherwise true
func (q *QueryTableResult) Next() bool {
	var row []string
	// set closing query in case of preliminary return
	closer := func() {
		if err := q.Close(); err != nil {
			message := err.Error()
			if q.err != nil {
				message = fmt.Sprintf("%s,%s", message, q.err.Error())
			}
			q.err = errors.New(message)
		}
	}
	defer func() {
		closer()
	}()
	parsingState := parsingStateNormal
	q.annotationsChanged = false
	dataTypeAnnotationFound := false

readRow:
	row, q.err = q.csvReader.Read()
	if q.err == io.EOF {
		q.err = nil
		return false
	}
	if q.err != nil {
		return false
	}

	if len(row) <= 1 {
		goto readRow
	}
	if len(row[0]) > 0 && row[0][0] == '#' {
		if parsingState == parsingStateNormal {
			q.annotationsChanged = true
			q.baseColumns = nil
			for range row[1:] {
				q.baseColumns = append(q.baseColumns, NewFluxColumn())
			}
			parsingState = parsingStateAnnotation
		}
	}
	expectedNcol := len(q.baseColumns)
	if expectedNcol == 0 {
		q.err = errors.New("parsing error, annotations not found")
		return false
	}
	ncol := len(row) - 1
	if ncol != expectedNcol {
		q.err = fmt.Errorf("parsing error, row has different number of columns than the table: %d vs %d", ncol, expectedNcol)
		return false
	}
	switch row[0] {
	case "":
		switch parsingState {
		case parsingStateAnnotation:
			if !dataTypeAnnotationFound {
				q.err = errors.New("parsing error, datatype annotation not found")
				return false
			}
			parsingState = parsingStateNameRow
			fallthrough
		case parsingStateNameRow:
			if row[1] == "error" {
				parsingState = parsingStateError
			} else {
				for i, n := range row[1:] {
					q.baseColumns[i].SetName(n)
				}
				q.metadata = NewFluxTableMetadataFull(q.baseColumns...)
				parsingState = parsingStateNormal
			}
			goto readRow
		case parsingStateError:
			var message string
			if len(row) > 1 && len(row[1]) > 0 {
				message = row[1]
			} else {
				message = "unknown query error"
			}
			reference := ""
			if len(row) > 2 && len(row[2]) > 0 {
				reference = fmt.Sprintf(",%s", row[2])
			}
			q.err = fmt.Errorf("%s%s", message, reference)
			return false
		}
		values := make(map[string]interface{})
		for i, v := range row[1:] {
			values[q.baseColumns[i].Name()], q.err = toValue(
				stringTernary(v, q.baseColumns[i].DefaultValue()), q.baseColumns[i].DataType(), q.baseColumns[i].Name())
			if q.err != nil {
				return false
			}
		}
		var prevRes string
		var prevId int64
		if q.record != nil {
			prevRes, prevId = q.record.Result(), q.record.TableId()
		}
		q.record, q.err = NewFluxRecord(q.metadata, values)
		q.resultChanged = q.record.Result() != prevRes
		q.tableIdChanged = q.record.TableId() != prevId
		if q.err != nil {
			return false
		}
	case "#datatype":
		dataTypeAnnotationFound = true
		for i, d := range row[1:] {
			t, err := ParseType(d)
			if err != nil {
				q.err = err
				return false
			}
			q.baseColumns[i].SetDataType(t)
		}
		goto readRow
	case "#group":
		for i, g := range row[1:] {
			q.baseColumns[i].SetGroup(g == "true")
		}
		goto readRow
	case "#default":
		for i, c := range row[1:] {
			q.baseColumns[i].SetDefaultValue(c)
		}
		goto readRow
	}
	// don't close query
	closer = func() {}
	return true
}

// Err returns an error raised during flux query response parsing
func (q *QueryTableResult) Err() error {
	return q.err
}

// stringTernary returns a if not empty, otherwise b
func stringTernary(a, b string) string {
	if a == "" {
		return b
	}
	return a
}

// toValue converts s into type by t
func toValue(s string, t ColType, name string) (interface{}, error) {
	if s == "" {
		return nil, nil
	}
	switch t {
	case StringDatatype:
		return s, nil
	case TimeDatatypeRFC:
		return time.Parse(time.RFC3339, s)
	case TimeDatatypeRFCNano:
		return time.Parse(time.RFC3339Nano, s)
	case DurationDatatype:
		return duration.RawDurationToTimeDuration(s)
	case DoubleDatatype:
		return strconv.ParseFloat(s, 64)
	case BoolDatatype:
		if strings.ToLower(s) == "false" {
			return false, nil
		}
		return true, nil
	case LongDatatype:
		return strconv.ParseInt(s, 10, 64)
	case ULongDatatype:
		return strconv.ParseUint(s, 10, 64)
	case Base64BinaryDataType:
		return base64.StdEncoding.DecodeString(s)
	default:
		return nil, fmt.Errorf("%s has unknown data type %v", name, t)
	}
}
