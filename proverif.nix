let pkgs = import <nixpkgs> {};

in pkgs.mkShell rec {
  name = "devenv";

  buildInputs = with pkgs; [
    proverif graphviz 
  ];
}