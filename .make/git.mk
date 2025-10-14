GIT_BRANCH = $(shell git branch --show-current)
GIT_COMMIT_HASH = $(shell git rev-parse HEAD)
GIT_COMMIT_MESSAGE = $(shell git show -s --format='%s')
GIT_DIRTY = $(if $(shell git status --porcelain),true,false)

.PHONY: git-check
git-check:
	$(info checking if commit messages follow conventions...)
	@cog check --from-latest-tag > /dev/null
