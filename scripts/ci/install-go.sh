#!/usr/bin/env bash
set -eo pipefail

declare -r GO_VERSION=1.16.5

function install_go_linux () {
    ARCHIVE=go${GO_VERSION}.linux-$(dpkg --print-architecture).tar.gz
    wget https://golang.org/dl/${ARCHIVE}
    tar -C $1 -xzf ${ARCHIVE}
    rm ${ARCHIVE}
}

function install_go_mac () {
    ARCHIVE=go${GO_VERSION}.darwin-amd64.tar.gz
    wget https://golang.org/dl/${ARCHIVE}
    tar -C $1 -xzf ${ARCHIVE}
    rm ${ARCHIVE}
}

function install_go_windows () {
    ARCHIVE=go${GO_VERSION}.windows-amd64.zip
    wget https://golang.org/dl/${ARCHIVE}
    unzip -qq -d $1 ${ARCHIVE}
    rm ${ARCHIVE}
}

function main () {
    if [[ $# != 1 ]]; then
        >&2 echo Usage: $0 '<install-dir>'
        exit 1
    fi
    local -r install_dir=$1

    rm -rf "$install_dir"
    mkdir -p "$install_dir"
    case $(uname) in
        Linux)
            install_go_linux "$install_dir"
            ;;
        Darwin)
            install_go_mac "$install_dir"
            ;;
        MSYS_NT*)
            install_go_windows "$install_dir"
            ;;
        *)
            >&2 echo Error: unknown OS $(uname)
            exit 1
            ;;
    esac

    "${install_dir}/go/bin/go" version
}

main ${@}
