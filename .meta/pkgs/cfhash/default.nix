{
  lib,
  yarn2nix-moretea,
}:
let
  inherit (yarn2nix-moretea)
    mkYarnPackage
    importOfflineCache
    ;
in
mkYarnPackage rec {
  name = "cfhash";
  version = "0.0.0";

  src =
    with lib.fileset;
    toSource {
      root = ./.;
      fileset = unions [
        ./index.mjs
        ./package.json
        ./yarn.lock
        ./yarn.nix
      ];
    };

  packageJSON = "${src}/package.json";
  yarnLock = "${src}/yarn.lock";
  yarnNix = "${src}/yarn.nix";
  offlineCache = importOfflineCache "${src}/yarn.nix";

  doDist = false;
  dontStrip = true;

  meta.mainProgram = "cfhash";
}
