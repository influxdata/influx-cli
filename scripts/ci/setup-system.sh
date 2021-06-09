#!/usr/bin/env bash
set -eo pipefail

function setup_linux () {
    sudo apt-get update
    sudo apt-get install -y --no-install-recommends make
}

function setup_mac () {
    # Python and TCL both come pre-installed on Circle's mac executors, and both depend on wget in some way.
    # Homebrew will auto-upgrade both of them when wget is installed/upgraded, triggering a chain of upgrades.
    # Uninstall them both before adding wget to avoid burning time in CI for things we don't need.
    brew remove --force python@3.9 tcl-tk
    HOMEBREW_NO_AUTO_UPDATE=1 brew install wget
}

function setup_windows () {
    choco install make mingw wget
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
