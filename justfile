set quiet

_default:
  echo "Build recipes:"
  remake --tasks
  echo
  just --list --list-heading $'Command recipes:\n' --list-prefix ''

# [arg("ENVIRONMENT", help="staging or production")]
[doc("deploys this site to Cloudflare Pages")]
deploy $ENVIRONMENT="":
  #!/usr/bin/env bash

  source .env

  : "${GIT_BRANCH:=$(git branch --show-current)}"
  if [[ "${GIT_BRANCH}" = "master" ]]; then
    : "${ENVIRONMENT:=production}"
  else
    : "${ENVIRONMENT:=staging}"
  fi
  : "${ASSETS_DIR:=dist/${ENVIRONMENT}}"
  : "${GIT_COMMIT_HASH:=$(git rev-parse HEAD)}"
  : "${GIT_COMMIT_MESSAGE:=$(git show -s --format='%s')}"

  if git status --porcelain | grep -q .; then
    GIT_DIRTY='true'
  else
    GIT_DIRTY='false'
  fi

  echo "{{BOLD + GREEN}}Environment{{NORMAL}}: ${ENVIRONMENT}"
  echo "{{BOLD + GREEN}}Git branch{{NORMAL}}: ${GIT_BRANCH}"
  echo "{{BOLD + GREEN}}Git commit hash{{NORMAL}}: ${GIT_COMMIT_HASH}"
  echo "{{BOLD + GREEN}}Git commit message{{NORMAL}}: ${GIT_COMMIT_MESSAGE}"
  echo "{{BOLD + GREEN}}Is git in a dirty state? ${GIT_DIRTY}"
  echo "{{BOLD + GREEN}}Assets dir{{NORMAL}}: {{BLUE}}${ASSETS_DIR}{{NORMAL}}"

  read -t 3 -N 1 -p "Continue? [y/N] " CONTINUE; echo
  case "${CONTINUE,,}" in y);; *) exit 2;; esac

  echo "building ${ASSETS_DIR}..."
  remake --assume-old=static "${ASSETS_DIR}"

  echo "deploying ${ASSETS_DIR} to ${ENVIRONMENT} environment..."
  netlifier -site wwmoraes -dir "${ASSETS_DIR}"

[doc("Inspects images EXIF metadata.")]
exif-inspect:
  remake bin/site
  git ls-files {archetypes,assets,content}'/**/*.'{jpg,png}'.json' | fzf -m | ifne xargs site image inspect

[doc("Inspects all images' EXIF metadata.")]
exif-inspect-all:
  remake bin/site
  git ls-files {archetypes,assets,content}'/**/*.'{jpg,png}'.json' | ifne xargs site image inspect

[doc("Fixes known Golang code issues.")]
golang-fix *FLAGS:
  golangci-lint run --fix --enable-only gci,gofumpt {{FLAGS}}

[doc("Purges Cloudflare's cache for this project.")]
purge-cache:
  #!/usr/bin/env bash
  source .env
  if command -v op > /dev/null; then
    # shellcheck disable=SC2046
    eval $(op inject --in-file=.env.secrets)
  fi

  : "${CLOUDFLARE_ZONE_ID:?Cloudflare zone ID not set}"
  : "${CLOUDFLARE_API_TOKEN:?Cloudflare API token not set}"

  curl -X POST -sS "https://api.cloudflare.com/client/v4/zones/${CLOUDFLARE_ZONE_ID}/purge_cache" \
    -H 'Content-Type: application/json' \
    -H "Authorization: Bearer ${CLOUDFLARE_API_TOKEN}" \
    -w 'HTTP_STATUS:%{http_code}' \
    ;

[doc("Starts a local server.")]
start ENVIRONMENT="development" *FLAGS='-p 8888':
  hugo server -e {{ENVIRONMENT}} {{FLAGS}}
