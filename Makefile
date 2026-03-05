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

export

define HUGO_SOURCES
$(strip
$(shell git ls-files {assets,config,content,data,i18n,layouts,static,themes}'/**/*.*')
$(patsubst %.json,%,$(strip $(shell git ls-files {archetypes,assets,content}'/**/*.'{jpg,png}'.json')))
$(patsubst %.puml,%.png,$(strip $(shell git ls-files 'content/**/*.puml')))
static
)
endef

#: Builds the entire project.
all: bin/site dist

#: Builds the CLI companion tool to manage the site.
bin/site:

.PHONY: clean
#: Delete all files that are normally created by building.
clean:
	-@${MAKE} -C cmd/site clean
	rm -rf dist resources

#: Generates all environments' site assets.
dist: dist/development dist/staging dist/production

#: Generates development site assets.
dist/development:

#: Generates staging site assets.
dist/staging:

#: Generates production site assets.
dist/production:

#: Perform self-checks such as linting and formatting.
check: check/prose
check: check/go
check: check/css
check:
	$(info linting editor config constraints...)
	@editorconfig-checker

.PHONY: check/css
check/css:
	$(info linting CSS...)
	@stylelint --allow-empty-input --cache --formatter compact '**.css' '**.scss'

.PHONY: check/git
check/git:
	$(info checking if commit messages follow conventions...)
	@cog check --from-latest-tag > /dev/null

.PHONY: check/go
check/go:
	$(info linting golang sources...)
	@golangci-lint run

.PHONY: check/prose
check/prose: .styles
	$(info linting prose...)
	@vale --minAlertLevel=error content

.PHONY: hooks/*
hooks/pre-commit: check
hooks/pre-push: all test check/git

.PHONY: test
#: Builds and runs tests.
test:
	go test -v ./...

## make magic, not war :)

.PRECIOUS: .styles
.SECONDARY: .styles
.styles: .vale.ini
	$(info updating prose linting settings...)
	@vale sync

.PHONY: bin/%
bin/%: cmd/%
	@${MAKE} -C $<

.PRECIOUS: content/%.png
content/%.png: content/%.puml
	$(info generating $@...)
	@kroki convert -o $@ $<

dist/%: ${HUGO_SOURCES}
	$(info generating static site for environment '$*' at '$@'...)
	@hugo --gc --cleanDestinationDir --environment '$*' --destination '$@'

.PHONY: static
#: Updates static assets.
static:
	@${MAKE} -C $@

.PRECIOUS: %.png
%.png: %.png.json bin/site
	$(info updating EXIF of $@...)
	@site image update $@ > /dev/null
	@touch $@

.PRECIOUS: %.jpg
%.jpg: %.jpg.json bin/site
	$(info updating EXIF of $@...)
	@site image update $@ > /dev/null
	@touch $@
