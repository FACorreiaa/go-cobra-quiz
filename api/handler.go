package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"sort"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type Handler struct {
	ctx     context.Context
	service *ServiceStore
}

func NewHandler(ctx context.Context, s *ServiceStore) *Handler {
	return &Handler{ctx: ctx, service: s}
}

func (h *Handler) StartSession(w http.ResponseWriter, r *http.Request) {
	var user User
	var session Session

	user, err := h.service.User.createUserID(h.ctx, user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = h.service.User.getUser(h.ctx, user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	session, err = h.service.Session.createSessionID(h.ctx, session)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//ctx := context.WithValue(r.Context(), "sessionID", session.ID)

	response := CreateUser{
		UserID:    user.ID,
		SessionID: session.ID,
	}

	//h.SetName(w, r.WithContext(ctx))

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(response)
}

func (h *Handler) SetName(w http.ResponseWriter, r *http.Request) {
	userIDParam := chi.URLParam(r, "user_id")
	userID, err := uuid.Parse(userIDParam)

	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	var newName struct {
		Name string `json:"name"`
	}

	if err := json.NewDecoder(r.Body).Decode(&newName); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if newName.Name == "" {
		http.Error(w, "Empty name provided", http.StatusBadRequest)
		return
	}

	if err := h.service.User.updateUserName(h.ctx, userID, newName.Name); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := UpdateUser{
		UserID:   userID,
		Username: newName.Name,
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(response)

}

func (h *Handler) SubmitQuiz(w http.ResponseWriter, r *http.Request) {
	var userAnswers map[string]string
	if err := json.NewDecoder(r.Body).Decode(&userAnswers); err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	userIDStr := chi.URLParam(r, "user_id")
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}
	user, err := h.service.User.getUserByID(h.ctx, userID)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	score, correctAnswers, err := h.service.Quiz.createUserAnswers(h.ctx, userAnswers, user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.service.User.updateUserScore(h.ctx, user, score); err != nil {
		http.Error(w, "Failed to save user information", http.StatusInternalServerError)
		return
	}

	usersWithAnswers, err := h.service.Ranking.getUsersResults(h.ctx)
	if err != nil {
		http.Error(w, "Failed to get users with answers", http.StatusInternalServerError)
		return
	}

	percentile := h.service.User.calculateUserPercent(h.ctx, usersWithAnswers, score)
	percentileMessage := fmt.Sprintf("You are better than %.2f%% of users who already submitted their quiz", percentile)

	response := QuestionsReply{
		Score:          score,
		CorrectAnswers: correctAnswers,
		Percentile:     percentile,
		Message:        percentileMessage,
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(response)
}

func (h *Handler) GetAllScores(w http.ResponseWriter, r *http.Request) {
	usersWithAnswers, err := h.service.Ranking.getUsersResults(h.ctx)
	if err != nil {
		http.Error(w, "Failed to get users with answers", http.StatusInternalServerError)
		return
	}

	sort.Slice(usersWithAnswers, func(i, j int) bool {
		return usersWithAnswers[i].Score > usersWithAnswers[j].Score
	})

	jsonResponse, err := json.Marshal(usersWithAnswers)
	if err != nil {
		http.Error(w, "Failed to marshal JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(jsonResponse)

}

func (h *Handler) GetRanking(w http.ResponseWriter, r *http.Request) {
	usersWithAnswers, err := h.service.Ranking.getUsersResults(h.ctx)
	if err != nil {
		http.Error(w, "Failed to get users with answers", http.StatusInternalServerError)
		return
	}

	sort.Slice(usersWithAnswers, func(i, j int) bool {
		return usersWithAnswers[i].Score > usersWithAnswers[j].Score
	})

	rank := make([]Ranking, 0)

	for _, user := range usersWithAnswers {
		rank = append(rank, Ranking{
			UserID:   user.ID,
			Username: user.Name,
			Score:    user.Score,
		})
	}

	// Serialize the response slice into JSON format
	jsonResponse, err := json.Marshal(rank)
	if err != nil {
		http.Error(w, "Failed to marshal JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(jsonResponse)
}

func (h *Handler) GetQuestions(w http.ResponseWriter, r *http.Request) {
	// Extract text and options from each question
	var questions []map[string]interface{}
	for _, q := range MultipleChoiceQuestions {
		question := map[string]interface{}{
			"id":       q.ID,
			"question": q.Question,
			"options":  q.Options,
		}
		questions = append(questions, question)
	}

	// Write the questions as JSON response
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(questions)
}
