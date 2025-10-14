.PRECIOUS: static/favicon.ico
.SECONDARY: static/favicon.ico
static/favicon.ico: assets/images/brand/favicon.svg
	$(info generating $@...)
	@magick -background none $< -gravity center -extent 48x48 -quality 92 $@
