{
  description = "mywebfw - minimal Nix flake for Go server";

  inputs.nixpkgs.url = "github:nixos/nixpkgs/nixos-24.05";

  outputs = { self, nixpkgs }:
    let
      system = "x86_64-linux";
      pkgs = import nixpkgs { inherit system; };

      server = pkgs.buildGoModule {
        pname = "mywebfw-server";
        version = "0.1.0";

        # このリポジトリ全体を Go モジュールとして使う
        src = ./.;

        # ルートに go.mod がある前提
        modRoot = ".";

        # 実際の main パッケージは cmd/server
        subPackages = [ "./cmd/server" ];

        # go mod vendor してないので null でOK（←お前の言う通り）
        vendorHash = null;
      };
    in {
      # パッケージ
      packages.${system} = {
        server = server;
        default = server;
      };

      defaultPackage.${system} = server;

      # nix run .#server で起動できるように
      apps.${system}.default = {
        type = "app";
        program = "${server}/bin/server";
      };
    };
}

