GOFLAGS += -mod=readonly -race -trimpath
LDFLAGS += -s -w

GO_SOURCES = $(strip $(shell git ls-files '*.go'))

.PHONY: golang-test
golang-test:
	go test -v ./...

.PHONY: golang-check
golang-check:
	$(info linting golang sources...)
	@golangci-lint run

.PHONY: golang-fix
#: Automatically fixes known Golang code issues.
golang-fix:
	golangci-lint run --fix --enable-only gci,gofumpt {{ .CLI_ARGS }}

.PHONY: golang-build/%
.SECONDARY: golang-build/%
golang-build/%:
	$(info building $*...)
	@go build -ldflags='${LDFLAGS}' -o ./$* ./$<

go.sum: GOFLAGS-=-mod=readonly
go.sum: ${GO_SOURCES} go.mod
	$(info updating golang dependencies...)
	@go mod tidy -v -x
	@touch $@

gomod2nix.toml: go.sum
	gomod2nix generate
