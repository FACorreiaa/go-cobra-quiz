package server

import (
	"github.com/FACorreiaa/go-cobra-quiz/api"

	"github.com/go-chi/chi/v5"
)

func Router(s *api.Service) *chi.Mux {
	// could have just stayed with 1.22 router but used chi group routing
	h := api.NewHandler(s)

	r := chi.NewRouter()

	r.Route("/session", func(r chi.Router) {
		r.Post("/", h.StartSession)
		r.Post("/set-name/{user_id}", h.SetName)
		r.Post("/{user_id}/submit-quiz", h.SubmitQuiz)
		r.Get("/ranking", h.GetRanking)
	})

	r.Get("/quiz/list", h.GetAllScores)

	return r
}
