#!/usr/bin/env bash

declare -r ETC_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" &> /dev/null && pwd)"
declare -r ROOT_DIR="$(dirname ${ETC_DIR})"

declare -r GENERATOR_DOCKER_IMG=openapitools/openapi-generator-cli:v5.1.0

# Run the generator - This produces many more files than we want to track in git.
docker run --rm -it -u "$(id -u):$(id -g)" \
  -v "${ROOT_DIR}":/influx \
  ${GENERATOR_DOCKER_IMG} \
  generate \
  -g go \
  -i /influx/internal/api/api.yml \
  -o /influx/internal/api \
  -t /influx/internal/api/templates \
  --additional-properties packageName=api,enumClassPrefix=true,generateInterfaces=true

# Edit the generated files.
(
  # Clean up files we don't care about.
  cd "${ROOT_DIR}/internal/api"
  rm -rf go.mod go.sum git_push.sh api docs .openapi-generator .travis.yml .gitignore

  # Clean up the generated code.
  cd "${ROOT_DIR}"
  make fmt
)
