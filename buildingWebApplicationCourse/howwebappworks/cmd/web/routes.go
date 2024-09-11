package main

import (
	"net/http"

	"github.com/Esbaevnurdos/buildingapp/pkg/config"
	"github.com/Esbaevnurdos/buildingapp/pkg/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func routes(app *config.AppConfig) http.Handler {
    mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

    mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
    mux.Get("/about", http.HandlerFunc(handlers.Repo.About))
    return mux
}
