package main

import (
	"log"
	"net/http"

	"mywebfw/framework"
	"mywebfw/apps/portfolio"
)

func main() {
	r := framework.NewRouter()

	serveHTML := func(path string) framework.HandlerFunc {
		return func(ctx *framework.Context) {
			http.ServeFile(ctx.W, ctx.Req, path)
		}
	}

	// ページ
	r.Handle("GET", "/",        serveHTML("apps/portfolio/index.html"))
	r.Handle("GET", "/about",   serveHTML("apps/portfolio/about.html"))
	r.Handle("GET", "/works",   serveHTML("apps/portfolio/works.html"))
	r.Handle("GET", "/contact", serveHTML("apps/portfolio/contact.html"))

	// API: プロジェクト一覧
	r.Handle("GET", "/api/projects", portfolio.ProjectsAPI)

	// 既存のやつはそのまま
	r.Handle("GET", "/health", func(ctx *framework.Context) {
		framework.JSON(ctx, http.StatusOK, map[string]any{
			"status": "ok",
		})
	})

	r.Handle("GET", "/users/:id", func(ctx *framework.Context) {
		id := ctx.Params["id"]
		framework.JSON(ctx, http.StatusOK, map[string]any{
			"userId": id,
		})
	})

	staticHandler := http.StripPrefix("/static/", framework.Static("apps/portfolio"))
	r.Handle("GET", "/static/:filePath", func(ctx *framework.Context) {
		staticHandler.ServeHTTP(ctx.W, ctx.Req)
	})

	addr := ":8080"
	log.Println("Listening on", addr)

	handler := framework.LoggingMiddleware(r)
	if err := http.ListenAndServe(addr, handler); err != nil {
		log.Fatal(err)
	}
}

