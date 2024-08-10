let
  pkgs = import (fetchTarball {
    url = "https://github.com/NixOS/nixpkgs/archive/refs/tags/24.05.tar.gz";
    sha256 = "1lr1h35prqkd1mkmzriwlpvxcb34kmhc9dnr48gkm8hh089hifmx";
  }) {};
  unstable = import (fetchTarball {
    name = "nixos-unstable-a14c5d651cee9ed70f9cd9e83f323f1e531002db";
    # url = "https://github.com/NixOS/nixpkgs/archive/refs/heads/nixpkgs-unstable.tar.gz";
    url = "https://github.com/NixOS/nixpkgs/archive/a14c5d651cee9ed70f9cd9e83f323f1e531002db.tar.gz";
    sha256 = "1b2dwbqm5vdr7rmxbj5ngrxm7sj5r725rqy60vnlirbbwks6aahb";
  }) {};
  kaizen = import (fetchTarball {
    name = "kaizen-8075b45edf93d8f95a00958fd3a1cc606ba3405c";
    url = "https://github.com/wwmoraes/kaizen/archive/8075b45edf93d8f95a00958fd3a1cc606ba3405c.tar.gz";
    sha256 = "1w6fkd5kqm2l2aij31kd9ddk84phbcir7flrbha6mxw2y4j36z09";
  }) { inherit pkgs; };
  inherit (pkgs) lib mkShell;
in mkShell rec {
  packages = [
    kaizen.go-commitlint
    pkgs.editorconfig-checker
    pkgs.git
    pkgs.go-task
    pkgs.gofumpt
    pkgs.hugo
    pkgs.imagemagick
    pkgs.lefthook
    pkgs.markdownlint-cli
    pkgs.stylelint # TODO replace with native code tool
    pkgs.typos
    pkgs.vale
    pkgs.yarn # TODO remove after replacing stylelint
    unstable.go
    unstable.golangci-lint
  ] ++ lib.optionals (builtins.getEnv "CI" != "") [ # CI-only
  ] ++ lib.optionals (builtins.getEnv "CI" == "") [ # local-only
    kaizen.gopium
    kaizen.goutline
    pkgs.gopls
    pkgs.gotools
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
