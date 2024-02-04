.PHONY: check
check: .lefthook-local.yaml
	@lefthook run pre-commit
	@lefthook run pre-push

.PHONY: hooks
hooks:
	@${RM} .git/hooks/*
	@lefthook install

.lefthook-local.yaml: .lefthook-system.yaml
	@ln -s $< $@
