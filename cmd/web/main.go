package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/hannasoderstromdev/go-web-2/pkg/config"
	"github.com/hannasoderstromdev/go-web-2/pkg/handlers"
	"github.com/hannasoderstromdev/go-web-2/pkg/render"
)

const portNumber = ":8080"

// main is the main application function
func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app) // Pass app as reference, returns repo
	handlers.NewHandlers(repo) // Give repo to NewHandlers

	render.NewTemplates(&app)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))

	srv := &http.Server {
		Addr: portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}