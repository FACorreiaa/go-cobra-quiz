package api

import (
	"fmt"

	"github.com/google/uuid"
)

type UserRepository struct {
	users     map[uuid.UUID]*User
	questions []Question
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		users:     make(map[uuid.UUID]*User),
		questions: []Question{},
	}
}

type IUserRepository interface {
	generateUserID(user User) (User, error)
	generateSessionID(session Session) (Session, error)
	getUserByID(id uuid.UUID) (*User, error)
	getUsersResults() ([]User, error)
	updateUser(user *User) error
	createUser(user User) error
}

type Repository struct {
	User IUserRepository
}

func NewRepository() *Repository {
	return &Repository{
		User: NewUserRepository(),
	}
}

func (r *UserRepository) generateUserID(user User) (User, error) {
	// Save the user in the repository
	user.ID = uuid.New()
	return user, nil
}

func (r *UserRepository) generateSessionID(session Session) (Session, error) {
	// Save the user in the repository
	session.ID = uuid.New()
	return session, nil
}

func (r *UserRepository) getUserByID(id uuid.UUID) (*User, error) {
	user, ok := r.users[id]
	if !ok {
		return nil, fmt.Errorf("user not found")
	}
	userCopy := *user
	return &userCopy, nil
}

func (r *UserRepository) getUsersResults() ([]User, error) {
	var users []User
	for _, u := range r.users {
		users = append(users, *u)
	}
	return users, nil
}

func (r *UserRepository) updateUser(user *User) error {
	_, ok := r.users[user.ID]
	if !ok {
		return fmt.Errorf("user not found")
	}
	r.users[user.ID] = user
	return nil
}

func (r *UserRepository) createUser(user User) error {
	// Check if the user already exists
	_, ok := r.users[user.ID]
	if ok {
		return fmt.Errorf("user already exists")
	}

	// Add the user to the repository
	r.users[user.ID] = &user

	return nil
}
