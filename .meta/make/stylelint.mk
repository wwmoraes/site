check::
	$(info linting CSS...)
	@stylelint --allow-empty-input --cache --formatter compact '**.css' '**.scss'
