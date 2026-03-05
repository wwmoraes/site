ROOT ?= $(shell git rev-parse --show-toplevel)
GOBIN ?= ${ROOT}/bin
COMPONENT = $(notdir $(realpath $(dir $(lastword ${MAKEFILE_LIST}))))

$(eval ${COMPONENT}_SOURCES = $(strip $(shell git ls-files '*.go') ${GO_SOURCES}) ${ROOT}/go.sum)

.PHONY: all
all: ${GOBIN}/${COMPONENT}

.PHONY: check
check: $(value ${COMPONENT}_SOURCES)
	golangci-lint run

.PHONY: clean
clean:
	rm -f ${GOBIN}/${COMPONENT}

.PHONY: test
test: $(value ${COMPONENT}_SOURCES)
	go test -v ./...

## make magic, not war :)

${GOBIN}/${COMPONENT}: $(value ${COMPONENT}_SOURCES)
	$(info building $(@F)...)
	@go build -ldflags='${LDFLAGS}' -o $@ .
	@touch $@
