package csv2lp

import (
	"bytes"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

// Test_CsvToLineProtocol_variousBufferSize tests conversion of annotated CSV data to line protocol data on various buffer sizes
func Test_CsvToLineProtocol_variousBufferSize(t *testing.T) {
	var tests = []struct {
		name  string
		csv   string
		lines string
		err   string
	}{
		{
			"simple1",
			"_measurement,a,b\ncpu,1,1\ncpu,b2\n",
			"cpu a=1,b=1\ncpu a=b2\n",
			"",
		},
		{
			"simple1_withSep",
			"sep=;\n_measurement;a;b\ncpu;1;1\ncpu;b2\n",
			"cpu a=1,b=1\ncpu a=b2\n",
			"",
		},
		{
			"simple2",
			"_measurement,a,b\ncpu,1,1\ncpu,\n",
			"",
			"no field data",
		},
		{
			"simple3",
			"_measurement,a,_time\ncpu,1,1\ncpu,2,invalidTime\n",
			"",
			"_time", // error in _time column
		},
		{
			"constant_annotations",
			"#constant,measurement,,cpu\n" +
				"#constant,tag,xpu,xpu1\n" +
				"#constant,tag,cpu,cpu1\n" +
				"#constant,long,of,100\n" +
				"#constant,dateTime,,2\n" +
				"x,y\n" +
				"1,2\n" +
				"3,4\n",
			"cpu,cpu=cpu1,xpu=xpu1 x=1,y=2,of=100i 2\n" +
				"cpu,cpu=cpu1,xpu=xpu1 x=3,y=4,of=100i 2\n",
			"", // no error
		},
		{
			"timezone_annotation-0100",
			"#timezone,-0100\n" +
				"#constant,measurement,cpu\n" +
				"#constant,dateTime:2006-01-02,1970-01-01\n" +
				"x,y\n" +
				"1,2\n",
			"cpu x=1,y=2 3600000000000\n",
			"", // no error
		},
		{
			"timezone_annotation_EST",
			"#timezone,EST\n" +
				"#constant,measurement,cpu\n" +
				"#constant,dateTime:2006-01-02,1970-01-01\n" +
				"x,y\n" +
				"1,2\n",
			"cpu x=1,y=2 18000000000000\n",
			"", // no error
		},
	}
	bufferSizes := []int{40, 7, 3, 1}

	for _, test := range tests {
		for _, bufferSize := range bufferSizes {
			t.Run(test.name+"_"+strconv.Itoa(bufferSize), func(t *testing.T) {
				reader := CsvToLineProtocol(strings.NewReader(test.csv))
				buffer := make([]byte, bufferSize)
				lines := make([]byte, 0, 100)
				for {
					n, err := reader.Read(buffer)
					if err != nil {
						if err == io.EOF {
							break
						}
						if test.err != "" {
							// fmt.Println(err)
							if err := err.Error(); !strings.Contains(err, test.err) {
								require.Equal(t, err, test.err)
							}
							return
						}
						require.Nil(t, err.Error())
						break
					}
					lines = append(lines, buffer[:n]...)
				}
				if test.err == "" {
					require.Equal(t, test.lines, string(lines))
				} else {
					require.Fail(t, "error message with '"+test.err+"' expected")
				}
			})
		}
	}
}

// Test_CsvToLineProtocol_samples tests conversion of annotated CSV data to line protocol data
func Test_CsvToLineProtocol_samples(t *testing.T) {
	var tests = []struct {
		name  string
		csv   string
		lines string
		err   string
	}{
		{
			"queryResult_19452", // https://github.com/influxdata/influxdb/issues/19452
			"#datatype,string,long,dateTime:RFC3339,dateTime:RFC3339,dateTime:RFC3339,long,string,string,string\n" +
				"#group,false,false,true,true,false,false,true,true,true\n" +
				"#default,_result,,,,,,,,\n" +
				",result,table,_start,_stop,_time,_value,_field,_measurement,host\n" +
				",,0,2020-08-26T22:59:23.598653Z,2020-08-26T23:00:23.598653Z,2020-08-26T22:59:30Z,15075651584,active,mem,ip-192-168-86-25.ec2.internal\n",
			"mem,host=ip-192-168-86-25.ec2.internal active=15075651584i 1598482770000000000\n",
			"", // no error
		},
		{
			"queryResult_19452_group_first", // issue 19452, but with group annotation first
			"#group,false,false,true,true,false,false,true,true,true\n" +
				"#datatype,string,long,dateTime:RFC3339,dateTime:RFC3339,dateTime:RFC3339,long,string,string,string\n" +
				"#default,_result,,,,,,,,\n" +
				",result,table,_start,_stop,_time,_value,_field,_measurement,host\n" +
				",,0,2020-08-26T22:59:23.598653Z,2020-08-26T23:00:23.598653Z,2020-08-26T22:59:30Z,15075651584,active,mem,ip-192-168-86-25.ec2.internal\n",
			"mem,host=ip-192-168-86-25.ec2.internal active=15075651584i 1598482770000000000\n",
			"", // no error
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			reader := CsvToLineProtocol(strings.NewReader(test.csv))
			buffer := make([]byte, 100)
			lines := make([]byte, 0, 100)
			for {
				n, err := reader.Read(buffer)
				if err != nil {
					if err == io.EOF {
						break
					}
					if test.err != "" {
						// fmt.Println(err)
						if err := err.Error(); !strings.Contains(err, test.err) {
							require.Equal(t, err, test.err)
						}
						return
					}
					require.Nil(t, err.Error())
					break
				}
				lines = append(lines, buffer[:n]...)
			}
			if test.err == "" {
				require.Equal(t, test.lines, string(lines))
			} else {
				require.Fail(t, "error message with '"+test.err+"' expected")
			}
		})
	}
}

// Test_CsvToLineProtocol_LogTableColumns checks correct logging of table columns
func Test_CsvToLineProtocol_LogTableColumns(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	oldFlags := log.Flags()
	log.SetFlags(0)
	oldPrefix := log.Prefix()
	prefix := "::PREFIX::"
	log.SetPrefix(prefix)
	defer func() {
		log.SetOutput(os.Stderr)
		log.SetFlags(oldFlags)
		log.SetPrefix(oldPrefix)
	}()

	csv := "_measurement,a,b\ncpu,1,1\ncpu,b2\n"

	reader := CsvToLineProtocol(strings.NewReader(csv)).LogTableColumns(true)
	require.False(t, reader.skipRowOnError)
	require.True(t, reader.logTableDataColumns)
	// read all the data
	ioutil.ReadAll(reader)

	out := buf.String()
	// fmt.Println(out)
	messages := strings.Count(out, prefix)
	require.Equal(t, messages, 1)
}

// Test_CsvToLineProtocol_LogTimeZoneWarning checks correct logging of timezone warning
func Test_CsvToLineProtocol_LogTimeZoneWarning(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	oldFlags := log.Flags()
	log.SetFlags(0)
	oldPrefix := log.Prefix()
	prefix := "::PREFIX::"
	log.SetPrefix(prefix)
	defer func() {
		log.SetOutput(os.Stderr)
		log.SetFlags(oldFlags)
		log.SetPrefix(oldPrefix)
	}()

	csv := "#timezone 1\n" +
		"#constant,dateTime:2006-01-02,1970-01-01\n" +
		"_measurement,a,b\ncpu,1,1"

	reader := CsvToLineProtocol(strings.NewReader(csv))
	bytes, _ := ioutil.ReadAll(reader)

	out := buf.String()
	// fmt.Println(out) // "::PREFIX::WARNING:  #timezone annotation: unknown time zone 1
	messages := strings.Count(out, prefix)
	require.Equal(t, messages, 1)
	require.Equal(t, string(bytes), "cpu a=1,b=1 0\n")
}

// Test_CsvToLineProtocol_SkipRowOnError tests that error rows are skipped
func Test_CsvToLineProtocol_SkipRowOnError(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	oldFlags := log.Flags()
	log.SetFlags(0)
	oldPrefix := log.Prefix()
	prefix := "::PREFIX::"
	log.SetPrefix(prefix)
	defer func() {
		log.SetOutput(os.Stderr)
		log.SetFlags(oldFlags)
		log.SetPrefix(oldPrefix)
	}()

	csv := "_measurement,a,_time\n,1,1\ncpu,2,2\ncpu,3,3a\n"

	reader := CsvToLineProtocol(strings.NewReader(csv)).SkipRowOnError(true)
	require.Equal(t, reader.skipRowOnError, true)
	require.Equal(t, reader.logTableDataColumns, false)
	// read all the data
	ioutil.ReadAll(reader)

	out := buf.String()
	// fmt.Println(out)
	messages := strings.Count(out, prefix)
	require.Equal(t, messages, 2)
}

// Test_CsvToLineProtocol_RowSkipped tests that error rows are reported to configured RowSkipped listener
func Test_CsvToLineProtocol_RowSkipped(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	oldFlags := log.Flags()
	log.SetFlags(0)
	defer func() {
		log.SetOutput(os.Stderr)
		log.SetFlags(oldFlags)
	}()

	type ActualArguments = struct {
		src *CsvToLineReader
		err error
		row []string
	}
	type ExpectedArguments = struct {
		errorString string
		row         []string
	}

	csv := "sep=;\n_measurement;a|long:strict\n;1\ncpu;2.1\ncpu;3a\n"
	calledArgs := []ActualArguments{}
	expectedArgs := []ExpectedArguments{
		{
			"line 3: column '_measurement': no measurement supplied",
			[]string{"", "1"},
		},
		{
			"line 4: column 'a': '2.1' cannot fit into long data type",
			[]string{"cpu", "2.1"},
		},
		{
			"line 5: column 'a': strconv.ParseInt:",
			[]string{"cpu", "3a"},
		},
	}

	reader := CsvToLineProtocol(strings.NewReader(csv)).SkipRowOnError(true)
	reader.RowSkipped = func(src *CsvToLineReader, err error, _row []string) {
		// make a copy of _row
		row := make([]string, len(_row))
		copy(row, _row)
		// remember for comparison
		calledArgs = append(calledArgs, ActualArguments{
			src, err, row,
		})
	}
	// read all the data
	ioutil.ReadAll(reader)

	out := buf.String()
	require.Empty(t, out, "No log messages expected because RowSkipped handler is set")

	require.Len(t, calledArgs, 3)
	for i, expected := range expectedArgs {
		require.Equal(t, reader, calledArgs[i].src)
		require.Contains(t, calledArgs[i].err.Error(), expected.errorString)
		require.Equal(t, expected.row, calledArgs[i].row)
	}
}

// Test_CsvLineError tests CsvLineError error format
func Test_CsvLineError(t *testing.T) {
	var tests = []struct {
		err   CsvLineError
		value string
	}{
		{
			CsvLineError{Line: 1, Err: errors.New("cause")},
			"line 1: cause",
		},
		{
			CsvLineError{Line: 2, Err: CsvColumnError{"a", errors.New("cause")}},
			"line 2: column 'a': cause",
		},
		{
			CsvLineError{Line: -1, Err: CsvColumnError{"a", errors.New("cause")}},
			"column 'a': cause",
		},
	}
	for _, test := range tests {
		require.Equal(t, test.value, test.err.Error())
	}
}

// Test_CsvToLineProtocol_lineNumbers tests that correct line numbers are reported
func Test_CsvToLineProtocol_lineNumbers(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	oldFlags := log.Flags()
	log.SetFlags(0)
	defer func() {
		log.SetOutput(os.Stderr)
		log.SetFlags(oldFlags)
	}()

	type ActualArguments = struct {
		src *CsvToLineReader
		err error
		row []string
	}
	type ExpectedArguments = struct {
		errorString string
		row         []string
	}

	// note: csv contains a field with newline and an extra empty lines
	csv := `sep=;

_measurement;a|long:strict
;1
cpu;"2
.1"

cpu;3a
`
	calledArgs := []ActualArguments{}
	expectedArgs := []ExpectedArguments{
		{
			"line 4: column '_measurement': no measurement supplied",
			[]string{"", "1"},
		},
		{
			"line 6: column 'a': '2\n.1' cannot fit into long data type",
			[]string{"cpu", "2\n.1"},
		},
		{
			"line 8: column 'a': strconv.ParseInt:",
			[]string{"cpu", "3a"},
		},
	}

	reader := CsvToLineProtocol(strings.NewReader(csv)).SkipRowOnError(true)
	reader.RowSkipped = func(src *CsvToLineReader, err error, _row []string) {
		// make a copy of _row
		row := make([]string, len(_row))
		copy(row, _row)
		// remember for comparison
		calledArgs = append(calledArgs, ActualArguments{
			src, err, row,
		})
	}
	// read all the data
	ioutil.ReadAll(reader)

	out := buf.String()
	require.Empty(t, out, "No log messages expected because RowSkipped handler is set")

	require.Len(t, calledArgs, 3)
	for i, expected := range expectedArgs {
		require.Equal(t, reader, calledArgs[i].src)
		require.Contains(t, calledArgs[i].err.Error(), expected.errorString)
		require.Equal(t, expected.row, calledArgs[i].row)
	}
	// 8 lines were read
	require.Equal(t, 8, reader.LineNumber)
}
