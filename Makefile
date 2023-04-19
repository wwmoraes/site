DIAGRAMS_SOURCES = $(shell find content -type f -iname "*.puml")
DIAGRAMS_TARGETS = $(patsubst %.puml,%.png,${DIAGRAMS_SOURCES})

start:
	@exec hugo server -p 8888

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


metrics:
	@mkdir -p tmp/goatcounter/db
	docker run --rm -it \
		--name artero-goatcounter \
		-e GOATCOUNTER_DOMAIN=stats.artero.dev \
		-e GOATCOUNTER_EMAIL=webmaster@artero.dev \
		-e GOATCOUNTER_PASSWORD=webmaster \
		-e GOATCOUNTER_DEBUG=1 \
		-v ${PWD}/tmp/goatcounter/db:/goatcounter/db \
		-p 8080:8080 \
		baethon/goatcounter
.PRECIOUS: %.png
%.png: %.puml
	plantuml -tpng -darkmode -theme reddress-darkblue $<

index:
	@op run --env-file=.markscribe.env -- markscribe -write content/_index.md content/_index.md.tmpl

books:
	go run ./... update goodreads --list 138333248-william --shelf currently-reading
	go run ./... update goodreads --list 138333248-william --shelf read
	go run ./... update goodreads --list 138333248-william --shelf to-read
