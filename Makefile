### Environment setup
export GOPATH=$(shell go env GOPATH)
export GOOS=$(shell go env GOOS)
export GOARCH=$(shell go env GOARCH)
export GOVERSION=$(shell go list -m -f '{{.GoVersion}}')

ifeq ($(OS), Windows_NT)
	VERSION := $(shell git describe --exact-match --tags 2>nil)
else
	VERSION := $(shell git describe --exact-match --tags 2>/dev/null)
endif
COMMIT := $(shell git rev-parse --short HEAD)

LDFLAGS := $(LDFLAGS) -X main.commit=$(COMMIT) -X main.date=$(shell date -u +'%Y-%m-%dT%H:%M:%SZ')
ifdef VERSION
	LDFLAGS += -X main.version=$(VERSION)
endif
export GO_BUILD=go build -ldflags "$(LDFLAGS)"

SOURCES := $(shell find . -name '*.go' -not -name '*_test.go') go.mod go.sum
SOURCES_NO_VENDOR := $(shell find . -path ./vendor -prune -o -name "*.go" -not -name '*_test.go' -print)

# Allow for `go test` to be swapped out by other tooling, i.e. `gotestsum`
export GO_TEST=go test
# Allow for a subset of tests to be specified.
GO_TEST_PATHS=./...

### Build / dependency management
api/api.gen.go: api/api.yml api/gen.go
	go generate ./api

openapi: api/api.gen.go

fmt: $(SOURCES_NO_VENDOR)
	gofmt -w -s $^
	go run github.com/daixiang0/gci -w $^

bin/$(GOOS)/influx: $(SOURCES)
	$(GO_BUILD) -o $@ ./cmd/$(shell basename "$@")

influx: bin/$(GOOS)/influx

vendor: go.mod go.sum
	go mod vendor

build: openapi fmt influx

clean:
	$(RM) -r bin
	$(RM) -r vendor

### Linters
checkfmt:
	./etc/checkfmt.sh

checktidy:
	./etc/checktidy.sh

staticcheck: $(SOURCES) vendor
	go run honnef.co/go/tools/cmd/staticcheck -go $(GOVERSION) ./...

vet:
	go vet ./...

# Testing
test:
	$(GO_TEST) $(GO_TEST_PATHS)

test-race:
	$(GO_TEST) -v -race -count=1 $(GO_TEST_PATHS)

### List of all targets that don't produce a file
.PHONY: influx openapi fmt build checkfmt checktidy staticcheck vet test test-race
