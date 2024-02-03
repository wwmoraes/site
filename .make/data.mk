## make magic, not war ;)

data: ${SITE}
	$(info updating data files)
	@op run --env-file=.env.secrets -- ./$< data update
	@touch $@
