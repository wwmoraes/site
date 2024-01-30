.PHONY: spell
spell: .styles
	@vale .

.styles: .vale.ini
	@vale sync
