.PHONY: check
check: .lefthook-local.yaml
	@lefthook run pre-commit --all-files --force

.PHONY: hooks
hooks: .git/hooks/pre-commit

.git/hooks/%: $(wildcard .lefthook*.yaml)
	$(info installing $* hook)
	@lefthook add $*
	@touch $@

.lefthook-local.yaml: .lefthook-system.yaml
	@ln -s $< $@
