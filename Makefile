# Disable built-in rules and variables and suffixes
MAKEFLAGS += --no-builtin-rules
MAKEFLAGS += --no-builtin-variables
.SUFFIXES:

-include .env
-include .env.local
export

SITE = bin/site
SITE_SOURCES = $(shell git ls-files {assets,config,content,data,i18n,layouts,static}/* themes/*/{assets,data,i18n,layouts,static}/* themes/*/hugo.*)

-include .make/*.mk

.PHONY: all
all: build

.PHONY: build
build: public

.PHONY: clean
clean:
	@${RM} -r bin
	@${RM} -r public

.PHONY: test
test:
	@${GO} test ./... | column -t
