package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// Routes() function will be attached to app *application. Remember the receiver function concept.
// we can access the appplication struct at main from here by using receiver variable (app).
func (app *application) Routes() http.Handler {
	// create a router mux
	mux := chi.NewRouter()

	// using a chi middleware.
	// this middleware will be executed for every request.
	mux.Use(middleware.Recoverer)

	// handling request
	mux.Get("/", HomeHandler)

	return mux
}
