.PHONY: check
check: .lefthook-local.yaml
	@lefthook run pre-commit
	@lefthook run pre-push

.PHONY: hooks
hooks:
	@${RM} .git/hooks/*
	@lefthook install

.PHONY: lefthook
lefthook: .lefthook-local.yaml

ifeq (${CI},)
.lefthook-local.yaml: .lefthook-system.yaml
else
.lefthook-local.yaml: .lefthook-ci.yaml
endif
	$(info linking $@ to $<)
	@ln -sf $< $@
