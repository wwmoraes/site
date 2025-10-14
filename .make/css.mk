css-check:
	$(info linting CSS...)
	@stylelint -f compact assets/css/**/*.scss themes/*/assets/css/**/*.scss

yarn.nix: yarn.lock
	$(info updating yarn nix...)
	@yarn2nix > $@

yarn.lock: package.json
	$(info updating yarn lock...)
	@yarn install --mode update-lockfile
