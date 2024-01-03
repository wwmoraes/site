KROKI ?= kroki
PLANTUML_SOURCES = $(shell find content -type f -name "*.puml")
PLANTUML_TARGETS = $(patsubst %.puml,%.png,${PLANTUML_SOURCES})

.PHONY: diagrams
diagrams: ${PLANTUML_TARGETS}

.PRECIOUS: %.png
%.png: %.puml
	@${KROKI} convert --format png $<
# plantuml -tpng -darkmode -theme reddress-darkblue $<
