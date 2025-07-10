{ system ? builtins.currentSystem
, sources ? import ./nix/sources.nix
}:
let
  pkgs = import sources.nixpkgs {
    inherit system;
    config.packageOverrides = pkgs: {
      nur = import sources.NUR { inherit pkgs; };
      unstable = import sources.unstable { inherit system pkgs; };
    };
  };
  inherit (pkgs) lib mkShell;
in mkShell rec {
  packages = [
    pkgs.editorconfig-checker
    pkgs.git
    pkgs.lefthook
    pkgs.markdownlint-cli
    pkgs.nodejs-slim # needed for stylelint 💀
    pkgs.typos
    pkgs.unstable.cocogitto
    pkgs.unstable.dart-sass
    pkgs.unstable.go
    pkgs.unstable.go-task
    pkgs.unstable.golangci-lint
    pkgs.unstable.hugo
    pkgs.vale
  ] ++ lib.optionals (builtins.getEnv "CI" != "") [ # CI-only
  ] ++ lib.optionals (builtins.getEnv "CI" == "") [ # local-only
    pkgs.dockerfile-language-server-nodejs
    pkgs.imagemagick
    pkgs.markdown-oxide # markdown LSP
    pkgs.niv
    pkgs.nur.repos.wwmoraes.gopium
    pkgs.nur.repos.wwmoraes.goutline
    pkgs.nur.repos.wwmoraes.kroki-cli
    pkgs.unstable.delve
    pkgs.unstable.gopls
    pkgs.unstable.gotools
    pkgs.yarn # needed for stylelint 💀
    ## TODO kroki
  ];

  offlineCache = pkgs.fetchYarnDeps {
    yarnLock = ./yarn.lock;
    # hash = lib.fakeHash;
    hash = "sha256-M8LKHLH5DlHo9q7XFju6MEJskiYLxXLqCt2vG53g0O0=";
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
