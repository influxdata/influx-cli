#!/usr/bin/env bash
set -exo pipefail

function build_test_binaries () {
    local tags=""
    if [ "$(go env GOOS)" = windows ]; then
        tags="-tags timetzdata"
    fi

    CGO_ENABLED=0 go-test-compile ${tags} -o "${1}/" ./...
}

function build_test_tools () {
    # Copy pre-built gotestsum out of the cross-builder.
    local ext=""
    if [ "$(go env GOOS)" = windows ]; then
        ext=".exe"
    fi
    cp "/usr/local/bin/gotestsum_$(go env GOOS)_$(go env GOARCH)${ext}" "$1/gotestsum${ext}"

    # Build test2json from the installed Go distribution.
    CGO_ENABLED=0 go build -o "${1}/" -ldflags="-s -w" cmd/test2json
}

function write_test_metadata () {
    # Write version that should be reported in test results.
    echo "$(go env GOVERSION) $(go env GOOS)/$(go env GOARCH)" > "${1}/go.version"

    # Write list of all packages.
    go list ./... > "${1}/tests.list"
}

function main () {
    if [[ $# != 1 ]]; then
        >&2 echo Usage: $0 '<output-dir>'
        >&2 echo '<output-dir>' will be created if it does not already exist
        exit 1
    fi
    local -r out_dir="$1"

    mkdir -p "$out_dir"

    # Build all test binaries.
    build_test_binaries "${out_dir}/"
    # Build gotestsum and test2json so downstream jobs can use it without needing `go`.
    build_test_tools "$out_dir"
    # Write other metadata needed for testing.
    write_test_metadata "$out_dir"
}

main ${@}
