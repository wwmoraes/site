GNUPGHOME ?= ${HOME}/.gnupg

PGP_SOURCES = $(wildcard ${GNUPGHOME}/private-keys-v1.d/*.key) ${GNUPGHOME}/pubring.kbx ${GNUPGHOME}/trustdb.gpg
PGP_TARGETS = static/pgp.asc $(foreach EMAIL,${PGP_EMAILS},static/.well-known/openpgpkey/artero.dev/hu/$(value PGP_WKD_HASH_${EMAIL}))

PGP_GEN_FILE := $(dir $(lastword ${MAKEFILE_LIST}))pgp.gen.mk

getWkdHash = gpg --options /dev/null --list-keys --with-wkd-hash '$(1)' | grep -A1 '<$(1)>' | tail -n1 | xargs | cut -d@ -f1

define generateWkdVars
PGP_WKD_HASH_$(1) = $(2)
PGP_WKD_EMAIL_$(2) = $(1)
endef

## make magic, not war :)

${PGP_TARGETS}: ${PGP_GEN_FILE}

${PGP_GEN_FILE}: ${PGP_SOURCES}
	$(info generating $@...)
	$(file >$@)
	$(foreach EMAIL,${PGP_EMAILS},$(file >>$@,$(call generateWkdVars,${EMAIL},$(shell $(call getWkdHash,${EMAIL})))))

.PRECIOUS: static/.well-known/openpgpkey/artero.dev/hu/%
.SECONDARY: static/.well-known/openpgpkey/artero.dev/hu/%
static/.well-known/openpgpkey/artero.dev/hu/%: EMAIL=$(value PGP_WKD_EMAIL_$*)
static/.well-known/openpgpkey/artero.dev/hu/%: ${PGP_SOURCES}
	$(info generating $@...)
	@gpg \
	--options /dev/null \
	--yes \
	--no-armor \
	--output '$@' \
	--export-filter drop-subkey='usage =~ a || usage =~ c' \
	--export-filter keep-uid='uid =~ <${EMAIL}>' \
	--export-options export-minimal \
	--export '${EMAIL}'

.PRECIOUS: static/pgp.asc
.SECONDARY: static/pgp.asc
static/pgp.asc: EMAIL=$(firstword ${PGP_EMAILS})
static/pgp.asc: ${PGP_SOURCES}
	$(info generating $@...)
	@gpg \
	--options /dev/null \
	--yes \
	--armor \
	--output $@ \
	--export-options export-minimal \
	--export '${EMAIL}'
