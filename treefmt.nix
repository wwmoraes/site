{
  projectRootFile = "flake.nix";

  programs.jsonfmt = {
    enable = true;
    excludes = [
      "themes/pico/layouts/list.feed.json"
    ];
  };
  programs.keep-sorted.enable = true;
  programs.nixf-diagnose = {
    enable = true;
    excludes = [
      "**/yarn.nix"
    ];
  };
  programs.nixfmt = {
    enable = true;
    excludes = [
      "**/yarn.nix"
    ];
  };
  programs.statix.enable = true;
  programs.typos = {
    enable = true;
    excludes = [
      "*.asc"
    ];
  };
  programs.yamlfmt = {
    enable = true;
    settings = {
      formatter = {
        type = "basic";
        indentless_arrays = true;
        scan_folded_as_literal = true;
      };
    };
  };
}
