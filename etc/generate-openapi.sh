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
  --additional-properties packageName=api,enumClassPrefix=true,generateInterfaces=true

# Clean up files we don't care about.
(
  cd "${ROOT_DIR}/internal/api"
  rm -rf go.mod go.sum git_push.sh api docs .openapi-generator .travis.yml .gitignore
)

# Since we deleted the generated go.mod, run `go mod tidy` to update parent dependencies.
(
  cd "${ROOT_DIR}"
  go mod tidy
)

# Inject linter directives into generated files to make staticcheck happy.
(
  cd "${ROOT_DIR}"
  cat <<EOF > internal/api/client.go
//lint:file-ignore ST1005 Ignore capitalized errors, they're generated
//lint:file-ignore SA6005 Ignore old-fashioned way of comparing strings, it's generated

$(cat internal/api/client.go)
EOF

  cat <<EOF > internal/api/configuration.go
//lint:file-ignore ST1005 Ignore capitalized errors, they're generated

$(cat internal/api/configuration.go)
EOF
)
