GO ?= go

bin/%: $(shell ${GO} list -f '{{ range .GoFiles }}{{ printf "%s/%s\n" $$.Dir . }}{{ end }}' ./...)
	$(info building $@)
	@${GO} build -race -o ./$(dir $@) ./cmd/$*/...
