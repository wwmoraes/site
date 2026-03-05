pushd .meta/nix/packages/stylelint > /dev/null || exit 1

yarn outdated
yarn install --mode update-lockfile
yarn2nix > yarn.nix

VERSION=$(yarn info --json stylelint | jq -r '.data.version')
jq --arg VERSION "${VERSION}" '.version = $VERSION' package.json | sponge package.json

popd > /dev/null || exit 1
