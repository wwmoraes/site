{
  pkgs,
}:
let
  inherit (pkgs) lib mkShell;
in
mkShell rec {
  packages =
    [
      pkgs.cocogitto
      pkgs.editorconfig-checker
      pkgs.git
      pkgs.nodejs-slim # needed for stylelint 💀
      pkgs.remake
      pkgs.unstable.cocogitto
      pkgs.unstable.dart-sass
      pkgs.unstable.go
      pkgs.unstable.golangci-lint
      pkgs.unstable.hugo
      pkgs.vale
    ]
    ++ lib.optionals (builtins.getEnv "CI" != "") [
      # CI-only
    ]
    ++ lib.optionals (builtins.getEnv "CI" == "") [
      # local-only
      pkgs.imagemagick
      pkgs.nur.repos.wwmoraes.kroki-cli
      pkgs.unstable.gotools
      pkgs.yarn # needed for stylelint 💀
      pkgs.yarn2nix # needed for stylelint 💀
    ];

  yarnOfflineCache = pkgs.yarn2nix-moretea.importOfflineCache ./yarn.nix;

  shellHook = ''
    ${lib.getExe pkgs.yarn} config --offline set yarn-offline-mirror ${yarnOfflineCache}
    ${lib.getExe' pkgs.fixup-yarn-lock "fixup-yarn-lock"} yarn.lock
    ${lib.getExe pkgs.yarn}  install --offline --frozen-lockfile --ignore-platform --ignore-scripts --no-progress --non-interactive
    patchShebangs node_modules/
  '';
}
