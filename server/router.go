package server

import (
	"github.com/FACorreiaa/go-cobra-quiz/api/handler"
	"github.com/go-chi/chi/v5"
)

func Router() *chi.Mux {
	// could have just stayed with 1.22 router but chi route grouping is cool

	r := chi.NewRouter()
	r.Get("/", handler.HelloHandler)
	return r
}
