package write_test

import (
	"bufio"
	"bytes"
	"compress/gzip"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
	"testing"

	"github.com/influxdata/influx-cli/v2/clients/write"
	"github.com/stretchr/testify/require"
)

func readLines(reader io.Reader) []string {
	scanner := bufio.NewScanner(reader)
	retVal := make([]string, 0, 3)
	for scanner.Scan() {
		retVal = append(retVal, scanner.Text())
	}
	return retVal
}

func createTempFile(t *testing.T, suffix string, contents []byte, compress bool) string {
	t.Helper()

	file, err := ioutil.TempFile("", "influx_writeTest*."+suffix)
	require.NoError(t, err)
	defer file.Close()

	var writer io.Writer = file
	if compress {
		gzipWriter := gzip.NewWriter(writer)
		defer gzipWriter.Close()
		writer = gzipWriter
	}

	_, err = writer.Write(contents)
	require.NoError(t, err)

	return file.Name()
}

type mockClient struct {
	t    *testing.T
	fail bool
}

func (c *mockClient) Do(req *http.Request) (*http.Response, error) {
	if c.fail {
		resp := http.Response{StatusCode: 500}
		return &resp, nil
	}
	query := req.URL.Query()
	resp := http.Response{Header: map[string][]string{}}
	if contentType := query.Get("Content-Type"); contentType != "" {
		resp.Header.Set("Content-Type", contentType)
	}
	if encoding := query.Get("encoding"); encoding != "" {
		resp.Header.Set("Content-Encoding", encoding)
	}
	compress := query.Get("compress") != ""
	resp.StatusCode = http.StatusOK
	if data := query.Get("data"); data != "" {
		body := bytes.Buffer{}
		var writer io.Writer = &body
		if compress {
			gzw := gzip.NewWriter(writer)
			defer gzw.Close()
			writer = gzw
		}
		_, err := writer.Write([]byte(data))
		require.NoError(c.t, err)
		resp.Body = ioutil.NopCloser(&body)
	}

	return &resp, nil
}

func TestLineReader(t *testing.T) {
	gzipStdin := func(uncompressed string) io.Reader {
		contents := &bytes.Buffer{}
		writer := gzip.NewWriter(contents)
		_, err := writer.Write([]byte(uncompressed))
		require.NoError(t, err)
		require.NoError(t, writer.Close())
		return contents
	}

	lpContents := "f1 b=f2,c=f3,d=f4"
	lpFile := createTempFile(t, "txt", []byte(lpContents), false)
	gzipLpFile := createTempFile(t, "txt.gz", []byte(lpContents), true)
	gzipLpFileNoExt := createTempFile(t, "lp", []byte(lpContents), true)
	stdInLpContents := "stdin3 i=stdin1,j=stdin2,k=stdin4"

	csvContents := "_measurement,b,c,d\nf1,f2,f3,f4"
	csvFile := createTempFile(t, "csv", []byte(csvContents), false)
	gzipCsvFile := createTempFile(t, "csv.gz", []byte(csvContents), true)
	gzipCsvFileNoExt := createTempFile(t, "csv", []byte(csvContents), true)
	stdInCsvContents := "i,j,_measurement,k\nstdin1,stdin2,stdin3,stdin4"

	defer func() {
		for _, f := range []string{lpFile, gzipLpFile, gzipLpFileNoExt, csvFile, gzipCsvFile, gzipCsvFileNoExt} {
			_ = os.Remove(f)
		}
	}()

	var tests = []struct {
		name string
		// input
		args                       []string
		files                      []string
		urls                       []string
		format                     write.InputFormat
		compression                write.InputCompression
		encoding                   string
		headers                    []string
		skipHeader                 int
		ignoreDataTypeInColumnName bool
		stdIn                      io.Reader
		// output
		firstLineCorrection int // 0 unless shifted by prepended headers or skipped rows
		lines               []string
	}{
		{
			name:                "read data from LP file",
			files:               []string{lpFile},
			firstLineCorrection: 0,
			lines: []string{
				lpContents,
			},
		},
		{
			name:                "read data from LP file using non-UTF encoding",
			files:               []string{lpFile},
			encoding:            "ISO_8859-1",
			firstLineCorrection: 0,
			lines: []string{
				lpContents,
			},
		},
		{
			name:                "read compressed LP data from file",
			files:               []string{gzipLpFileNoExt},
			compression:         write.InputCompressionGZIP,
			firstLineCorrection: 0,
			lines: []string{
				lpContents,
			},
		},
		{
			name:                "read compressed data from LP file using non-UTF encoding",
			files:               []string{gzipLpFileNoExt},
			compression:         write.InputCompressionGZIP,
			encoding:            "ISO_8859-1",
			firstLineCorrection: 0,
			lines: []string{
				lpContents,
			},
		},
		{
			name:                "read compressed LP data from file ending in .gz",
			files:               []string{gzipLpFile},
			firstLineCorrection: 0,
			lines: []string{
				lpContents,
			},
		},
		{
			name:                "read compressed and uncompressed LP data from file in the same call",
			files:               []string{gzipLpFile, lpFile},
			firstLineCorrection: 0,
			lines: []string{
				lpContents,
				lpContents,
			},
		},
		{
			name:  "read LP data from stdin",
			stdIn: strings.NewReader(stdInLpContents),
			lines: []string{
				stdInLpContents,
			},
		},
		{
			name:        "read compressed LP data from stdin",
			compression: write.InputCompressionGZIP,
			stdIn:       gzipStdin(stdInLpContents),
			lines: []string{
				stdInLpContents,
			},
		},
		{
			name:  "read LP data from stdin using '-' argument",
			args:  []string{"-"},
			stdIn: strings.NewReader(stdInLpContents),
			lines: []string{
				stdInLpContents,
			},
		},
		{
			name:        "read compressed LP data from stdin using '-' argument",
			compression: write.InputCompressionGZIP,
			args:        []string{"-"},
			stdIn:       gzipStdin(stdInLpContents),
			lines: []string{
				stdInLpContents,
			},
		},
		{
			name: "read LP data from 1st argument",
			args: []string{stdInLpContents},
			lines: []string{
				stdInLpContents,
			},
		},
		{
			name: "read LP data from URL",
			urls: []string{fmt.Sprintf("/a?data=%s", url.QueryEscape(lpContents))},
			lines: []string{
				lpContents,
			},
		},
		{
			name:        "read compressed LP data from URL",
			urls:        []string{fmt.Sprintf("/a?data=%s&compress=true", url.QueryEscape(lpContents))},
			compression: write.InputCompressionGZIP,
			lines: []string{
				lpContents,
			},
		},
		{
			name: "read compressed LP data from URL ending in .gz",
			urls: []string{fmt.Sprintf("/a.gz?data=%s&compress=true", url.QueryEscape(lpContents))},
			lines: []string{
				lpContents,
			},
		},
		{
			name: "read compressed LP data from URL with gzip encoding",
			urls: []string{fmt.Sprintf("/a?data=%s&compress=true&encoding=gzip", url.QueryEscape(lpContents))},
			lines: []string{
				lpContents,
			},
		},
		{
			name:                "read data from CSV file + transform to line protocol",
			files:               []string{csvFile},
			firstLineCorrection: 0, // no changes
			lines: []string{
				lpContents,
			},
		},
		{
			name:                "read compressed CSV data from file + transform to line protocol",
			files:               []string{gzipCsvFileNoExt},
			compression:         write.InputCompressionGZIP,
			firstLineCorrection: 0,
			lines: []string{
				lpContents,
			},
		},
		{
			name:                "read compressed CSV data from file ending in .csv.gz + transform to line protocol",
			files:               []string{gzipCsvFile},
			firstLineCorrection: 0,
			lines: []string{
				lpContents,
			},
		},
		{
			name:                "read CSV data from --header and --file + transform to line protocol",
			headers:             []string{"x,_measurement,y,z"},
			files:               []string{csvFile},
			firstLineCorrection: -1, // shifted back by header line
			lines: []string{
				"b x=_measurement,y=c,z=d",
				"f2 x=f1,y=f3,z=f4",
			},
		},
		{
			name:                "read CSV data from --header and @file argument with 1st row in file skipped + transform to line protocol",
			headers:             []string{"x,_measurement,y,z"},
			skipHeader:          1,
			args:                []string{"@" + csvFile},
			firstLineCorrection: 0, // shifted (-1) back by header line, forward (+1) by skipHeader
			lines: []string{
				"f2 x=f1,y=f3,z=f4",
			},
		},
		{
			name:   "read CSV data from stdin + transform to line protocol",
			format: write.InputFormatCSV,
			stdIn:  strings.NewReader(stdInCsvContents),
			lines: []string{
				stdInLpContents,
			},
		},
		{
			name:        "read compressed CSV data from stdin + transform to line protocol",
			format:      write.InputFormatCSV,
			compression: write.InputCompressionGZIP,
			stdIn:       gzipStdin(stdInCsvContents),
			lines: []string{
				stdInLpContents,
			},
		},
		{
			name:   "read CSV data from stdin using '-' argument + transform to line protocol",
			format: write.InputFormatCSV,
			args:   []string{"-"},
			stdIn:  strings.NewReader(stdInCsvContents),
			lines: []string{
				stdInLpContents,
			},
		},
		{
			name:        "read compressed CSV data from stdin using '-' argument + transform to line protocol",
			format:      write.InputFormatCSV,
			compression: write.InputCompressionGZIP,
			args:        []string{"-"},
			stdIn:       gzipStdin(stdInCsvContents),
			lines: []string{
				stdInLpContents,
			},
		},
		{
			name:   "read CSV data from 1st argument + transform to line protocol",
			format: write.InputFormatCSV,
			args:   []string{stdInCsvContents},
			lines: []string{
				stdInLpContents,
			},
		},
		{
			name: "read data from .csv URL + transform to line protocol",
			urls: []string{fmt.Sprintf("/a.csv?data=%s", url.QueryEscape(csvContents))},
			lines: []string{
				lpContents,
			},
		},
		{
			name:        "read compressed CSV data from URL + transform to line protocol",
			urls:        []string{fmt.Sprintf("/a.csv?data=%s&compress=true", url.QueryEscape(csvContents))},
			compression: write.InputCompressionGZIP,
			lines: []string{
				lpContents,
			},
		},
		{
			name: "read compressed CSV data from URL ending in .csv.gz + transform to line protocol",
			urls: []string{fmt.Sprintf("/a.csv.gz?data=%s&compress=true", url.QueryEscape(csvContents))},
			lines: []string{
				lpContents,
			},
		},
		{
			name: "read compressed CSV data from URL with gzip encoding + transform to line protocol",
			urls: []string{fmt.Sprintf("/a.csv?data=%s&compress=true&encoding=gzip", url.QueryEscape(csvContents))},
			lines: []string{
				lpContents,
			},
		},
		{
			name:       "read data from .csv URL + change header line + transform to line protocol",
			urls:       []string{fmt.Sprintf("/a.csv?data=%s", url.QueryEscape(csvContents))},
			headers:    []string{"k,j,_measurement,i"},
			skipHeader: 1,
			lines: []string{
				"f3 k=f1,j=f2,i=f4",
			},
		},
		{
			name: "read data from URL with text/csv Content-Type + transform to line protocol",
			urls: []string{fmt.Sprintf("/a?Content-Type=text/csv&data=%s", url.QueryEscape(csvContents))},
			lines: []string{
				lpContents,
			},
		},
		{
			name: "read compressed data from URL with text/csv Content-Type and gzip Content-Encoding + transform to line protocol",
			urls: []string{fmt.Sprintf("/a?Content-Type=text/csv&data=%s&compress=true&encoding=gzip", url.QueryEscape(csvContents))},
			lines: []string{
				lpContents,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			r := &write.MultiInputLineReader{
				StdIn:                      test.stdIn,
				HttpClient:                 &mockClient{t: t},
				Args:                       test.args,
				Files:                      test.files,
				URLs:                       test.urls,
				Format:                     test.format,
				Compression:                test.compression,
				Headers:                    test.headers,
				SkipHeader:                 test.skipHeader,
				IgnoreDataTypeInColumnName: test.ignoreDataTypeInColumnName,
			}
			reader, closer, err := r.Open(context.Background())
			require.NotNil(t, closer)
			defer closer.Close()
			require.NoError(t, err)
			require.NotNil(t, reader)
			lines := readLines(reader)
			require.Equal(t, test.lines, lines)
		})
	}
}

func TestLineReaderErrors(t *testing.T) {
	csvFile1 := createTempFile(t, "csv", []byte("_measurement,b,c,d\nf1,f2,f3,f4"), false)
	defer os.Remove(csvFile1)

	var tests = []struct {
		name     string
		encoding string
		files    []string
		urls     []string
		message  string
	}{
		{
			name:     "unsupported encoding",
			encoding: "green",
			message:  "https://www.iana.org/assignments/character-sets/character-sets.xhtml", // hint to available values
		},
		{
			name:    "file not found",
			files:   []string{csvFile1 + "x"},
			message: csvFile1,
		},
		{
			name:    "unsupported URL",
			urls:    []string{"wit://whatever"},
			message: "wit://whatever",
		},
		{
			name:    "invalid URL",
			urls:    []string{"http://test%zy"}, // 2 hex digits after % expected
			message: "http://test%zy",
		},
		{
			name:    "URL with 500 status code",
			urls:    []string{"/test"},
			message: "/test",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			r := write.MultiInputLineReader{
				HttpClient: &mockClient{t: t, fail: true},
				Files:      test.files,
				URLs:       test.urls,
				Encoding:   test.encoding,
			}
			_, closer, err := r.Open(context.Background())
			require.NotNil(t, closer)
			defer closer.Close()
			require.Error(t, err)
			require.Contains(t, err.Error(), test.message)
		})
	}
}

func TestLineReaderErrorOut(t *testing.T) {
	stdInContents := "_measurement,a|long:strict\nm,1\nm,1.1"
	errorOut := bytes.Buffer{}

	r := write.MultiInputLineReader{
		StdIn:    strings.NewReader(stdInContents),
		ErrorOut: &errorOut,
		Format:   write.InputFormatCSV,
	}
	reader, closer, err := r.Open(context.Background())
	require.NoError(t, err)
	defer closer.Close()

	out := bytes.Buffer{}
	_, err = io.Copy(&out, reader)
	require.NoError(t, err)

	require.Equal(t, "m a=1i", strings.Trim(out.String(), "\n"))
	errorLines := errorOut.String()
	require.Equal(t, "# error : line 3: column 'a': '1.1' cannot fit into long data type\nm,1.1", strings.Trim(errorLines, "\n"))
}
