# Disable built-in rules and variables and suffixes
MAKEFLAGS += --no-builtin-rules
MAKEFLAGS += --no-builtin-variables
.SUFFIXES:

-include .env
-include .env.local
export

SITE = bin/site

-include .make/*.mk

.PHONY: all
all: build

.PHONY: build
build: diagrams exif radar favicon
	@${HUGO} --gc --cleanDestinationDir

.PHONY: clean
clean:
	@${RM} -r bin
	@${RM} -r public

.PHONY: test
test:
	@${GO} test ./... | column -t

.PHONY: publish
publish: data build

.PHONY: spell
spell: .styles
	@vale .

.styles: .vale.ini
	@vale sync
