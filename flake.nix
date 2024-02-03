{
  inputs.flake-utils.url = "github:numtide/flake-utils";

  outputs = { self, nixpkgs, flake-utils }: flake-utils.lib.eachDefaultSystem (system:
    let pkgs = nixpkgs.legacyPackages.${system}; in
    {
      devShells.default = import ./shell.nix { inherit pkgs; };

      # packages = rec {
      #   site = pkgs.buildGoModule {
      #     pname = "site";
      #     version = self.rev or "dirty";

      #     # src = builtins.fetchGit {
      #     #   url = ./.;
      #     # };
      #     src = pkgs.lib.sources.sourceFilesBySuffices self [
      #       ".go"
      #       "go.mod"
      #       "go.sum"
      #     ];

      #     subPackages = [ "cmd/site" ];

      #     # vendorHash = pkgs.lib.fakeHash;
      #     vendorHash = "sha256-hSVoBd3/uUbAvqrXddKBfG94PmGLP6qZEf4lNjIVAUE=";

      #     doCheck = false;

      #     CGO_ENABLED = 0;

      #     meta = with pkgs.lib; {
      #       description = "Batteries toolkit for Hugo static sites";
      #       homepage = "https://github.com/wwmoraes/site";
      #       license = licenses.mit;
      #       maintainers = with maintainers; [ wwmoraes ];
      #     };
      #   };
      #   public = pkgs.stdenv.mkDerivation {
      #     pname = "public";
      #     version = self.rev or "dirty";
      #     src = pkgs.lib.sources.cleanSourcewith {
      #       src = self;
      #       filter = path: type: builtins.any (path) [".git"];
      #     };
      #     # src = pkgs.fetchgit {
      #     #   url = self;
      #     #   leaveDotGit = true;
      #     # };

      #     buildPhase = ''
      #       source $stdenv/setup
      #       make build
      #     '';

      #     installPhase = ''
      #       mkdir -p $out
      #       cp -r public $out
      #     '';
      #   };
      #   default = site;
      # };
    }
  );
}
