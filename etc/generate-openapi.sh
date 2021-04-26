#!/usr/bin/env bash

declare -r ETC_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" &> /dev/null && pwd)"
declare -r ROOT_DIR="$(dirname ${ETC_DIR})"

declare -r GENERATOR_DOCKER_IMG=openapitools/openapi-generator-cli:v5.1.0
declare -r OPENAPI_COMMIT=5ab3ef0b9a6aee68b3b34e1858ca6d55c153650a

# Download our target API spec.
# NOTE: openapi-generator supports HTTP references to API docs, but using that feature
# causes the host of the URL to be injected into the base paths of generated code.
curl -o ${ROOT_DIR}/internal/api/cli.yml https://raw.githubusercontent.com/influxdata/openapi/${OPENAPI_COMMIT}/contracts/cli.yml

# Run the generator - This produces many more files than we want to track in git.
docker run --rm -it -u "$(id -u):$(id -g)" \
  -v "${ROOT_DIR}":/influx \
  ${GENERATOR_DOCKER_IMG} \
  generate \
  -g go \
  -i /influx/internal/api/cli.yml \
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
  >/dev/null make fmt
)
