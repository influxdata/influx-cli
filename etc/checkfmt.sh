#!/bin/bash

HAS_FMT_ERR=0
# For every Go file in the project, excluding vendor...
for file in $(go list -f '{{$dir := .Dir}}{{range .GoFiles}}{{printf "%s/%s\n" $dir .}}{{end}}' ./...); do
  # ... if file does not contain standard generated code comment (https://golang.org/s/generatedcode)...
  if ! grep -Exq '^// Code generated .* DO NOT EDIT\.$' $file; then
    FMT_OUT="$(gofmt -l -d -e $file)" # gofmt exits 0 regardless of whether it's formatted.
    GCI_OUT="$(go run github.com/daixiang0/gci -d $file)"

    # Work around annoying output of gci
    if [[ "$GCI_OUT" = "skip file $file since no import" ]]; then
      GCI_OUT=""
    fi

    if [[ -n "$FMT_OUT" || -n "$GCI_OUT" ]]; then
      HAS_FMT_ERR=1
      echo "Not formatted: $file"
    fi
  fi
done

if [ "$HAS_FMT_ERR" -eq "1" ]; then
    echo 'Commit includes files that are not formatted' && \
    echo 'run "make fmt"' && \
    echo ''
fi
exit "$HAS_FMT_ERR"
