package main

import (
	"log"
	"net/http"

	"github.com/tenelol/nixar/apps/simple"
	"github.com/tenelol/nixar/framework"
)

func main() {
	app := framework.NewApp()

	// ミドルウェア
	app.Use(framework.Logging())

	// ページ
	app.Get("/", func(ctx *framework.Context) {
		ctx.HTMLFile("apps/simple/index.html")
	})

	app.Get("/about", func(ctx *framework.Context) {
		ctx.HTMLFile("apps/simple/about.html")
	})

	// API: サンプル
	app.Get("/api/hello", simple.HelloAPI)

	// ヘルスチェック
	app.Get("/health", func(ctx *framework.Context) {
		ctx.JSON(http.StatusOK, map[string]any{
			"status": "ok",
		})
	})

	// 静的ファイル
	staticHandler := http.StripPrefix("/static/", framework.Static("apps/simple"))
	app.Get("/static/*file", framework.WrapHTTPHandler(staticHandler))

	addr := ":8080"
	log.Println("Listening on", addr)

	if err := http.ListenAndServe(addr, app); err != nil {
		log.Fatal(err)
	}
}
