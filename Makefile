GO ?= go

DIAGRAMS_SOURCES = $(shell find content -type f -name "*.puml")
DIAGRAMS_TARGETS = $(patsubst %.puml,%.png,${DIAGRAMS_SOURCES})

IMAGES = $(wildcard archetypes/*/*.jpg)
IMAGES += $(wildcard assets/images/*.jpg)
IMAGES += $(wildcard content/posts/*/*.jpg)

start:
	@hugo server -p 8888

start-prod:
	@hugo server -e production

publish: bin/site
	@${RM} -r public
	@./$< radar update
	@hugo --gc --cleanDestinationDir

diagrams: ${DIAGRAMS_TARGETS}

.PHONY: build
build: bin/site

bin/site: $(shell ${GO} list -f '{{ range .GoFiles }}{{ printf "%s/%s\n" $$.Dir . }}{{ end }}' ./cmd/$(notdir $@)/...)
	$(info building $@)
	@go build -race -o ./$(dir $@) ./cmd/$(notdir $@)/...

.PRECIOUS: %.png
%.png: %.puml
	plantuml -tpng -darkmode -theme reddress-darkblue $<

.PHONY: data
data: bin/site
	@op run --env-file=.env -- ./$< data update

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
		-Artist \
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

.PHONY: radar
radar: content/radar/radar.svg

content/radar/radar.svg: bin/site content/radar/radar.svg.tmpl $(wildcard content/radar/*.md)
	@./$< radar update

blip: QUADRANT=$(shell echo "languages\nplatforms\ntechniques\ntools" | fzf -1)
blip: TIER=$(shell echo "adopt\ntrial\nassess\nhold" | fzf -1)
blip: NAME=$(shell read -p "Name: " && echo $$REPLY | tr '[A-Z]' '[a-z]' | tr ' ' '-')
blip: bin/site
	@./$< radar blip create -q ${QUADRANT} -t ${TIER} "${NAME}"
	@${MAKE} radar
