#!/usr/bin/env bash
set -euo pipefail

declare -r ETC_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" &> /dev/null && pwd)"
declare -r ROOT_DIR="$(dirname ${ETC_DIR})"
declare -r API_DIR="${ROOT_DIR}/api"

declare -r GENERATED_PATTERN='^// Code generated .* DO NOT EDIT\.$'
declare -r MERGE_DOCKER_IMG=quay.io/influxdb/swagger-cli
declare -r GENERATOR_DOCKER_IMG=openapitools/openapi-generator-cli:v5.1.0
declare -r TAG_STRIP_IMG=python:3.9-alpine3.15

# Clean up all the generated files in the target directory.
rm -f $(grep -Elr "${GENERATED_PATTERN}" "${API_DIR}")

# Merge all API contracts into a single file to drive codegen.
docker run --rm -it -u "$(id -u):$(id -g)" \
  -v "${API_DIR}":/api \
  ${MERGE_DOCKER_IMG} \
  swagger-cli bundle /api/contract/cli.yml \
  --outfile /api/cli.gen.yml \
  --type yaml

# Merge extras to a separate file
docker run --rm -it -u "$(id -u):$(id -g)" \
  -v "${API_DIR}":/api \
  ${MERGE_DOCKER_IMG} \
  swagger-cli bundle /api/contract/cli-extras.yml \
  --outfile /api/cli-extras.gen.yml \
  --type yaml

# Strip certain tags to prevent duplicated and conflicting codegen.
docker run --rm -it -u "$(id -u):$(id -g)" \
  -v "${ROOT_DIR}":/api \
  ${TAG_STRIP_IMG} \
  sh -c "python3 /api/etc/stripGroupTags.py /api/api/cli.gen.yml > /api/api/cli-stripped.gen.yml"

# Run the generator - This produces many more files than we want to track in git.
docker run --rm -it -u "$(id -u):$(id -g)" \
  -v "${API_DIR}":/api \
  ${GENERATOR_DOCKER_IMG} \
  generate \
  -g go \
  -i /api/cli-stripped.gen.yml \
  -o /api \
  -t /api/templates \
  --additional-properties packageName=api,enumClassPrefix=true,generateInterfaces=true

# Run the generator for extras
docker run --rm -it -u "$(id -u):$(id -g)" \
  -v "${API_DIR}":/api \
  ${GENERATOR_DOCKER_IMG} \
  generate \
  -g go \
  -i /api/cli-extras.gen.yml \
  -o /api/extras \
  -t /api/templates \
  --additional-properties packageName=extras,enumClassPrefix=true,generateInterfaces=true

rm ${API_DIR}/extras/{client.go,configuration.go,response.go}

# Edit the generated files.
for DIR in "${API_DIR}" "${API_DIR}/extras" ; do
(
  # Clean up files we don't care about.
  cd $DIR
  rm -rf go.mod go.sum git_push.sh api docs .openapi-generator .travis.yml .gitignore

  # Change extension of generated files.
  for f in $(grep -El "${GENERATED_PATTERN}" *.go); do
    base=$(basename ${f} .go)
    mv ${f} ${base}.gen.go
  done
)
done

# Clean up the generated code.
cd "${ROOT_DIR}"
>/dev/null make fmt
