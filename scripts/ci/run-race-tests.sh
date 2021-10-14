#!/usr/bin/env bash
set -exo pipefail

function main () {
    if [[ $# != 1 ]]; then
        >&2 echo Usage: $0 '<output-dir>'
        exit 1
    fi
    if [[ $(go env GOOS) != linux || $(go env GOARCH) != amd64 ]]; then
        >&2 echo Race tests only supported on linux/amd64
        exit 1
    fi

    local -r out_dir="$1"
    mkdir -p "$out_dir"

    gotestsum --junitfile "${out_dir}/report.xml" -- -race ./...
}

main ${@}
