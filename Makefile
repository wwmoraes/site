MAKEFLAGS += --no-builtin-rules
MAKEFLAGS += --no-builtin-variables
MAKEFLAGS += --jobs --max-load

SHELL := $(shell which bash)
.SHELLFLAGS := -euo pipefail -c
.DEFAULT_GOAL := all

-include .env
-include .env.local

ROOT ?= $(shell git rev-parse --show-toplevel)

-include .make/*.mk
-include .meta/make/*.mk

export

.PHONY: all
#: Builds the entire project.
all: bin/site dist

#: Builds the CLI companion tool to manage the site.
bin/site:

.PHONY: clean
#: Delete all files that are normally created by building.
clean::
	-@${MAKE} -C cmd/site clean
	rm -rf dist resources

.PHONY: check
#: Lints and validates sources.
check::

.PHONY: test
#: Builds and runs tests.
test::
