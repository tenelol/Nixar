package main

import (
	"log"
	"net/http"

	"nexar/framework"
	"nexar/apps/portfolio"
)

func main() {
	app := framework.NewApp()

	// ミドルウェア
	app.Use(framework.Logging())

	// ページ
	app.Get("/", func(ctx *framework.Context) {
		ctx.HTMLFile("apps/portfolio/index.html")
	})

	app.Get("/about", func(ctx *framework.Context) {
		ctx.HTMLFile("apps/portfolio/about.html")
	})

	app.Get("/works", func(ctx *framework.Context) {
		ctx.HTMLFile("apps/portfolio/works.html")
	})

	app.Get("/contact", func(ctx *framework.Context) {
		ctx.HTMLFile("apps/portfolio/contact.html")
	})

	// API: プロジェクト一覧
	app.Get("/api/projects", portfolio.ProjectsAPI)

	// ヘルスチェック
	app.Get("/health", func(ctx *framework.Context) {
		ctx.JSON(http.StatusOK, map[string]any{
			"status": "ok",
		})
	})

	// 静的ファイル
	staticHandler := http.StripPrefix("/static/", framework.Static("apps/portfolio"))
	app.Get("/static/:filePath", framework.WrapHTTPHandler(staticHandler))

	addr := ":8080"
	log.Println("Listening on", addr)

	if err := http.ListenAndServe(addr, app); err != nil {
		log.Fatal(err)
	}
}

