## EXIF tags management
## https://exiftool.org/examples.html
## https://exiftool.org/TagNames/EXIF.html

IMAGES = $(shell find archetypes assets content -type f -name '*.jpg' -or -name '*.png')

.PHONY: exif
exif: ${IMAGES}

.PHONY: exif-show
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

%.jpg: %.jpg.json
	$(info updating EXIF: $@)
	@exiftool -q -overwrite_original_in_place -"*=" -j=$< $@
