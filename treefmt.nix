{
  mkFormatterModule,
  pkgs,
  ...
}: {
  imports = [
    (mkFormatterModule {
      name = "stylelint";
      package = "stylelint";
      args = [
        "--cache"
        "--cache-location"
        "$TMPDIR/"
      ];
      includes = pkgs.lib.mkDefault [
        "'**.css'"
        "'**.scss'"
      ];
    })
    (
      { config, lib, ... }:
      let
        cfg = config.programs.stylelint;
      in
      {
        options.programs.stylelint = {
          configFile = lib.mkOption {
            type = with lib.types; nullOr str;
            default = null;
            example = ".stylelint.yaml";
            description = "Custom config file";
          };
          formatter = lib.mkOption {
            type =
              with lib.types;
              nullOr (
                either str (oneOf [
                  "compact"
                  "github"
                  "json"
                  "string"
                  "tap"
                  "unix"
                  "verbose"
                ])
              );
            default = "compact";
            example = "compact";
            description = "An output formatter.";
          };
        };

        config = lib.mkIf cfg.enable {
          settings.formatter.stylelint = {
            options =
              [
                "--formatter"
                cfg.formatter
              ]
              ++ lib.optionals (cfg.configFile != null) [
                "--config"
                cfg.configFile
              ]
              ++ (if (builtins.length cfg.includes) > 0 then cfg.includes else ".");
          };
        };
      }
    )
  ];

  projectRootFile = "flake.nix";

  programs.jsonfmt = {
    enable = true;
    excludes = [
      "themes/pico/layouts/_default/list.feed.json"
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
  programs.stylelint = {
    enable = true;
    includes = [
      "'**.scss'"
    ];
  };
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
