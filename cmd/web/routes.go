package main

import (
	"bookings/internal/config"
	"bookings/internal/handlers"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func routes(app *config.AppConfig) http.Handler {
	//mux := pat.New()
	//mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	//mux.Get("/about", http.HandlerFunc(handlers.Repo.About))

	mux := chi.NewRouter()
	mux.Use(WriteToConsole)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)
	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	mux.Get("/make-reservation", handlers.Repo.Reservation)
	mux.Get("/majors-suite", handlers.Repo.Majors)
	mux.Get("/generals-quarters", handlers.Repo.Generals)
	mux.Get("/search-availability", handlers.Repo.Availability)
	mux.Post("/search-availability", handlers.Repo.PostAvailability)
	mux.Post("/search-availability-json", handlers.Repo.AvailabilityJson)
	mux.Get("/contact", handlers.Repo.Contact)

	//Create a file server for static files
	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))
	return mux
}
