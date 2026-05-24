{
  pkgs,
}:
rec {
  default = pkgs.mkShell {
    packages = [
      pkgs.d2
      pkgs.editorconfig-checker
      pkgs.git
      pkgs.remake
      pkgs.remarshal
      pkgs.unstable.cocogitto
      pkgs.unstable.go
      pkgs.unstable.golangci-lint
      pkgs.unstable.hugo
      pkgs.unstable.just
      pkgs.vale
      pkgs.wrangler
    ];
  };

  ci = default.overrideAttrs (
    final: prev: {
      nativeBuildInputs = pkgs.lib.subtractLists [
        pkgs.wrangler # no need for it during integration
      ] prev.nativeBuildInputs;
    }
  );

  terminal = default.overrideAttrs (
    final: prev: {
      nativeBuildInputs = [
        pkgs.unstable.gotools
      ]
      ++ prev.nativeBuildInputs;
    }
  );
}
