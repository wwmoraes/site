HUGO_ENVIRONMENT ?= development
HUGO_SERVER_PORT ?= 8888

HUGO_SOURCES = $(strip $(shell git ls-files {assets,config,content,data,i18n,layouts,static,themes}'/**/*.*'))

ifeq (${GIT_BRANCH},master)
HUGO_ENVIRONMENT = production
else ifeq (${GIT_BRANCH},staging)
HUGO_ENVIRONMENT = staging
else
HUGO_ENVIRONMENT = development
endif

.PHONY: start
#: Starts a local server.
start:
	hugo server -e ${HUGO_ENVIRONMENT} -p ${HUGO_SERVER_PORT}

#: Builds the static site pages.
public: ${HUGO_SOURCES}
	$(info generating static site at '$@'...)
	@hugo --gc --cleanDestinationDir
