set quiet

_default:
  echo "Build recipes:"
  remake --tasks
  echo
  just --list --list-heading $'Command recipes:\n' --list-prefix ''

# [arg("ENVIRONMENT", help="staging or production")]
[doc("deploys this site to Cloudflare Pages")]
deploy $ENVIRONMENT="staging":
  #!/usr/bin/env bash

  source .env

  : "${CLOUDFLARE_PROJECT_NAME:?Cloudflare project name not set}"

  : "${ASSETS_DIR:=dist/${ENVIRONMENT}}"
  : "${GIT_BRANCH:=$(git branch --show-current)}"
  : "${GIT_COMMIT_HASH:=$(git rev-parse HEAD)}"
  : "${GIT_COMMIT_MESSAGE:=$(git show -s --format='%s')}"

  if git status --porcelain | grep -q .; then
    GIT_DIRTY='true'
  else
    GIT_DIRTY='false'
  fi

  if [ "${ENVIRONMENT}" == "production" ]; then
    CLOUDFLARE_PAGES_BRANCH=master
  else
    CLOUDFLARE_PAGES_BRANCH=staging
  fi

  echo "{{BOLD + GREEN}}Environment{{NORMAL}}: ${ENVIRONMENT}"
  echo "{{BOLD + GREEN}}Git branch{{NORMAL}}: ${GIT_BRANCH}"
  echo "{{BOLD + GREEN}}Git commmit hash{{NORMAL}}: ${GIT_COMMIT_HASH}"
  echo "{{BOLD + GREEN}}Git commmit message{{NORMAL}}: ${GIT_COMMIT_MESSAGE}"
  echo "{{BOLD + GREEN}}Is git in a dirty state? ${GIT_DIRTY}"
  echo "{{BOLD + GREEN}}Assets dir{{NORMAL}}: {{BLUE}}${ASSETS_DIR}{{NORMAL}}"
  echo "{{BOLD + GREEN}}Cloudflare project name{{NORMAL}}: ${CLOUDFLARE_PROJECT_NAME}"
  echo "{{BOLD + GREEN}}Cloudflare Pages branch{{NORMAL}}: ${CLOUDFLARE_PAGES_BRANCH}"

  read -t 3 -N 1 -p "Continue? [y/N] " CONTINUE; echo
  case "${CONTINUE,,}" in y);; *) exit 2;; esac

  echo "building ${ASSETS_DIR}..."
  remake "${ASSETS_DIR}"

  echo "deploying ${ASSETS_DIR} to ${ENVIRONMENT} environment (Pages branch ${CLOUDFLARE_PAGES_BRANCH})..."
  wrangler pages deploy "${ASSETS_DIR}" \
    --no-bundle \
    --upload-source-maps \
    --project-name "${CLOUDFLARE_PROJECT_NAME}" \
    --branch "${CLOUDFLARE_PAGES_BRANCH}" \
    --commit-hash "${GIT_COMMIT_HASH}" \
    --commit-message "${GIT_COMMIT_MESSAGE}" \
    --commit-dirty "${GIT_DIRTY}" \
    ;

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
  # shellcheck disable=SC2046
  eval $(op inject --in-file=.env.secrets)

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
