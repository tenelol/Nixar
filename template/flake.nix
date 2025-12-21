{
  description = "Nixar template app";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs/nixos-24.05";
  };

  outputs = { self, nixpkgs }:
    let
      system = "x86_64-linux";
      pkgs = import nixpkgs { inherit system; };
    in {
      packages.${system}.default = pkgs.buildGoModule {
        pname = "nixar-template";
        version = "0.1.0";

        src = ./.;
        modRoot = ".";
        subPackages = [ "./cmd/server" ];
        vendorHash = "sha256-MMDjNtjZdpC0JZf+6jg2a5MPmf3gSFlvB1nyY0GowEw=";
      };

      apps.${system}.default = {
        type = "app";
        program = "${self.packages.${system}.default}/bin/server";
      };

      devShells.${system}.default = pkgs.mkShell {
        buildInputs = [
          pkgs.go
          pkgs.gopls
        ];
      };
    };
}
