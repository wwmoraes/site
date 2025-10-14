PLANTUML_SOURCES = $(strip $(shell git ls-files 'content/**/*.puml'))
PLANTUML_TARGETS = $(patsubst %.puml,%.png,${PLANTUML_SOURCES})

.PRECIOUS: content/%.png
content/%.png: content/%.puml
	$(info generating $@...)
	@kroki convert -o $@ $<

