define PGP_EMAILS
$(strip
william@artero.dev
contact@artero.dev
security@artero.dev
git@artero.dev
)
endef

GNUPGHOME ?= ${HOME}/.gnupg

# Lists the trust database, public and private keys files as sources. This
# is more than what's necessary as it may trigger due to trust/key changes on
# unrelated entries. Better be safe than sorry.
PGP_SOURCES = $(wildcard ${GNUPGHOME}/private-keys-v1.d/*.key) ${GNUPGHOME}/pubring.kbx ${GNUPGHOME}/trustdb.gpg
WKD_TARGETS = $(foreach EMAIL,${PGP_EMAILS},.well-known/openpgpkey/artero.dev/hu/$(value PGP_WKD_HASH_${EMAIL}))

# make file with the generated variables to avoid calling GPG on every make
# invocation.
PGP_GEN_FILE := $(realpath $(dir $(lastword ${MAKEFILE_LIST})))/.pgp.gen.mk

# getWkdHash <email> returns the shell command to retrieve the WKD hash for that email address.
getWkdHash = gpg --options /dev/null --list-keys --with-wkd-hash '$(1)' | grep -A1 '<$(1)>' | tail -n1 | xargs | cut -d@ -f1

# generates the make variables to simulate associative arrays in both directions
# between WKD hashes and email addresses.
define generateWkdVars
PGP_WKD_HASH_$(1) = $(2)
PGP_WKD_EMAIL_$(2) = $(1)
endef

.DEFAULT_GOAL := all

.PHONY: all
all: favicon.ico pgp.asc ${WKD_TARGETS}

.well-known/openpgpkey/artero.dev: ${WKD_TARGETS}
	touch $@

## make magic, not war :)

${WKD_TARGETS}: ${PGP_GEN_FILE}

${PGP_GEN_FILE}: ${PGP_SOURCES}
	$(info generating $@...)
	$(file >$@)
	$(foreach EMAIL,${PGP_EMAILS},$(file >>$@,$(call generateWkdVars,${EMAIL},$(shell $(call getWkdHash,${EMAIL})))))

.PRECIOUS: .well-known/openpgpkey/artero.dev/hu/%
.SECONDARY: .well-known/openpgpkey/artero.dev/hu/%
.well-known/openpgpkey/artero.dev/hu/%: EMAIL=$(value PGP_WKD_EMAIL_$*)
.well-known/openpgpkey/artero.dev/hu/%: ${PGP_SOURCES}
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

.PRECIOUS: favicon.ico
.SECONDARY: favicon.ico
favicon.ico: ../assets/images/brand/favicon.svg
	$(info generating $@...)
	@magick -background none $< -gravity center -extent 48x48 -quality 92 $@

.PRECIOUS: pgp.asc
.SECONDARY: pgp.asc
pgp.asc: EMAIL=$(firstword ${PGP_EMAILS})
pgp.asc: ${PGP_SOURCES}
	$(info generating $@...)
	@gpg \
	--options /dev/null \
	--yes \
	--armor \
	--output $@ \
	--export-options export-minimal \
	--export '${EMAIL}'
