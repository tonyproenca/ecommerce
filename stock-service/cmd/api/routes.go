package main

import (
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (app *Config) routes() http.Handler {
	mux := chi.NewRouter()

	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	mux.Use(middleware.Heartbeat("/health"))

	mux.Post("/stock-product", app.StoreNewStockProduct)
	mux.Get("/stock-product/{productCode}", app.RetrieveStockProduct)
	mux.Delete("/stock-product/{productCode}", app.DeleteStockProduct)
	mux.Put("/stock-product", app.UpdateStockProduct)

	return mux
}
