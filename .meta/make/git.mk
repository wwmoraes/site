check::
	$(info checking if commit messages follow conventions...)
	@cog check --from-latest-tag > /dev/null

.PHONY: hooks/*
hooks/pre-commit: check
hooks/pre-push: all test check/git
