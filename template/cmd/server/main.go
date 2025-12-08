package main

import (
	"log"
	"net/http"

	"github.com/tenelol/nixar/framework"
	"github.com/tenelol/nixar/apps/web"
)

func main() {
	app := framework.NewApp()

	app.Use(framework.Logging())

	app.Get("/", web.Index)

	addr := ":8080"
	log.Println("Listening on", addr)

	if err := http.ListenAndServe(addr, app); err != nil {
		log.Fatal(err)
	}
}

