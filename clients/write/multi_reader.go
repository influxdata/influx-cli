package write

import (
	"compress/gzip"
	"context"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/influxdata/influx-cli/v2/pkg/csv2lp"
)

var _ HttpClient = (*http.Client)(nil)

type HttpClient interface {
	Do(*http.Request) (*http.Response, error)
}

type InputFormat int

const (
	InputFormatDerived InputFormat = iota
	InputFormatCSV
	InputFormatLP
)

func (i *InputFormat) Set(v string) error {
	switch v {
	case "":
		*i = InputFormatDerived
	case "lp":
		*i = InputFormatLP
	case "csv":
		*i = InputFormatCSV
	default:
		return fmt.Errorf("unsupported format: %q", v)
	}
	return nil
}

func (i InputFormat) String() string {
	switch i {
	case InputFormatLP:
		return "lp"
	case InputFormatCSV:
		return "csv"
	case InputFormatDerived:
		fallthrough
	default:
		return ""
	}
}

type InputCompression int

const (
	InputCompressionDerived InputCompression = iota
	InputCompressionGZIP
	InputCompressionNone
)

func (i *InputCompression) Set(v string) error {
	switch v {
	case "":
		*i = InputCompressionDerived
	case "none":
		*i = InputCompressionNone
	case "gzip":
		*i = InputCompressionGZIP
	default:
		return fmt.Errorf("unsupported compression: %q", v)
	}
	return nil
}

func (i InputCompression) String() string {
	switch i {
	case InputCompressionNone:
		return "none"
	case InputCompressionGZIP:
		return "gzip"
	case InputCompressionDerived:
		fallthrough
	default:
		return ""
	}
}

type MultiInputLineReader struct {
	StdIn      io.Reader
	HttpClient HttpClient
	ErrorOut   io.Writer

	Args        []string
	Files       []string
	URLs        []string
	Format      InputFormat
	Compression InputCompression
	Encoding    string

	// CSV-specific options.
	Headers                    []string
	SkipRowOnError             bool
	SkipHeader                 int
	IgnoreDataTypeInColumnName bool
	Debug                      bool
}

func (r *MultiInputLineReader) Open(ctx context.Context) (io.Reader, io.Closer, error) {
	if r.Debug {
		log.Printf("WriteFlags%+v", r)
	}

	args := r.Args
	files := r.Files
	if len(args) > 0 && len(args[0]) > 1 && args[0][0] == '@' {
		// backward compatibility: @ in arg denotes a file
		files = append(files, args[0][1:])
		args = args[:0]
	}

	readers := make([]io.Reader, 0, 2*len(r.Headers)+2*len(files)+2*len(r.URLs)+1)
	closers := make([]io.Closer, 0, len(files)+len(r.URLs))

	// validate and setup decoding of files/stdin if encoding is supplied
	decode, err := csv2lp.CreateDecoder(r.Encoding)
	if err != nil {
		return nil, csv2lp.MultiCloser(closers...), err
	}

	// utility to manage common steps used to decode / decompress input sources,
	// while tracking resources that must be cleaned-up after reading.
	addReader := func(r io.Reader, name string, compressed bool) error {
		if compressed {
			rcz, err := gzip.NewReader(r)
			if err != nil {
				return fmt.Errorf("failed to decompress %s: %w", name, err)
			}
			closers = append(closers, rcz)
			r = rcz
		}
		readers = append(readers, decode(r), strings.NewReader("\n"))
		return nil
	}

	// prepend header lines
	if len(r.Headers) > 0 {
		for _, header := range r.Headers {
			readers = append(readers, strings.NewReader(header), strings.NewReader("\n"))

		}
		if r.Format == InputFormatDerived {
			r.Format = InputFormatCSV
		}
	}

	// add files
	for _, file := range files {
		f, err := os.Open(file)
		if err != nil {
			return nil, csv2lp.MultiCloser(closers...), fmt.Errorf("failed to open %q: %v", file, err)
		}
		closers = append(closers, f)

		fname := file
		compressed := r.Compression == InputCompressionGZIP || (r.Compression == InputCompressionDerived && strings.HasSuffix(fname, ".gz"))
		if compressed {
			fname = strings.TrimSuffix(fname, ".gz")
		}
		if r.Format == InputFormatDerived && strings.HasSuffix(fname, ".csv") {
			r.Format = InputFormatCSV
		}

		if err = addReader(f, file, compressed); err != nil {
			return nil, csv2lp.MultiCloser(closers...), err
		}
	}

	// allow URL data sources, a simple alternative to `curl -f -s http://... | influx batcher ...`
	for _, addr := range r.URLs {
		u, err := url.Parse(addr)
		if err != nil {
			return nil, csv2lp.MultiCloser(closers...), fmt.Errorf("failed to open %q: %v", addr, err)
		}
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, addr, nil)
		if err != nil {
			return nil, csv2lp.MultiCloser(closers...), fmt.Errorf("failed to open %q: %v", addr, err)
		}
		req.Header.Set("Accept-Encoding", "gzip")
		resp, err := r.HttpClient.Do(req)
		if err != nil {
			return nil, csv2lp.MultiCloser(closers...), fmt.Errorf("failed to open %q: %v", addr, err)
		}
		if resp.Body != nil {
			closers = append(closers, resp.Body)
		}
		if resp.StatusCode/100 != 2 {
			return nil, csv2lp.MultiCloser(closers...), fmt.Errorf("failed to open %q: response status_code=%d", addr, resp.StatusCode)
		}

		compressed := r.Compression == InputCompressionGZIP ||
			resp.Header.Get("Content-Encoding") == "gzip" ||
			(r.Compression == InputCompressionDerived && strings.HasSuffix(u.Path, ".gz"))
		if compressed {
			u.Path = strings.TrimSuffix(u.Path, ".gz")
		}
		if r.Format == InputFormatDerived &&
			(strings.HasSuffix(u.Path, ".csv") || strings.HasPrefix(resp.Header.Get("Content-Type"), "text/csv")) {
			r.Format = InputFormatCSV
		}

		if err = addReader(resp.Body, addr, compressed); err != nil {
			return nil, csv2lp.MultiCloser(closers...), err
		}
	}

	// add stdin or a single argument
	switch {
	case len(args) == 0:
		// use also stdIn if it is a terminal
		if r.StdIn != nil && !isCharacterDevice(r.StdIn) {
			if err = addReader(r.StdIn, "stdin", r.Compression == InputCompressionGZIP); err != nil {
				return nil, csv2lp.MultiCloser(closers...), err
			}
		}
	case args[0] == "-":
		// "-" also means stdin
		if err = addReader(r.StdIn, "stdin", r.Compression == InputCompressionGZIP); err != nil {
			return nil, csv2lp.MultiCloser(closers...), err
		}
	default:
		if err = addReader(strings.NewReader(args[0]), "arg 0", r.Compression == InputCompressionGZIP); err != nil {
			return nil, csv2lp.MultiCloser(closers...), err
		}
	}

	// skipHeader lines when set
	if r.SkipHeader > 0 {
		// find the last non-string reader (stdin or file)
		for i := len(readers) - 1; i >= 0; i-- {
			_, stringReader := readers[i].(*strings.Reader)
			if !stringReader { // ignore headers and new lines
				readers[i] = csv2lp.SkipHeaderLinesReader(r.SkipHeader, readers[i])
				break
			}
		}
	}

	// create writer for errors-file, if supplied
	var errorsFile *csv.Writer
	var rowSkippedListener func(*csv2lp.CsvToLineReader, error, []string)
	if r.ErrorOut != nil {
		errorsFile = csv.NewWriter(r.ErrorOut)
		rowSkippedListener = func(source *csv2lp.CsvToLineReader, lineError error, row []string) {
			log.Println(lineError)
			errorsFile.Comma = source.Comma()
			errorsFile.Write([]string{fmt.Sprintf("# error : %v", lineError)})
			if err := errorsFile.Write(row); err != nil {
				log.Printf("Unable to batcher to error-file: %v\n", err)
			}
			errorsFile.Flush() // flush is required
		}
	}

	// concatenate readers
	reader := io.MultiReader(readers...)
	if r.Format == InputFormatCSV {
		csvReader := csv2lp.CsvToLineProtocol(reader)
		csvReader.LogTableColumns(r.Debug)
		csvReader.SkipRowOnError(r.SkipRowOnError)
		csvReader.Table.IgnoreDataTypeInColumnName(r.IgnoreDataTypeInColumnName)
		// change LineNumber to report file/stdin line numbers properly
		csvReader.LineNumber = r.SkipHeader - len(r.Headers)
		csvReader.RowSkipped = rowSkippedListener
		reader = csvReader
	} else if r.SkipRowOnError {
		reader = csv2lp.LineProtocolFilter(reader)
	}

	return reader, csv2lp.MultiCloser(closers...), nil
}

// isCharacterDevice returns true if the supplied reader is a character device (a terminal)
func isCharacterDevice(reader io.Reader) bool {
	file, isFile := reader.(*os.File)
	if !isFile {
		return false
	}
	info, err := file.Stat()
	if err != nil {
		return false
	}
	return (info.Mode() & os.ModeCharDevice) == os.ModeCharDevice
}
