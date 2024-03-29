package api

import (
	"fmt"

	"github.com/google/uuid"
)

type Repository struct {
	users     map[uuid.UUID]*User
	questions []Question
}

func NewRepository() *Repository {
	return &Repository{
		users:     make(map[uuid.UUID]*User),
		questions: []Question{},
	}
}

func (r *Repository) generateUserID(user User) (User, error) {
	// Save the user in the repository
	user.ID = uuid.New()
	return user, nil
}

func (r *Repository) generateSessionID(session Session) (Session, error) {
	// Save the user in the repository
	session.ID = uuid.New()
	return session, nil
}

func (r *Repository) getUserByID(id uuid.UUID) (*User, error) {
	user, ok := r.users[id]
	if !ok {
		return nil, fmt.Errorf("user not found")
	}
	userCopy := *user
	return &userCopy, nil
}

func (r *Repository) getUsers() ([]User, error) {
	var users []User
	for _, u := range r.users {
		users = append(users, *u)
	}
	return users, nil
}

func (r *Repository) updateUser(user *User) error {
	_, ok := r.users[user.ID]
	if !ok {
		return fmt.Errorf("user not found")
	}
	r.users[user.ID] = user
	return nil
}

func (r *Repository) addUser(user User) error {
	// Check if the user already exists
	_, ok := r.users[user.ID]
	if ok {
		return fmt.Errorf("user already exists")
	}

	// Add the user to the repository
	r.users[user.ID] = &user

	return nil
}
