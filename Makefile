### Environment setup
export GOPATH=$(shell go env GOPATH)
export GOOS=$(shell go env GOOS)
export GOARCH=$(shell go env GOARCH)
export GOVERSION=$(shell go list -m -f '{{.GoVersion}}')

LDFLAGS := $(LDFLAGS) -X main.date=$(shell date -u +'%Y-%m-%dT%H:%M:%SZ')

ifdef VERSION
	LDFLAGS += -X main.version=$(VERSION)
endif
ifdef COMMIT
	LDFLAGS += -X main.commit=$(COMMIT)
endif

# Use default flags, but allow adding -gcflags "..." if desired. Eg, for debug
# builds, may want to use GCFLAGS="all=-N -l" in the build environment.
GCFLAGS ?=
ifneq ($(GCFLAGS),)
GCFLAGS := -gcflags "$(GCFLAGS)"
endif

export GO_BUILD=go build $(GCFLAGS) -ldflags "$(LDFLAGS)"

# SOURCES are the files that affect building the main binary.
SOURCES := $(shell find . -name '*.go' -not -name '*_test.go') go.mod go.sum

# FMT_FILES are all files that should be formatted according to our rules.
FMT_FILES := $(shell find . -path ./vendor -prune -o -name "*.go" -print)

# Allow for `go test` to be swapped out by other tooling, i.e. `gotestsum`
export GO_TEST=go test
# Allow for a subset of tests to be specified.
GO_TEST_PATHS=./...

### Build / dependency management
openapi:
	./etc/generate-openapi.sh

fmt: $(FMT_FILES)
	# Format everything, but the import-format doesn't match our desired pattern.
	gofmt -w -s $^
	# Remove unused imports.
	go run golang.org/x/tools/cmd/goimports -w $^
	# Format imports.
	go run github.com/daixiang0/gci -w $^

bin/$(GOOS)/influx: $(SOURCES)
	CGO_ENABLED=0 $(GO_BUILD) -o $@ ./cmd/$(shell basename "$@")

.DEFAULT_GOAL := influx
influx: bin/$(GOOS)/influx

clean:
	$(RM) -r bin

### Linters
checkfmt:
	./etc/checkfmt.sh

checktidy:
	./etc/checktidy.sh

checkopenapi:
	./etc/checkopenapi.sh

staticcheck: $(SOURCES)
	go run honnef.co/go/tools/cmd/staticcheck -go $(GOVERSION) ./...

vet:
	go vet ./...

# Testing
mock: ./internal/mock/gen.go
	go generate ./internal/mock/

test:
	CGO_ENABLED=0 $(GO_TEST) $(GO_TEST_PATHS)

test-race:
	# Race-checking requires CGO.
	$(GO_TEST) -v -race -count=1 $(GO_TEST_PATHS)

### List of all targets that don't produce a file
.PHONY: influx openapi fmt checkfmt checktidy staticcheck vet mock test test-race
