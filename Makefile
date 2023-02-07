start:
	@exec hugo server -p 8888

start-prod:
	@hugo server -e production

build:
	@rm -rf public
	@hugo
