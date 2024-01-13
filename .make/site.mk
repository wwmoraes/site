.PHONY: data
data: ${SITE}
	$(info updating data files)
	@op run --env-file=.env.secrets -- ./$< data update

.PHONY: radar
radar: content/radar/radar.svg

.PHONY: favicon
favicon: static/favicon.ico

## make magic, not war ;)

content/radar/radar.svg: ${SITE} content/radar/radar.svg.tmpl $(wildcard content/radar/*.md)
	$(info updating radar files)
	@./$< radar update

static/favicon.ico: assets/images/brand/favicon.svg
	$(info generating $@)
	@convert \
		-background none \
		$< \
		-gravity center \
		-extent 48x48 \
		-quality 92 \
		$@
