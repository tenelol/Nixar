{ packages ? import <nixpkgs> {} }:

let
  baseShell = import ../../shells/shell.nix { inherit packages; };
in
packages.mkShell {
  inherit (baseShell) pure;

  buildInputs = baseShell.buildInputs ++ (with packages; [
    nodejs_22
    nodePackages.pnpm
  ]);

  shellHook = ''
    ${baseShell.shellHook}
    echo "Node.js frontend environment"
    echo "Node: $(node --version)"
    echo "pnpm: $(pnpm --version)"
  '';
}

