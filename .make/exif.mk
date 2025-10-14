EXIF_SOURCES = $(strip $(shell git ls-files {archetypes,assets,content}'/**/*.'{jpg,png}'.json'))
EXIF_TARGETS = $(patsubst %.json,%,${EXIF_SOURCES})

.PHONY: exif-inspect-all
#: Inspects all images' EXIF metadata.
exif-inspect-all: bin/site
	@site image inspect ${EXIF_SOURCES}

.PHONY: exif-inspect
#: Inspects specific image EXIF metadata.
exif-inspect: bin/site
	@echo ${EXIF_SOURCES} | tr ' ' '\n' | fzf -m | ifne xargs site image inspect

## make magic, not war :)

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
