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
    pkgs.go-task
    pkgs.gofumpt
    pkgs.hugo
    pkgs.lefthook
    pkgs.markdownlint-cli
    pkgs.nodejs-slim # needed for stylelint 💀
    pkgs.typos
    pkgs.unstable.go
    pkgs.unstable.golangci-lint
    pkgs.vale
  ] ++ lib.optionals (builtins.getEnv "CI" != "") [ # CI-only
  ] ++ lib.optionals (builtins.getEnv "CI" == "") [ # local-only
    pkgs.dockerfile-language-server-nodejs
    pkgs.git
    pkgs.gopls
    pkgs.gotools
    pkgs.imagemagick
    pkgs.markdown-oxide
    pkgs.niv
    pkgs.nur.repos.wwmoraes.go-commitlint
    pkgs.nur.repos.wwmoraes.gopium
    pkgs.nur.repos.wwmoraes.goutline
    pkgs.nur.repos.wwmoraes.kroki-cli
    pkgs.unstable.delve
    pkgs.unstable.golangci-lint-langserver
    pkgs.unstable.gopls
    pkgs.unstable.gotools
    pkgs.yarn # needed for stylelint 💀
    ## TODO kroki
  ];

  offlineCache = pkgs.fetchYarnDeps {
    yarnLock = ./yarn.lock;
    hash = "sha256-RF+sFG4eQXDSNd3dWutWgIBpdJAo/jHWd+4IVLndvOU=";
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
