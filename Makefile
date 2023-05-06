DIAGRAMS_SOURCES = $(shell find content -type f -name "*.puml")
DIAGRAMS_TARGETS = $(patsubst %.puml,%.png,${DIAGRAMS_SOURCES})
GOODREADS_LIST = 138333248-william
GOODREADS_SHELVES = currently-reading,read,to-read

IMAGES = $(wildcard archetypes/*/*.jpg)
IMAGES += $(wildcard assets/images/*.jpg)
IMAGES += $(wildcard content/posts/*/*.jpg)

start:
	@hugo server -p 8888

start-prod:
	@hugo server -e production

build:
	@${RM} -r public
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

github:
	@op run --env-file=.env -- go run ./... update github

books:
	@go run ./... update goodreads --list ${GOODREADS_LIST} --shelves ${GOODREADS_SHELVES}

## EXIF tags management
## https://exiftool.org/examples.html
## https://exiftool.org/TagNames/EXIF.html

exif: ${IMAGES}

exif-show: ${IMAGES}
	@$(foreach IMAGE,$^,exiftool -s \
		-Directory \
		-FileName \
		-Copyright \
		-ImageUniqueID \
		-ImageDescription \
		-Make \
		${IMAGE};echo "---";)

%.jpg: %.jpg.json ; ${buildExif}

define buildExif
$(info adjusting EXIF of $@)
@exiftool -overwrite_original_in_place -"*=" -j=$< $@
endef

define exifTemplate
{
	"Copyright": "2013 William Artero",
	"ImageDescription": "some image",
	"ImageUniqueID": "https://...",
	"Make": "#fff"
}
endef
