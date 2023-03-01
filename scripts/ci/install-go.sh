#!/usr/bin/env bash
set -eo pipefail

declare -r GO_VERSION=1.19.6

# Hashes are from the table at https://golang.org/dl/
function go_hash () {
  case $1 in
    linux_amd64)
      echo e3410c676ced327aec928303fef11385702a5562fd19d9a1750d5a2979763c3d
      ;;
    linux_arm64)
      echo e4d63c933a68e5fad07cab9d12c5c1610ce4810832d47c44314c3246f511ac4f
      ;;
    mac)
      echo 63386d51c69cef6c001ff0436832289635ba4a2649282451a18827e93507b444
      ;;
    windows)
      echo 8d84af29e46c38b1eec77f9310310517c9e394ac7489e1c7329a94b443b0388d
      ;;
  esac
}

function install_go_linux () {
    local -r arch=$(dpkg --print-architecture)
    ARCHIVE=go${GO_VERSION}.linux-${arch}.tar.gz
    wget https://golang.org/dl/${ARCHIVE}
    echo "$(go_hash linux_${arch})  ${ARCHIVE}" | sha256sum --check --
    tar -C $1 -xzf ${ARCHIVE}
    rm ${ARCHIVE}
}

function install_go_mac () {
    ARCHIVE=go${GO_VERSION}.darwin-amd64.tar.gz
    wget https://golang.org/dl/${ARCHIVE}
    echo "$(go_hash mac)  ${ARCHIVE}" | shasum -a 256 --check -
    tar -C $1 -xzf ${ARCHIVE}
    rm ${ARCHIVE}
}

function install_go_windows () {
    ARCHIVE=go${GO_VERSION}.windows-amd64.zip
    wget https://golang.org/dl/${ARCHIVE}
    echo "$(go_hash windows)  ${ARCHIVE}" | sha256sum --check --
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
