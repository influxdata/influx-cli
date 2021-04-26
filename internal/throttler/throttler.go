package throttler

import (
	"context"
	"fmt"
	"io"
	"regexp"
	"strconv"
	"strings"

	"github.com/fujiwara/shapeio"
	"github.com/influxdata/influx-cli/v2/pkg/csv2lp"
)

type Throttler struct {
	bytesPerSecond float64
}

func NewThrottler(rate string) (*Throttler, error) {
	bytesPerSec, err := ToBytesPerSecond(rate)
	if err != nil {
		return nil, err
	}
	return &Throttler{bytesPerSecond: bytesPerSec}, nil
}

func (t *Throttler) Throttle(ctx context.Context, in io.Reader) io.Reader {
	if t.bytesPerSecond == 0.0 {
		return in
	}

	// LineReader ensures that original reader is consumed in the smallest possible
	// units (at most one protocol line) to avoid bigger pauses in throttling
	throttledReader := shapeio.NewReaderWithContext(csv2lp.NewLineReader(in), ctx)
	throttledReader.SetRateLimit(t.bytesPerSecond)

	return throttledReader
}

var rateLimitRegexp = regexp.MustCompile(`^(\d*\.?\d*)(B|kB|MB)/?(\d*)?(s|sec|m|min)$`)
var bytesUnitMultiplier = map[string]float64{"B": 1, "kB": 1024, "MB": 1_048_576}
var timeUnitMultiplier = map[string]float64{"s": 1, "sec": 1, "m": 60, "min": 60}

// ToBytesPerSecond converts rate from string to number. The supplied string
// value format must be COUNT(B|kB|MB)/TIME(s|sec|m|min) with / and TIME being optional.
// All spaces are ignored, they can help with formatting. Examples: "5 MB / 5 min", 17kbs. 5.1MB5m.
func ToBytesPerSecond(rateLimit string) (float64, error) {
	// ignore all spaces
	strVal := strings.ReplaceAll(rateLimit, " ", "")
	if len(strVal) == 0 {
		return 0, nil
	}

	matches := rateLimitRegexp.FindStringSubmatch(strVal)
	if matches == nil {
		return 0, fmt.Errorf("invalid rate limit %q: it does not match format COUNT(B|kB|MB)/TIME(s|sec|m|min) with / and TIME being optional, rexpexp: %v", strVal, rateLimitRegexp)
	}
	bytes, err := strconv.ParseFloat(matches[1], 64)
	if err != nil {
		return 0, fmt.Errorf("invalid rate limit %q: '%v' is not count of bytes: %v", strVal, matches[1], err)
	}
	bytes = bytes * bytesUnitMultiplier[matches[2]]
	var time float64
	if len(matches[3]) == 0 {
		time = 1 // number is not specified, for example 5kbs or 1Mb/s
	} else {
		int64Val, err := strconv.ParseUint(matches[3], 10, 32)
		if err != nil {
			return 0, fmt.Errorf("invalid rate limit %q: time is out of range: %v", strVal, err)
		}
		if int64Val <= 0 {
			return 0, fmt.Errorf("invalid rate limit %q: positive time expected but %v supplied", strVal, matches[3])
		}
		time = float64(int64Val)
	}
	time = time * timeUnitMultiplier[matches[4]]
	return bytes / time, nil
}
