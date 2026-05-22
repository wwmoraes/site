{
  pkgs,
  ...
}:
{
  projectRootFile = "flake.nix";

  programs.jsonfmt = {
    enable = true;
    excludes = [
      "themes/pico/layouts/list.feed.json"
    ];
  };
  programs.keep-sorted.enable = true;
  programs.mdformat = {
    enable = true;
    excludes = [
      "archetypes/blip.md"
    ];
    package = pkgs.mdformat.withPlugins (
      ps: with ps; [
        mdformat-footnote
        mdformat-frontmatter
        mdformat-gfm
        mdformat-simple-breaks
      ]
    );
    settings = {
      number = true;
      # wrap = 120;
    };
  };
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
