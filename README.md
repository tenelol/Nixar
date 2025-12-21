# nixar

Minimal Go web framework built on `net/http`. The concept is a smallest-possible Go framework paired with a minimal `flake.nix`.

Languages: English | [日本語](README.ja.md)

## Features
- Simple routing with `:param` path parameters
- Middleware pipeline (context-based)
- JSON and HTML helpers
- Static file handler (no-cache)
- Adapter to wrap `http.Handler` / `http.HandlerFunc`

## Install
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
staticHandler := http.StripPrefix("/static/", framework.Static("public"))
app.Get("/static/*filePath", framework.WrapHTTPHandler(staticHandler))
```

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
- `template/` starter template (Go + Nix flake)

## License
MIT. See `LICENSE`.
