#!/usr/bin/env bash
set -eo pipefail

declare -r GO_VERSION=1.17.1

# Hashes are from the table at https://golang.org/dl/
function go_hash () {
  case $1 in
    linux_amd64)
      echo dab7d9c34361dc21ec237d584590d72500652e7c909bf082758fb63064fca0ef
      ;;
    linux_arm64)
      echo 53b29236fa03ed862670a5e5e2ab2439a2dc288fe61544aa392062104ac0128c
      ;;
    mac)
      echo 3c452046b1dfa27b70d3217c9fe6de266f9fd74d83aad81382fead70efcdffca
      ;;
    windows)
      echo 2f2d0a5d7c59fb38fcacaf1e272cf701bb8c050300ba8b609fc30d2c5800f02e
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
