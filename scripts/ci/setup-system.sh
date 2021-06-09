#!/usr/bin/env bash
set -eo pipefail

function setup_linux () {
    sudo apt-get update
    sudo apt-get install -y --no-install-recommends make
}

function setup_mac () {
    # Placeholder in case we ever need to `brew install` something
    return
}

function setup_windows () {
    choco install make wget
}

function main () {
    case $(uname) in
        Linux)
            setup_linux
            ;;
        Darwin)
            setup_mac
            ;;
        MSYS_NT*)
            setup_windows
            ;;
        *)
            >&2 echo Error: unknown OS $(uname)
            exit 1
            ;;
    esac
}

main ${@}
