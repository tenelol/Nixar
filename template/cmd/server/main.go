package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"path/filepath"

	"github.com/tenelol/nixar/framework"
	"nixar-template/apps/web"
)

func main() {
	port := flag.Int("port", 8080, "server listen port")
	pagesDir := flag.String("pages-dir", "apps/web", "directory for HTML pages")
	staticDir := flag.String("static-dir", "static", "directory for static assets")
	flag.Parse()

	app := framework.NewApp()

	app.Use(framework.Logging())

	app.Get("/", func(ctx *framework.Context) {
		ctx.HTMLFile(filepath.Join(*pagesDir, "index.html"))
	})

	app.Get("/about", func(ctx *framework.Context) {
		ctx.HTMLFile(filepath.Join(*pagesDir, "about.html"))
	})

	app.Get("/api/hello", web.HelloAPI)

	app.Get("/health", func(ctx *framework.Context) {
		ctx.JSON(http.StatusOK, map[string]any{
			"status": "ok",
		})
	})

	staticHandler := http.StripPrefix("/static/", framework.Static(*staticDir))
	app.Get("/static/*file", framework.WrapHTTPHandler(staticHandler))

	addr := fmt.Sprintf(":%d", *port)
	log.Println("Listening on", addr)

	if err := http.ListenAndServe(addr, app); err != nil {
		log.Fatal(err)
	}
}
