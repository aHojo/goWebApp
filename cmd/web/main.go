package main

import (
	"github.com/ahojo/hello-world-web/pkg/config"
	"github.com/ahojo/hello-world-web/pkg/handlers"
	"github.com/ahojo/hello-world-web/pkg/render"
	"log"
	"net/http"
)

const portNumber = ":8000"

func main() {
	//fmt.Println("Hello World")
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create tempalte cache")
	}
	app.TemplateCache = tc
	render.NewTemplates(&app)

	app.UseCache = false
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	http.HandleFunc("/about", handlers.Repo.About)

	http.HandleFunc("/", handlers.Repo.Home)

	_ = http.ListenAndServe(portNumber, nil)
}
