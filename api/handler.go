package api

import (
	"encoding/json"
	"fmt"
	"net/http"

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
	// Parse the user ID from the request parameters
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

	user, err := h.service.GetUserByID(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	user.Name = newName.Name

	if err := h.service.UpdateUser(user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonResponse := map[string]string{"message": "User name updated successfully!"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(jsonResponse)
}
