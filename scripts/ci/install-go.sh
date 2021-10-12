#!/usr/bin/env bash
set -eo pipefail

declare -r GO_VERSION=1.17.2

# Hashes are from the table at https://golang.org/dl/
function go_hash () {
  case $1 in
    linux_amd64)
      echo f242a9db6a0ad1846de7b6d94d507915d14062660616a61ef7c808a76e4f1676
      ;;
    linux_arm64)
      echo a5a43c9cdabdb9f371d56951b14290eba8ce2f9b0db48fb5fc657943984fd4fc
      ;;
    mac)
      echo 7914497a302a132a465d33f5ee044ce05568bacdb390ab805cb75a3435a23f94
      ;;
    windows)
      echo fa6da0b829a66f5fab7e4e312fd6aa1b2d8f045c7ecee83b3d00f6fe5306759a
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
