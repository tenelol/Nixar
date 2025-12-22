# nixar

Minimal Go web framework built on `net/http`. The concept is a smallest-possible Go framework paired with a minimal `flake.nix`.

Languages: English | [日本語](README.ja.md)

## Features
- Simple routing with `:param` path parameters
- Middleware pipeline (context-based)
- JSON and HTML helpers
- Static file handler (no-cache)
- Adapter to wrap `http.Handler` / `http.HandlerFunc`

## Install (Nix)
```bash
mkdir myapp
cd myapp
nix flake init -t github:tenelol/nixar
```

Run:
```bash
nix run
```

## Install (Go module)
```bash
go get github.com/tenelol/nixar
```

## Quick Start
```go
package main

import (
	"log"
	"net/http"

	"github.com/tenelol/nixar/framework"
)

func main() {
	app := framework.NewApp()
	app.Use(framework.Logging())

	app.Get("/", func(ctx *framework.Context) {
		ctx.JSON(http.StatusOK, map[string]any{"hello": "nixar"})
	})

	addr := ":8080"
	log.Println("Listening on", addr)
	log.Fatal(http.ListenAndServe(addr, app))
}
```

## Routing
```go
app.Get("/users/:id", func(ctx *framework.Context) {
	id := ctx.Param("id")
	ctx.JSON(200, map[string]string{"id": id})
})
```

## Static Files
```go
staticHandler := http.StripPrefix("/static/", framework.Static("static"))
app.Get("/static/*filePath", framework.WrapHTTPHandler(staticHandler))
```
`static` is just an example; you can serve any directory.

## Server flags
- `--port` (default: `8080`)
- `--pages-dir` (default: `apps/simple`)
- `--static-dir` (default: `static`)

## Nix flake
This repository includes a minimal `flake.nix` for building and running the example server.

Run the server:
```bash
nix run
```

Build the package:
```bash
nix build
```

Enter the dev shell:
```bash
nix develop
```

## Repository Layout
- `framework/` core framework
- `cmd/server/` example server entry
- `apps/simple/` example app (minimal site + API)
- `static/` static assets for the sample server
- `template/` starter template (Go + Nix flake)

## NixOS module
Enable the module in your NixOS configuration and it will run the server as a systemd service.
It also bundles the sample HTML and static assets and serves them automatically.

Example (flake-based):
```nix
{
  inputs.nixar.url = "github:tenelol/nixar";

  outputs = { self, nixar, nixpkgs, ... }: {
    nixosConfigurations.my-host = nixpkgs.lib.nixosSystem {
      system = "x86_64-linux";
      modules = [
        nixar.nixosModules.nixar
        {
          services.nixar.enable = true;
          services.nixar.port = 8080;
        }
      ];
    };
  };
}
```

## License
MIT. See `LICENSE`.
