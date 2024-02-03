HUGO ?= hugo
HUGO_SERVER_PORT ?= 8888

.PHONY: start
start:
	@${HUGO} server -p ${HUGO_SERVER_PORT}

.PHONY: start-prod
start-prod:
	@${HUGO} server -e production -p 8888

public: data ${SITE_SOURCES}
	@${HUGO} --gc --cleanDestinationDir
