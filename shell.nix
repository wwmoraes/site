{ system ? builtins.currentSystem
}:
let
  pkgs = import (fetchTarball {
    url = "https://github.com/NixOS/nixpkgs/archive/refs/tags/24.05.tar.gz";
    sha256 = "1lr1h35prqkd1mkmzriwlpvxcb34kmhc9dnr48gkm8hh089hifmx";
  }) {
    inherit system;
    config.packageOverrides = pkgs: {
      nur = import (builtins.fetchTarball "https://github.com/nix-community/NUR/archive/master.tar.gz") {
        inherit pkgs;
      };
      unstable = import (fetchTarball {
        name = "nixos-unstable-a14c5d651cee9ed70f9cd9e83f323f1e531002db";
        # url = "https://github.com/NixOS/nixpkgs/archive/refs/heads/nixpkgs-unstable.tar.gz";
        url = "https://github.com/NixOS/nixpkgs/archive/a14c5d651cee9ed70f9cd9e83f323f1e531002db.tar.gz";
        sha256 = "1b2dwbqm5vdr7rmxbj5ngrxm7sj5r725rqy60vnlirbbwks6aahb";
      }) { inherit system pkgs; };
    };
  };
  inherit (pkgs) lib mkShell;
in mkShell rec {
  packages = [
    pkgs.editorconfig-checker
    pkgs.go-task
    pkgs.gofumpt
    pkgs.hugo
    pkgs.lefthook
    pkgs.markdownlint-cli
    pkgs.stylelint # TODO replace with native code tool
    pkgs.typos
    pkgs.unstable.go
    pkgs.unstable.golangci-lint
    pkgs.vale
    pkgs.yarn # TODO remove after replacing stylelint
  ] ++ lib.optionals (builtins.getEnv "CI" != "") [ # CI-only
  ] ++ lib.optionals (builtins.getEnv "CI" == "") [ # local-only
    pkgs.dockerfile-language-server-nodejs
    pkgs.git
    pkgs.gopls
    pkgs.gotools
    pkgs.imagemagick
    pkgs.markdown-oxide
    pkgs.nur.repos.wwmoraes.go-commitlint
    pkgs.nur.repos.wwmoraes.gopium
    pkgs.nur.repos.wwmoraes.goutline
    pkgs.nur.repos.wwmoraes.kroki-cli
    pkgs.unstable.delve
    pkgs.unstable.golangci-lint-langserver
    pkgs.unstable.gopls
    pkgs.unstable.gotools
    ## TODO kroki
  ];

  offlineCache = pkgs.fetchYarnDeps {
    yarnLock = ./yarn.lock;
    hash = "sha256-VxsMxEMathPYRXjg82dEoeNNX8RsAmzcCBVAdMlnicQ=";
  };

  configurePhase = ''
    runHook preConfigure

    export HOME=$(mktemp -d)
    yarn config --offline set yarn-offline-mirror ${offlineCache}
    fixup-yarn-lock yarn.lock
    yarn install --offline --frozen-lockfile --ignore-platform --ignore-scripts --no-progress --non-interactive
    patchShebangs node_modules/

    runHook postConfigure
  '';
}
