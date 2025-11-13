{
  description = "artero.dev website";

  inputs = {
    flake-parts.url = "github:hercules-ci/flake-parts";
    nixpkgs.url = "github:NixOS/nixpkgs/25.05";
    nur = {
      inputs.nixpkgs.follows = "nixpkgs";
      inputs.flake-parts.follows = "flake-parts";
      url = "github:nix-community/NUR";
    };
    # sops-nix = {
    # 	inputs.nixpkgs.follows = "nixpkgs";
    # 	url = "github:Mic92/sops-nix";
    # };
    systems.url = "github:nix-systems/default";
    treefmt-nix = {
      inputs.nixpkgs.follows = "nixpkgs";
      url = "github:numtide/treefmt-nix";
    };
    unstable.url = "github:NixOS/nixpkgs?rev=e38c80c027d6bbdfa2a305fc08e732b9fac4928a";
  };

  nixConfig = {
    substituters = [
      "https://wwmoraes.cachix.org/"
      "https://nix-community.cachix.org/"
      "https://cache.nixos.org/"
    ];
    trusted-public-keys = [
      "wwmoraes.cachix.org-1:N38Kgu19R66Jr62aX5rS466waVzT5p/Paq1g6uFFVyM="
      "nix-community.cachix.org-1:mB9FSh9qf2dCimDSUo8Zy7bkq5CX+/rkCWyvRCYg3Fs="
      "cache.nixos.org-1:6NCHdD59X431o0gWypbMrAURkbJ16ZPMQFGspcDShjY="
    ];
  };

  outputs =
    inputs@{
      self,
      flake-parts,
      nixpkgs,
      nur,
      # sops-nix,
      systems,
      treefmt-nix,
      unstable,
      ...
    }:
    (flake-parts.lib.mkFlake { inherit inputs; } {
      imports = [
        treefmt-nix.flakeModule
      ];

      flake = {
        overlays = {
          unstable = final: prev: {
            unstable = import unstable { inherit (prev) system; };
          };
          nur = nur.overlays.default;
          local = final: prev: {
            stylelint = final.callPackage .meta/pkgs/stylelint { };
            inherit (self.packages.${prev.system}) update-stylelint;
          };
        };
      };

      perSystem =
        { pkgs, system, ... }:
        {
          _module.args.pkgs = import nixpkgs {
            inherit system;
            overlays = [
              self.overlays.unstable
              self.overlays.nur
              self.overlays.local
            ];
            config = { };
          };

          devShells.default = import ./shell.nix { inherit pkgs; };

          packages = {
            update-stylelint =
              let
                inherit (pkgs.lib) getExe getExe';
              in
              pkgs.writeShellScriptBin "update-stylelint" ''
                pushd .meta/pkgs/stylelint > /dev/null
                ${getExe pkgs.yarn} outdated
                ${getExe pkgs.yarn} install --mode update-lockfile
                ${getExe' pkgs.yarn2nix "yarn2nix"} > yarn.nix
                VERSION=$(${getExe pkgs.yarn} info --json stylelint | ${getExe pkgs.jq} -r '.data.version')
                jq --arg VERSION $VERSION '.version = $VERSION' package.json | ${getExe' pkgs.moreutils "sponge"} package.json
                popd > /dev/null
              '';
            inherit (pkgs) stylelint;
            inherit (pkgs) cfhash;
          };

          treefmt = ./treefmt.nix;
        };

      systems = import systems;
    });
}
