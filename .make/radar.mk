.PHONY: radar-blip
#: Creates a new technology radar entry (a "blip").
radar-blip: bin/site
	@site radar blip create
	@${MAKE} content/radar/radar.svg

## make magic, not war :)

.PRECIOUS: content/radar/radar.svg
.SECONDARY: content/radar/radar.svg
content/radar/radar.svg: content/radar/radar.svg.tmpl $(wildcard content/radar/*.md) bin/site
	$(info generating $@...)
	@site radar update > /dev/null
