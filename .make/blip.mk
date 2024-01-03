.PHONY: blip
blip: ${SITE}
	@./$< radar blip create
	@${MAKE} radar
