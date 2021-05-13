package fluxcsv

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

const (
	ResultCol  = "result"
	TableIdCol = "table"
)

// FluxTableMetadata holds flux query result table information represented by collection of columns.
// Each new table is introduced by annotations
type FluxTableMetadata struct {
	resultColumn  *FluxColumn
	tableIdColumn *FluxColumn
	columns       []FluxColumn
	groupKeyCols  []string
}

// FluxColumn holds flux query table column properties
type FluxColumn struct {
	name         string
	dataType     ColType
	group        bool
	defaultValue string
}

// FluxRecord represents row in the flux query result table
type FluxRecord struct {
	metadata *FluxTableMetadata
	result   string
	tableId  int64
	values   map[string]interface{}
}

// NewFluxTableMetadataFull creates FluxTableMetadata containing the given columns
func NewFluxTableMetadataFull(columns ...*FluxColumn) *FluxTableMetadata {
	m := FluxTableMetadata{}
	for _, c := range columns {
		switch n := c.Name(); n {
		case ResultCol:
			m.resultColumn = c
		case TableIdCol:
			m.tableIdColumn = c
		default:
			m.columns = append(m.columns, *c)
			if c.IsGroup() {
				m.groupKeyCols = append(m.groupKeyCols, n)
			}
		}
	}

	return &m
}

// ResultColumn returns metadata about the column naming results as
// specified by the query
func (f *FluxTableMetadata) ResultColumn() *FluxColumn {
	return f.resultColumn
}

// TableIdColumn returns metadata about the column tracking table IDs
// within a result
func (f *FluxTableMetadata) TableIdColumn() *FluxColumn {
	return f.tableIdColumn
}

// Columns returns slice of flux query result table
func (f *FluxTableMetadata) Columns() []FluxColumn {
	return f.columns
}

// Column returns flux table column by index.
// Returns nil if index is out of the bounds.
func (f *FluxTableMetadata) Column(index int) *FluxColumn {
	if len(f.columns) == 0 || index < 0 || index >= len(f.columns) {
		return nil
	}
	return &f.columns[index]
}

// GroupKeyCols returns the names of the grouping columns
// in the table, sorted in ascending order.
func (f *FluxTableMetadata) GroupKeyCols() []string {
	return f.groupKeyCols
}

// String returns FluxTableMetadata string dump
func (f *FluxTableMetadata) String() string {
	var buffer strings.Builder
	for i, c := range f.columns {
		if i > 0 {
			buffer.WriteString(",")
		}
		buffer.WriteString("col")
		buffer.WriteString(c.String())
	}
	return buffer.String()
}

// NewFluxColumn creates FluxColumn for id
func NewFluxColumn() *FluxColumn {
	return &FluxColumn{}
}

// NewFluxColumnFull creates FluxColumn
func NewFluxColumnFull(dataType ColType, defaultValue string, name string, group bool) *FluxColumn {
	return &FluxColumn{name: name, dataType: dataType, group: group, defaultValue: defaultValue}
}

// SetDefaultValue sets default value for the column
func (f *FluxColumn) SetDefaultValue(defaultValue string) {
	f.defaultValue = defaultValue
}

// SetGroup set group flag for the column
func (f *FluxColumn) SetGroup(group bool) {
	f.group = group
}

// SetDataType sets data type for the column
func (f *FluxColumn) SetDataType(dataType ColType) {
	f.dataType = dataType
}

// SetName sets name of the column
func (f *FluxColumn) SetName(name string) {
	f.name = name
}

// DefaultValue returns default value of the column
func (f *FluxColumn) DefaultValue() string {
	return f.defaultValue
}

// IsGroup return true if the column is grouping column
func (f *FluxColumn) IsGroup() bool {
	return f.group
}

// DataType returns data type of the column
func (f *FluxColumn) DataType() ColType {
	return f.dataType
}

// Name returns name of the column
func (f *FluxColumn) Name() string {
	return f.name
}

// String returns FluxColumn string dump
func (f *FluxColumn) String() string {
	return fmt.Sprintf("{name: %s, datatype: %v, defaultValue: %s, group: %v}", f.name, f.dataType, f.defaultValue, f.group)
}

// NewFluxRecord returns new record for the table with values
func NewFluxRecord(metadata *FluxTableMetadata, values map[string]interface{}) (*FluxRecord, error) {
	res := stringValue(values, ResultCol)
	if res == "" && metadata.ResultColumn() != nil {
		res = metadata.ResultColumn().DefaultValue()
	}
	delete(values, ResultCol)

	var tid int64
	if v, ok := values[TableIdCol]; ok {
		if tid, ok = v.(int64); !ok {
			return nil, fmt.Errorf("invalid value for table ID: %s", v)
		}
	} else if metadata.TableIdColumn() != nil {
		if did := metadata.TableIdColumn().DefaultValue(); did != "" {
			if parsedId, err := strconv.Atoi(did); err != nil {
				return nil, fmt.Errorf("invalid default value for table ID: %s", did)
			} else {
				tid = int64(parsedId)
			}
		}
	}

	return &FluxRecord{metadata: metadata, result: res, tableId: tid, values: values}, nil
}

// Result returns the name of the result containing this record as specified by the query.
func (r *FluxRecord) Result() string {
	return r.result
}

// TableId returns index of the table record belongs to within its result.
func (r *FluxRecord) TableId() int64 {
	return r.tableId
}

// Start returns the inclusive lower time bound of all records in the current table.
// Returns empty time.Time if there is no column "_start".
func (r *FluxRecord) Start() time.Time {
	return timeValue(r.values, "_start")
}

// Stop returns the exclusive upper time bound of all records in the current table.
// Returns empty time.Time if there is no column "_stop".
func (r *FluxRecord) Stop() time.Time {
	return timeValue(r.values, "_stop")
}

// Time returns the time of the record.
// Returns empty time.Time if there is no column "_time".
func (r *FluxRecord) Time() time.Time {
	return timeValue(r.values, "_time")
}

// Value returns the default _value column value or nil if not present
func (r *FluxRecord) Value() interface{} {
	return r.ValueByKey("_value")
}

// Field returns the field name.
// Returns empty string if there is no column "_field".
func (r *FluxRecord) Field() string {
	return stringValue(r.values, "_field")
}

// Measurement returns the measurement name of the record
// Returns empty string if there is no column "_measurement".
func (r *FluxRecord) Measurement() string {
	return stringValue(r.values, "_measurement")
}

// Values returns map of the values where key is the column name
func (r *FluxRecord) Values() map[string]interface{} {
	return r.values
}

// ValueByKey returns value for given column key for the record or nil of result has no value the column key
func (r *FluxRecord) ValueByKey(key string) interface{} {
	return r.values[key]
}

// String returns FluxRecord string dump
func (r *FluxRecord) String() string {
	var buffer strings.Builder
	i := 0
	for k, v := range r.values {
		if i > 0 {
			buffer.WriteString(",")
		}
		buffer.WriteString(fmt.Sprintf("%s:%v", k, v))
		i++
	}
	return buffer.String()
}

// timeValue returns time.Time value from values map according to the key
// Empty time.Time value is returned if key is not found
func timeValue(values map[string]interface{}, key string) time.Time {
	if val, ok := values[key]; ok {
		if t, ok := val.(time.Time); ok {
			return t
		}
	}
	return time.Time{}
}

// timeValue returns string value from values map according to the key
// Empty string is returned if key is not found
func stringValue(values map[string]interface{}, key string) string {
	if val, ok := values[key]; ok {
		if s, ok := val.(string); ok {
			return s
		}
	}
	return ""
}
