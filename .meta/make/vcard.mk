define VCARD_TEMPLATE
BEGIN:VCARD
VERSION:3.0
ADR;TYPE=home:;;;Den Haag;;;Netherlands
EMAIL:op://Personal/Me/Contact/email
FN:op://Personal/Me/preferred name
IMPP:mastodon:@wwmoraes@hachyderm.io
IMPP:linkedin:wwmoraes
KEY;PGP:https://artero.dev/pgp.asc
LOGO:https://artero.dev/images/brand/logo.png
N:op://Personal/Me/last name;op://Personal/Me/first name;;;
PHOTO:https://artero.dev/images/avatar.jpg
TEL:op://Personal/Me/Home/cell
TITLE:op://Personal/Me/occupation
TZ:op://Personal/Me/Home/timezone
URL:https://artero.dev
END:VCARD
endef

export VCARD_TEMPLATE

.PHONY: vcard
vcard: tmp/me.vcf bin/site
	site vcard $<

.PRECIOUS: tmp/me.vcf
tmp/me.vcf:
	@mkdir -p $(dir $@)
	echo "$$VCARD_TEMPLATE" | op inject > $@
