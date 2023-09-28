package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/hannasoderstromdev/go-web-2/pkg/config"
	"github.com/hannasoderstromdev/go-web-2/pkg/handlers"
)

func routes(app *config.AppConfig) http.Handler {
	// mux := pat.New()

	// mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	// mux.Get("/about", http.HandlerFunc(handlers.Repo.About))

	mux := chi.NewRouter()

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	return mux
}