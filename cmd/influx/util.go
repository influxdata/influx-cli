//go:build go1.18

package main

import "strings"

func stringsCut(s, sep string) (before, after string, found bool) {
	return strings.Cut(s, sep)
}
