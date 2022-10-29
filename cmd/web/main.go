package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/steacep/hello-world/v2/pkg/config"
	"github.com/steacep/hello-world/v2/pkg/handlers"
	"github.com/steacep/hello-world/v2/pkg/render"
)

const portNumber = ":8080"

// main is the main application function
func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplate(&app)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	fmt.Printf("Starting application on port %s\n", portNumber)

	log.Fatal(http.ListenAndServe(portNumber, nil))
}
