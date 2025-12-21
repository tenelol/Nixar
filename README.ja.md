# nixar

Go標準の`net/http`をベースにした最小構成のWebフレームワークです。最小のGoフレームワークと`flake.nix`をセットで提供することをコンセプトにしています。

言語: [English](README.md) | 日本語

## 特徴
- `:param` 形式のシンプルなルーティング
- コンテキストベースのミドルウェア
- JSON/HTML のヘルパー
- 静的ファイル配信（no-cache）
- `http.Handler` / `http.HandlerFunc` のラッパー

## インストール（Nix）
```bash
mkdir myapp
cd myapp
nix flake init -t github:tenelol/nixar
```

起動:
```bash
nix run
```

## インストール（Goモジュール）
```bash
go get github.com/tenelol/nixar
```

## クイックスタート
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

## ルーティング
```go
app.Get("/users/:id", func(ctx *framework.Context) {
	id := ctx.Param("id")
	ctx.JSON(200, map[string]string{"id": id})
})
```

## 静的ファイル
```go
staticHandler := http.StripPrefix("/static/", framework.Static("public"))
app.Get("/static/*filePath", framework.WrapHTTPHandler(staticHandler))
```

## Nix flake
このリポジトリには、サンプルサーバのビルドと実行のための最小構成 `flake.nix` が含まれています。

サーバを起動:
```bash
nix run
```

パッケージをビルド:
```bash
nix build
```

開発シェルに入る:
```bash
nix develop
```

## リポジトリ構成
- `framework/` コアフレームワーク
- `cmd/server/` サンプルサーバのエントリ
- `apps/simple/` ミニマルサイト＋APIの例
- `template/` スターターテンプレート（Go + Nix flake）

## ライセンス
MIT. 詳細は `LICENSE` を参照してください。
