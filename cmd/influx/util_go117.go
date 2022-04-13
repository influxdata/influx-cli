//go:build !go1.18

// remove this backward compat shim once we either:
// a) upgrade cross-builder docker image in https://github.com/influxdata/edge/
// b) switch the CI system for this repo away from using cross-builder docker image.

package main

import "strings"

func stringsCut(s, sep string) (before, after string, found bool) {
	if i := strings.Index(s, sep); i >= 0 {
		return s[:i], s[i+len(sep):], true
	}
	return s, "", false
}
