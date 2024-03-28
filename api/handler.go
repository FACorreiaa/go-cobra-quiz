package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type Handler struct {
	service *Service
}

func NewHandler(s *Service) *Handler {
	return &Handler{service: s}
}

func (h *Handler) StartSession(w http.ResponseWriter, r *http.Request) {
	var user User
	var session Session

	user, err := h.service.GenerateUserID(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = h.service.AddUser(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	session, err = h.service.GenerateSessionID(session)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	userID := user.ID
	sessionID := session.ID
	fmt.Printf("User ID: %v\n", userID)
	response := struct {
		UserID    uuid.UUID `json:"user_id"`
		SessionID uuid.UUID `json:"session_id"`
	}{UserID: userID, SessionID: sessionID}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(response)
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

	if err := h.service.UpdateUserName(userID, newName.Name); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := struct {
		UserID   uuid.UUID `json:"user_id"`
		Username string    `json:"username"`
	}{
		UserID:   userID,
		Username: newName.Name,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// SubmitQuiz work on this later
func (h *Handler) SubmitQuiz(w http.ResponseWriter, r *http.Request) {
	// Parse the request body to get the user's answers
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
	user, err := h.service.GetUserByID(userID)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	var score, correctAnswers int
	for id, answer := range userAnswers {
		questionID, err := strconv.Atoi(id)
		if err != nil {
			http.Error(w, "Invalid question ID", http.StatusBadRequest)
			return
		}
		question := h.service.findQuestionByID(questionID)
		if question == nil {
			http.Error(w, "Question not found", http.StatusNotFound)
			return
		}
		if user.HasAnswered(question.ID) {
			http.Error(w, "User has already answered this question", http.StatusBadRequest)
			return
		}
		if h.service.isValidAnswer(answer, question.Options) {
			correctAnswers++
		}
		user.Answers = append(user.Answers, Answer{QuestionID: question.ID, Answer: answer})
	}

	// Calculate the score
	score = correctAnswers * 10

	// Update user's score
	if err := h.service.UpdateUserScore(user, score); err != nil {
		http.Error(w, "Failed to save user information", http.StatusInternalServerError)
		return
	}

	//calculate user percentile
	usersWithAnswers, err := h.service.GetUsers()
	if err != nil {
		http.Error(w, "Failed to get users with answers", http.StatusInternalServerError)
		return
	}
	var higherScores int
	for _, u := range usersWithAnswers {
		if u.Score > score {
			higherScores++
		}
	}
	totalUsers := len(usersWithAnswers)
	percentile := (float64(higherScores) / float64(totalUsers)) * 100

	response := struct {
		Score          int     `json:"score"`
		CorrectAnswers int     `json:"correct_answers"`
		Percentile     float64 `json:"percentile"`
	}{
		Score:          score,
		CorrectAnswers: correctAnswers,
		Percentile:     percentile,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *Handler) GetAllScores(w http.ResponseWriter, r *http.Request) {
	usersWithAnswers, err := h.service.GetUsers()
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
	_, err = w.Write(jsonResponse)
	if err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
		return
	}
}

func (h *Handler) GetRanking(w http.ResponseWriter, r *http.Request) {
	usersWithAnswers, err := h.service.GetUsers()
	if err != nil {
		http.Error(w, "Failed to get users with answers", http.StatusInternalServerError)
		return
	}

	sort.Slice(usersWithAnswers, func(i, j int) bool {
		return usersWithAnswers[i].Score > usersWithAnswers[j].Score
	})

	var response []struct {
		UserID   uuid.UUID `json:"user_id"`
		Username string    `json:"username"`
		Score    int       `json:"score"`
	}

	for _, user := range usersWithAnswers {
		response = append(response, struct {
			UserID   uuid.UUID `json:"user_id"`
			Username string    `json:"username"`
			Score    int       `json:"score"`
		}{
			UserID:   user.ID,
			Username: user.Name,
			Score:    user.Score,
		})
	}

	// Serialize the response slice into JSON format
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Failed to marshal JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(jsonResponse)
	if err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
		return
	}
}