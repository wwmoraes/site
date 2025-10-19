{
  pkgs,
}:
let
  inherit (pkgs) lib mkShell;
in
mkShell {
  packages =
    [
      pkgs.cfhash
      pkgs.cocogitto
      pkgs.editorconfig-checker
      pkgs.git
      pkgs.remake
      pkgs.stylelint
      pkgs.unstable.cocogitto
      pkgs.unstable.dart-sass
      pkgs.unstable.go
      pkgs.unstable.golangci-lint
      pkgs.unstable.hugo
      pkgs.update-stylelint
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
    ];
}
