GOFLAGS += -mod=readonly -race -trimpath
GO_SOURCES = $(addprefix ${ROOT}/,$(strip $(shell git ls-files 'pkg/*.go' 'internal/*.go')))
LDFLAGS += -s -w

check:: go.sum
	$(info linting golang sources...)
	@golangci-lint run

test::
	go test -v ./...

.PHONY: bin/%
bin/%: cmd/%
	@${MAKE} -C $<

################################################################################
# Make magic, not war :)
################################################################################

#: Update golang dependencies manually.
go.sum: GOFLAGS-=-mod=readonly
go.sum: ${GO_SOURCES} $(strip $(shell git ls-files 'cmd/*.go')) go.mod
	$(info updating golang dependencies...)
	@go mod tidy -v -x
	@touch $@

gomod2nix.toml: go.sum
	gomod2nix generate
