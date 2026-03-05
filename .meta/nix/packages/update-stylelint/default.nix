{
  pkgs,
}:
pkgs.writeShellApplication {
  name = "update-stylelint";
  runtimeInputs = with pkgs; [
    jq
    yarn
    yarn2nix
    moreutils
  ];
  text = ./script.bash;
}
