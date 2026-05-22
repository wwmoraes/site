{
  pkgs,
}:
rec {
  default = pkgs.mkShell {
    packages = [
      pkgs.cocogitto
      pkgs.d2
      pkgs.editorconfig-checker
      pkgs.git
      pkgs.remake
      pkgs.unstable.cocogitto
      pkgs.unstable.go
      pkgs.unstable.golangci-lint
      pkgs.unstable.hugo
      pkgs.unstable.just
      pkgs.vale
      pkgs.wrangler
      pkgs.remarshal
    ];
  };

  terminal = default.overrideAttrs (
    final: prev: {
      nativeBuildInputs = [
        pkgs.unstable.gotools
      ]
      ++ prev.nativeBuildInputs;
    }
  );
}
