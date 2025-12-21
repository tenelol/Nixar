package main

import (
	"log"
	"net/http"

	"github.com/tenelol/nixar/framework"
	"nixar-template/apps/web"
)

func main() {
	app := framework.NewApp()

	app.Use(framework.Logging())

	app.Get("/", func(ctx *framework.Context) {
		ctx.HTMLFile("apps/web/index.html")
	})

	app.Get("/about", func(ctx *framework.Context) {
		ctx.HTMLFile("apps/web/about.html")
	})

	app.Get("/api/hello", web.HelloAPI)

	app.Get("/health", func(ctx *framework.Context) {
		ctx.JSON(http.StatusOK, map[string]any{
			"status": "ok",
		})
	})

	staticHandler := http.StripPrefix("/static/", framework.Static("apps/web"))
	app.Get("/static/*file", framework.WrapHTTPHandler(staticHandler))

	addr := ":8080"
	log.Println("Listening on", addr)

	if err := http.ListenAndServe(addr, app); err != nil {
		log.Fatal(err)
	}
}
