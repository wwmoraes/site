.PHONY: vale-check
vale-check: .styles
	$(info linting prose...)
	@vale --minAlertLevel=error content

.PRECIOUS: .styles
.SECONDARY: .styles
.styles: .vale.ini
	$(info updating prose linting settings...)
	@vale sync
