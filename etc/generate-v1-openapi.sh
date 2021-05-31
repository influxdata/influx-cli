#!/usr/bin/env bash
set -euo pipefail

declare -r ETC_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" &>/dev/null && pwd)"
declare -r ROOT_DIR="$(dirname ${ETC_DIR})"
declare -r API_DIR="${ROOT_DIR}/api/v1"

declare -r GENERATED_PATTERN='^// Code generated .* DO NOT EDIT\.$'
declare -r GENERATOR_DOCKER_IMG=openapitools/openapi-generator-cli:v5.1.0

# Clean up all the generated files in the target directory.
rm -f $(grep -Elr "${GENERATED_PATTERN}" "${API_DIR}")

# cleanup is responsible for normalizing generated files.
function cleanup() {
  # Clean up files we don't care about.
  cd "$1"
  rm -rf go.mod go.sum git_push.sh api docs .openapi-generator .travis.yml .gitignore

  # Change extension of generated files.
  for f in $(grep -El "${GENERATED_PATTERN}" *.go); do
    base=$(basename ${f} .go)
    mv ${f} ${base}.gen.go
  done

  # Clean up the generated code.
  cd "${ROOT_DIR}"
  make >/dev/null fmt
}

# Run the generator for the v1 compatibility API - This produces many more files than we want to track in git.
docker run --rm -it -u "$(id -u):$(id -g)" \
  -v "${ROOT_DIR}/api":/api \
  ${GENERATOR_DOCKER_IMG} \
  generate \
  -g go \
  -i /api/contract/swaggerV1Compat.yml \
  -o /api/v1 \
  -t /api/templates \
  --global-property modelTests=false,modelDocs=false,apiTests=false,apiDocs=false \
  --additional-properties packageName=v1,enumClassPrefix=true,generateInterfaces=true

cleanup "$API_DIR"