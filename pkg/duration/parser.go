package duration

import (
	"errors"
	"fmt"
	"strconv"
	"time"
	"unicode"
	"unicode/utf8"
)

const (
	Day  = 24 * time.Hour
	Week = 7 * Day
)

type durations []duration

type duration struct {
	magnitude int64
	unit      string
}

var (
	ErrInvalidUnit = errors.New("duration must be week(w), day(d), hour(h), min(m), sec(s), millisec(ms), microsec(us), or nanosec(ns)")
	ErrInvalidRune = errors.New("invalid rune in declaration")
)

// RawDurationToTimeDuration extends the builtin duration-parser to support days(d) and weeks(w) as parseable units.
// The core implementation is copied from InfluxDB OSS's task engine, which itself copied the logic from an
// internal module in Flux.
func RawDurationToTimeDuration(raw string) (time.Duration, error) {
	if raw == "" {
		return 0, nil
	}

	if dur, err := time.ParseDuration(raw); err == nil {
		return dur, nil
	}

	parsed, err := parseSignedDuration(raw)
	if err != nil {
		return 0, err
	}

	var dur time.Duration
	for _, d := range parsed {
		if d.magnitude < 0 {
			return 0, errors.New("must be greater than 0")
		}
		mag := time.Duration(d.magnitude)
		switch d.unit {
		case "w":
			dur += mag * Week
		case "d":
			dur += mag * Day
		case "h":
			dur += mag * time.Hour
		case "m":
			dur += mag * time.Minute
		case "s":
			dur += mag * time.Second
		case "ms":
			dur += mag * time.Millisecond
		case "us":
			dur += mag * time.Microsecond
		case "ns":
			dur += mag * time.Nanosecond
		default:
			return 0, ErrInvalidUnit
		}
	}
	return dur, nil
}

func parseSignedDuration(text string) (durations, error) {
	if r, s := utf8.DecodeRuneInString(text); r == '-' {
		d, err := parseDuration(text[s:])
		if err != nil {
			return nil, err
		}
		for i := range d {
			d[i].magnitude = -d[i].magnitude
		}
		return d, nil
	}

	d, err := parseDuration(text)
	if err != nil {
		return nil, err
	}
	return d, nil
}

// parseDuration will convert a string into components of the duration.
func parseDuration(input string) (durations, error) {
	lit := input
	var values durations
	for len(lit) > 0 {
		n := 0
		for n < len(lit) {
			ch, size := utf8.DecodeRuneInString(lit[n:])
			if size == 0 {
				return nil, ErrInvalidRune
			}

			if !unicode.IsDigit(ch) {
				break
			}
			n += size
		}

		if n == 0 {
			return nil, fmt.Errorf("invalid duration %s", lit)
		}

		magnitude, err := strconv.ParseInt(lit[:n], 10, 64)
		if err != nil {
			return nil, err
		}
		lit = lit[n:]

		n = 0
		for n < len(lit) {
			ch, size := utf8.DecodeRuneInString(lit[n:])
			if size == 0 {
				return nil, ErrInvalidRune
			}

			if !unicode.IsLetter(ch) {
				break
			}
			n += size
		}

		if n == 0 {
			return nil, fmt.Errorf("duration is missing a unit: %s", input)
		}

		unit := lit[:n]
		if unit == "µs" {
			unit = "us"
		}
		values = append(values, duration{
			magnitude: magnitude,
			unit:      unit,
		})
		lit = lit[n:]
	}
	return values, nil
}
