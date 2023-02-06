start:
	@exec hugo server -p 8888

start-prod:
	@hugo server -e production

build:
	@rm -rf public
	@hugo

hook-install:
	@pre-commit install

hook-update:
	@pre-commit autoupdate

check:
	@pre-commit run --all-files
