.PHONY: hooks/*
hooks/pre-commit: check
hooks/pre-push: all test check/git
	$(info checking if commit messages follow conventions...)
	cog check > /dev/null
