package server

import (
	"context"

	"github.com/FACorreiaa/go-cobra-quiz/internal"

	"github.com/go-chi/chi/v5"
)

func Router(s *internal.ServiceStore) *chi.Mux {
	// could have just stayed with 1.22 router but used chi group routing
	ctx := context.Background()
	h := internal.NewHandler(ctx, s)

	r := chi.NewRouter()

	r.Route("/session", func(r chi.Router) {
		r.Post("/", h.StartSession)
		r.Post("/set-name/{user_id}", h.SetName)
		r.Post("/{user_id}/submit-quiz", h.SubmitQuiz)
		r.Get("/ranking", h.GetRanking)
	})

	r.Route("/quiz", func(r chi.Router) {
		r.Get("/scores", h.GetAllScores)
		r.Get("/list", h.GetQuestions)
	})

	return r
}
