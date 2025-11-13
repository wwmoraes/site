css-check:
	$(info linting CSS...)
	@stylelint --cache --formatter compact '**.scss'
