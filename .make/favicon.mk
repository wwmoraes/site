.PHONY: favicon
favicon: static/favicon.ico

## make magic, not war ;)

static/favicon.ico: assets/images/brand/favicon.svg
	$(info generating $@)
	@convert \
		-background none \
		$< \
		-gravity center \
		-extent 48x48 \
		-quality 92 \
		$@
