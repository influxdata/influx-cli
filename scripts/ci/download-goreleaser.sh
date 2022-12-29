#!/bin/bash
set -o nounset \
    -o errexit \
    -o pipefail

case ${1} in
v1.14.0)
  CHECKSUM=5f832026b88340318caaec5bd985951e7d363bd248bf49f25239ebb802304bcb
  ;;
v1.13.1)
  CHECKSUM=136fecfb2e2f3a7965274ad5e2571985d8b2fa724b6536874f082e4b0bb9f344
  ;;
v1.13.0)
  CHECKSUM=743dea6fa96f3acdf0fe99ce5f8c83f43afe72efedeb1506f88f5321a18f63f2
  ;;
*)
  printf 'Could not validate goreleaser version %s...\n' "${1}" 1>&2 ; exit 1
  ;;
esac

curl -LO "https://github.com/goreleaser/goreleaser/releases/download/${1}/goreleaser_Linux_x86_64.tar.gz"

printf '%s goreleaser_Linux_x86_64.tar.gz' "${CHECKSUM}" | sha256sum --check

tar -xf goreleaser_Linux_x86_64.tar.gz -C "${GOPATH}/bin"
