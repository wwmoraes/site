MAKEFLAGS += --no-builtin-rules
MAKEFLAGS += --no-builtin-variables
MAKEFLAGS += --jobs --max-load

SHELL := $(shell which bash)
.SHELLFLAGS := -euo pipefail -c

-include .env
-include .env.local

define PGP_EMAILS
$(strip
william@artero.dev
contact@artero.dev
security@artero.dev
git@artero.dev
)
endef

-include .make/*.mk

HUGO_SOURCES += ${EXIF_TARGETS} ${PGP_TARGETS} ${PLANTUML_TARGETS} content/radar/radar.svg static/favicon.ico

export

.SUFFIXES:

#: Builds the entire project.
all: bin/site public

.PHONY: clean
#: Delete all files that are normally created by building.
clean:
	rm -rf bin public dist resources content/radar/radar.svg

#: Generates site for the current branch.
dist: dist/${HUGO_ENVIRONMENT}

.PHONY: deploy
#: Deploys the site version for the current environment.
deploy: dist/${HUGO_ENVIRONMENT}
	$(info deploying '$<' to '${HUGO_ENVIRONMENT}' environment...)
	@nix run nixpkgs#wrangler -- pages deploy $< \
		--no-bundle \
		--upload-source-maps \
		--project-name '${CLOUDFLARE_PROJECT_NAME}' \
		--branch '${GIT_BRANCH}' \
		--commit-hash '${GIT_COMMIT_HASH}' \
		--commit-message '${GIT_COMMIT_MESSAGE}' \
		--commit-dirty '${GIT_DIRTY}' \
		;

.PHONY: purge-cache
purge-cache:
ifeq (${CLOUDFLARE_API_TOKEN},)
	$(warning no cloudflare API token set, skipping purge)
else
	@curl -sS 'https://api.cloudflare.com/client/v4/zones/${CLOUDFLARE_ZONE}/purge_cache' \
		-H 'Content-Type: application/json' \
		-H 'Authorization: Bearer ${CLOUDFLARE_API_TOKEN}' \
		-w 'HTTP_STATUS:%{http_code}'
endif

#: Perform self-checks such as linting and formatting.
check: vale-check
check: golang-check
check: css-check
check:
	$(info linting editor config constraints...)
	@editorconfig-checker

.PHONY: test
#: Builds and runs tests.
test: golang-test

## make magic, not war :)

#: Builds the CLI companion tool to manage the site.
bin/site: golang-build/bin/site
	@touch $@

golang-build/bin/site: cmd/site ${GO_SOURCES} go.sum

dist/%: ${HUGO_SOURCES}
	$(info generating static site for environment '$*' at '$@'...)
	@hugo --gc --cleanDestinationDir --environment '$*' --destination '$@'
