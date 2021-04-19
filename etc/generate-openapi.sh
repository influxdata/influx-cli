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

# Clean up files we don't care about.
(
  cd "${ROOT_DIR}/internal/api"
  rm -rf go.mod go.sum git_push.sh api docs .openapi-generator .travis.yml .gitignore
)

# Edit the generated files.
(
  # Replace all uses of int32 with int64 in our generated models.
  #
  # See https://github.com/OpenAPITools/openapi-generator/issues/9280 for a feature request
  # to make the Go generator's number-handling compliant with the spec, so we can generate int64
  # fields directly without making our public docs invalid.
  for m in internal/api/model_*.go; do
    sed -i.bak -e 's/int32/int64/g' ${m}
  done

  rm internal/api/*.bak
)

# Since we deleted the generated go.mod, run `go mod tidy` to update parent dependencies.
(
  cd "${ROOT_DIR}"
  make fmt
  go mod tidy
)
