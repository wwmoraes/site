GOFLAGS += -mod=readonly -race -trimpath
GO_SOURCES = $(addprefix ${ROOT}/,$(strip $(shell git ls-files 'pkg/*.go' 'internal/*.go')))
LDFLAGS += -s -w

#: Update golang dependencies manually.
go.sum: GOFLAGS-=-mod=readonly
go.sum: ${GO_SOURCES} $(strip $(shell git ls-files 'cmd/*.go')) go.mod
	$(info updating golang dependencies...)
	@go mod tidy -v -x
	@touch $@

gomod2nix.toml: go.sum
	gomod2nix generate
