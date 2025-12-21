{
  description = "Nixar template app";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs/nixos-24.05";
    nixar.url = "github:tenelol/Nixar";
  };

  outputs = { self, nixpkgs, nixar }:
    let
      system = "x86_64-linux";
      pkgs = import nixpkgs { inherit system; };
      nixarPkg = nixar.packages.${system}.server;
    in {
      packages.${system}.default = pkgs.buildGoModule {
        pname = "nixar-template";
        version = "0.1.0";

        src = ./.;
        modRoot = ".";
        subPackages = [ "./cmd/server" ];
        vendorHash = null;
        postPatch = ''
          go mod edit -replace github.com/tenelol/nixar=${nixar}
        '';
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
