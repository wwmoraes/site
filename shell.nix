{
  # pkgs ? import <nixpkgs> { }
  pkgs ? import (fetchTarball "https://github.com/NixOS/nixpkgs/archive/nixpkgs-23.11-darwin.tar.gz") {}
}: with pkgs; mkShell {
  packages = with pkgs; [
    checkmake
    editorconfig-checker
    fish
    git
    gnumake
    go
    gofumpt
    golangci-lint
    hugo
    imagemagick
    lefthook
    perl536Packages.ImageExifTool
    reviewdog
    stylelint
    typos
    vale
    yarn
  ];

  # works only with nix develop --impure
  # GOCACHE = "${builtins.getEnv "HOME"}/Library/Caches/go-build";
  # GOMODCACHE = "${builtins.getEnv "HOME"}/go/pkg/mod";
  # GOBIN = "${pkgs.lib.getBin pkgs.go}/bin";

  installPhase = ''
    source $stdenv/setup
    yarn global add stylelint-config-standard-scss
    go install github.com/yuzutech/kroki-cli/cmd/kroki@latest
  '';

  # shellHook = ''
  #   # git status
  # '';
}
