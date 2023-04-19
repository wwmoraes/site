DIAGRAMS_SOURCES = $(shell find content -type f -iname "*.puml")
DIAGRAMS_TARGETS = $(patsubst %.puml,%.png,${DIAGRAMS_SOURCES})
GOODREADS_LIST = 138333248-william
GOODREADS_SHELVES = currently-reading read to-read

start:
	@hugo server -p 8888

start-prod:
	@hugo server -e production

build:
	@rm -rf public
	@hugo --gc --cleanDestinationDir

hook-install:
	@pre-commit install

hook-update:
	@pre-commit autoupdate

check:
	@pre-commit run --all-files

clean:
	${RM} -r public

diagrams: ${DIAGRAMS_TARGETS}

.PRECIOUS: %.png
%.png: %.puml
	plantuml -tpng -darkmode -theme reddress-darkblue $<

index:
	@op run --env-file=.markscribe.env -- markscribe -write content/_index.md content/_index.md.tmpl

books: $(addprefix data/goodreads/,${GOODREADS_SHELVES})

.PRECIOUS: data/goodreads/%
.PHONY: data/goodreads/%
data/goodreads/%:
	@echo go run ./... update goodreads --list ${GOODREADS_LIST} --shelf $*
