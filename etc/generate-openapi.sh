#!/usr/bin/env bash

declare -r ETC_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" &> /dev/null && pwd)"
declare -r ROOT_DIR="$(dirname ${ETC_DIR})"
declare -r CODEGEN_TARGET_DIR="${ROOT_DIR}/internal/api"

declare -r GENERATED_PATTERN='^// Code generated .* DO NOT EDIT\.$'
declare -r MERGE_DOCKER_IMG=quay.io/influxdb/swagger-cli
declare -r GENERATOR_DOCKER_IMG=openapitools/openapi-generator-cli:v5.1.0

# Clean up all the generated files in the target directory.
rm $(grep -Elr "${GENERATED_PATTERN}" "${CODEGEN_TARGET_DIR}")

# Merge all API contracts into a single file to drive codegen.
docker run --rm -it -u "$(id -u):$(id -g)" \
  -v "${ROOT_DIR}":/influx \
  ${MERGE_DOCKER_IMG} \
  swagger-cli bundle /influx/api-contract/cli.yml \
  --outfile /influx/internal/api/cli.yml \
  --type yaml

# Run the generator - This produces many more files than we want to track in git.
docker run --rm -it -u "$(id -u):$(id -g)" \
  -v "${CODEGEN_TARGET_DIR}":/api \
  ${GENERATOR_DOCKER_IMG} \
  generate \
  -g go \
  -i /api/cli.yml \
  -o /api \
  -t /api/templates \
  --additional-properties packageName=api,enumClassPrefix=true,generateInterfaces=true

# Edit the generated files.
(
  # Clean up files we don't care about.
  cd "${CODEGEN_TARGET_DIR}"
  rm -rf go.mod go.sum git_push.sh api docs .openapi-generator .travis.yml .gitignore

  # Change extension of generated files.
  for f in $(grep -El "${GENERATED_PATTERN}" *.go); do
    base=$(basename ${f} .go)
    mv ${f} ${base}.gen.go
  done

  # Clean up the generated code.
  cd "${ROOT_DIR}"
  >/dev/null make fmt
)
