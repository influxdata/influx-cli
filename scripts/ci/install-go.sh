#!/usr/bin/env bash
set -eo pipefail

declare -r GO_VERSION=1.19.4

# Hashes are from the table at https://golang.org/dl/
function go_hash () {
  case $1 in
    linux_amd64)
      echo c9c08f783325c4cf840a94333159cc937f05f75d36a8b307951d5bd959cf2ab8
      ;;
    linux_arm64)
      echo 9df122d6baf6f2275270306b92af3b09d7973fb1259257e284dba33c0db14f1b
      ;;
    mac)
      echo 44894862d996eec96ef2a39878e4e1fce4d05423fc18bdc1cbba745ebfa41253
      ;;
    windows)
      echo ada490e188bfb57c7388da7c5eba7565390992b6496204d30e710d37755956b0
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
