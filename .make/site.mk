.PHONY: data
data: ${SITE}
	$(info updating data files)
	@op run --env-file=.env.secrets -- ./$< data update

.PHONY: radar
radar: content/radar/radar.svg

## make magic, not war ;)

content/radar/radar.svg: ${SITE} content/radar/radar.svg.tmpl $(wildcard content/radar/*.md)
	$(info updating radar files)
	@./$< radar update
