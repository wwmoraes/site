VALE_STYLES_DIR = .styles
VALE_STYLES = ${VALE_STYLES_DIR}/alex ${VALE_STYLES_DIR}/Joblint ${VALE_STYLES_DIR}/proselint ${VALE_STYLES_DIR}/Readability ${VALE_STYLES_DIR}/write-good

.PHONY: vale
vale: ${VALE_STYLES}
	@vale .

.PHONY: vale-sync
vale-sync: ${VALE_STYLES}

${VALE_STYLES}: .vale.ini
	@vale sync
