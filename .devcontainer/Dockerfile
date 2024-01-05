# cspell:disable
FROM golang:1.21-alpine3.19

RUN apk add --no-cache \
  build-base \
  delve \
  exiftool \
  git \
  golangci-lint \
  gopls \
  hugo \
  ;

RUN : \
  && adduser -D -u 1000 code code \
  && adduser code code \
  ;

RUN : \
  && mkdir -m 0775 -p /go/pkg \
  && chown -R root:code /go/pkg \
  ;

USER code
RUN mkdir -p ~/.cache/go-build
RUN : \
  && go install github.com/1pkg/gopium@latest \
  && go install github.com/1pkg/goutline@latest \
  && go install github.com/yuzutech/kroki-cli/cmd/kroki@latest \
  ;

VOLUME ["/go/pkg/mod"]
VOLUME ["/home/code/.cache/go-build"]
