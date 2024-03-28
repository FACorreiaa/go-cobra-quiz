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

func (r *Repository) calculateScore(answers []Answer, questions []MultipleChoiceQuestion) int {
	score := 0
	for _, ans := range answers {
		for _, q := range questions {
			if q.ID == ans.QuestionID && q.CorrectAns == ans.Answer {
				score += 10
				break
			}
		}
	}
	return score
}

func (r *Repository) GenerateUserID(user User) (User, error) {
	// Save the user in the repository
	user.ID = uuid.New()
	return user, nil
}

func (r *Repository) GenerateSessionID(session Session) (Session, error) {
	// Save the user in the repository
	session.ID = uuid.New()
	return session, nil
}

func (r *Repository) GetUserByID(id uuid.UUID) (*User, error) {
	user, ok := r.users[id]
	if !ok {
		return nil, fmt.Errorf("user not found")
	}
	userCopy := *user
	return &userCopy, nil
}

func (r *Repository) GetUsers() ([]User, error) {
	var users []User
	for _, u := range r.users {
		users = append(users, *u)
	}
	return users, nil
}

func (r *Repository) UpdateUser(user *User) error {
	_, ok := r.users[user.ID]
	if !ok {
		return fmt.Errorf("user not found")
	}
	r.users[user.ID] = user
	return nil
}

func (r *Repository) AddUser(user User) error {
	// Check if the user already exists
	_, ok := r.users[user.ID]
	if ok {
		return fmt.Errorf("user already exists")
	}

	// Add the user to the repository
	r.users[user.ID] = &user

	return nil
}
