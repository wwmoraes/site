{
  # pkgs ? import <nixpkgs> { }
  pkgs ? import (fetchTarball "https://github.com/NixOS/nixpkgs/archive/nixpkgs-23.11-darwin.tar.gz") {}
}: let
  commitlint = pkgs.buildGoModule rec {
    pname = "commitlint";
    version = "0.10.1";

    src = pkgs.fetchFromGitHub {
      owner = "conventionalcommit";
      repo = "commitlint";
      rev = "v${version}";
      hash = "sha256-OJCK6GEfs/pcorIcKjylBhdMt+lAzsBgBVUmdLfcJR0=";
    };

    # vendorHash = pkgs.lib.fakeHash;
    vendorHash = "sha256-4fV75e1Wqxsib0g31+scwM4DYuOOrHpRgavCOGurjT8=";

    meta = with pkgs.lib; {
      description = "commitlint checks if your commit messages meets the conventional commit format";
      homepage = "https://github.com/conventionalcommit/commitlint";
      license = licenses.mit;
      maintainers = with maintainers; [ wwmoraes ];
    };
  };
in with pkgs; mkShell {
  packages = [
    checkmake
    commitlint
    editorconfig-checker
    fish
    git
    go
    go-task
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
