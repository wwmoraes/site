.PHONY: blip
blip: QUADRANT=$(shell echo "languages\nplatforms\ntechniques\ntools" | fzf -1)
blip: TIER=$(shell echo "adopt\ntrial\nassess\nhold" | fzf -1)
blip: NAME=$(shell read -p "Name: " && echo $$REPLY | tr '[A-Z]' '[a-z]' | tr ' ' '-')
blip: ${SITE}
	@./$< radar blip create -q ${QUADRANT} -t ${TIER} "${NAME}"
	@${MAKE} radar
