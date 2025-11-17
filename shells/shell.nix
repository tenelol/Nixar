{ packages ? import <nixpkgs> {} }:

packages.mkShell {
  buildInputs = with packages; [
    git
    go
  ];

  pure = true;

  shellHook = ''
    echo "mywebfw base dev environment"
  '';
}

