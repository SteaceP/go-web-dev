package main

import (
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/steacep/hello-world/v2/pkg/config"
	"github.com/steacep/hello-world/v2/pkg/handlers"
)

func routes(app *config.AppConfig) http.Handler {
	mux := pat.New()

	mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	mux.Get("/about", http.HandlerFunc(handlers.Repo.About))

	return mux
}
