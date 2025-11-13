{
  lib,
  makeBinaryWrapper,
  yarn2nix-moretea,
}:
let
  inherit (lib) importJSON;
  inherit (yarn2nix-moretea)
    mkYarnModules
    importOfflineCache
    ;
  modules = mkYarnModules rec {
    pname = "stylelint";
    inherit ((importJSON packageJSON)) version;

    packageJSON = ./package.json;
    yarnLock = ./yarn.lock;
    yarnNix = ./yarn.nix;
    offlineCache = importOfflineCache ./yarn.nix;

    pkgConfig.stylelint.nativeBuildInputs = [
      makeBinaryWrapper
    ];

    postBuild = ''
      mkdir -p $out/bin
      # ln -sf $out/deps/stylelint/node_modules/.bin/stylelint $out/bin/stylelint
      makeWrapper $out/node_modules/.bin/stylelint $out/bin/stylelint \
        --set NODE_PATH $out/node_modules
    '';
  };
in
modules.overrideAttrs {
  meta.mainProgram = "stylelint";
}
