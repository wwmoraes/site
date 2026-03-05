{
  pkgs,
}:
rec {
  default = pkgs.mkShell {
    packages = [
      pkgs.cocogitto
      pkgs.editorconfig-checker
      pkgs.git
      pkgs.remake
      pkgs.stylelint
      pkgs.unstable.cocogitto
      pkgs.unstable.go
      pkgs.unstable.golangci-lint
      pkgs.unstable.hugo
      pkgs.unstable.just
      pkgs.vale
      pkgs.wrangler
    ];
  };

  terminal = default.overrideAttrs (
    final: prev: {
      nativeBuildInputs = [
        pkgs.imagemagick
        pkgs.nur.repos.wwmoraes.kroki-cli
        pkgs.unstable.gotools
        pkgs.update-stylelint
      ]
      ++ prev.nativeBuildInputs;
    }
  );
}
