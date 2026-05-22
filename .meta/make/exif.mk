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
