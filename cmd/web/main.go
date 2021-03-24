package main

import (
	"github.com/ahojo/hello-world-web/pkg/config"
	"github.com/ahojo/hello-world-web/pkg/handlers"
	"github.com/ahojo/hello-world-web/pkg/render"
	"github.com/alexedwards/scs/v2"
	"log"
	"net/http"
	"time"
)

const portNumber = ":8000"

var app config.AppConfig

var session *scs.SessionManager

func main() {
	//fmt.Println("Hello World")

	// change to true when in production
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create tempalte cache")
	}
	app.TemplateCache = tc
	render.NewTemplates(&app)

	app.UseCache = false
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	//http.HandleFunc("/about", handlers.Repo.About)
	//http.HandleFunc("/", handlers.Repo.Home)
	//_ = http.ListenAndServe(portNumber, nil)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
