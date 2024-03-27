package server

import (
	"github.com/FACorreiaa/go-cobra-quiz/api"

	"github.com/go-chi/chi/v5"
)

func Router(s *api.Service) *chi.Mux {
	// could have just stayed with 1.22 router but used chi group routing
	h := api.NewHandler(s)

	r := chi.NewRouter()

	r.Post("/session", h.StartSession)
	r.Post("/set-name/{user_id}", h.SetName)
	//

	return r
}
