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

        src = ./.;
        modRoot = ".";
        subPackages = [ "./cmd/server" ];
        vendorHash = null;
      };
    in {
      # ================= packages / apps =================
      packages.${system} = {
        server = server;
        default = server;
      };

      defaultPackage.${system} = server;

      apps.${system}.default = {
        type = "app";
        program = "${server}/bin/server";
      };

      devShells.${system}.default = pkgs.mkShell {
        buildInputs = [
          pkgs.go
          pkgs.gopls
          pkgs.golangci-lint
          pkgs.air
          pkgs.nodejs_22
        ];
        GO111MODULE = "on";
      };

      # ================= NixOS module =================
      nixosModules.mywebfw = { config, lib, pkgs, ... }:
        let
          cfg = config.services.mywebfw;
        in {
          options.services.mywebfw = {
            enable = lib.mkEnableOption "mywebfw Go web server";

            package = lib.mkOption {
              type = lib.types.package;
              default = server;
              description = "mywebfw server package to run";
            };

            user = lib.mkOption {
              type = lib.types.str;
              default = "mywebfw";
              description = "User account under which the mywebfw service runs.";
            };

            group = lib.mkOption {
              type = lib.types.str;
              default = "mywebfw";
              description = "Group under which the mywebfw service runs.";
            };

            port = lib.mkOption {
              type = lib.types.int;
              default = 8080;
              description = "Port the mywebfw server listens on.";
            };

            workingDir = lib.mkOption {
              type = lib.types.path;
              default = "/var/lib/mywebfw";
              description = "Working directory for mywebfw server.";
            };

            extraArgs = lib.mkOption {
              type = lib.types.listOf lib.types.str;
              default = [ ];
              description = "Extra arguments passed to the server binary.";
            };

            openFirewall = lib.mkOption {
              type = lib.types.bool;
              default = true;
              description = "Whether to open the firewall for the mywebfw port.";
            };
          };

          config = lib.mkIf cfg.enable {
            users.users.${cfg.user} = {
              isSystemUser = true;
              group = cfg.group;
              home = cfg.workingDir;
            };

            users.groups.${cfg.group} = {};

            systemd.services.mywebfw = {
              description = "mywebfw Go web server";
              after = [ "network.target" ];
              wantedBy = [ "multi-user.target" ];

              serviceConfig = {
                WorkingDirectory = cfg.workingDir;
                ExecStart = "${cfg.package}/bin/server --port ${toString cfg.port} ${lib.concatStringsSep " " cfg.extraArgs}";
                User = cfg.user;
                Group = cfg.group;
                Restart = "on-failure";
              };
            };

            # 必要ならファイアウォールも自動で開ける
            networking.firewall.allowedTCPPorts =
              lib.mkIf cfg.openFirewall [ cfg.port ];
          };
        };
    };
}
