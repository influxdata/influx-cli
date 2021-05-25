#!/usr/bin/env bash
set -e

declare -r ETC_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" &> /dev/null && pwd)"
${ETC_DIR}/generate-openapi.sh

if ! git --no-pager diff --exit-code -- api; then
  >&2 echo "openapi generated client doesn't match spec, please run 'make openapi'"
  exit 1
fi
