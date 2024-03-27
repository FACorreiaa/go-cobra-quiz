package api

import (
	"fmt"

	"github.com/google/uuid"
)

type Repository struct {
	users     map[uuid.UUID]*User
	questions []Question
	responses []Response
}

func NewRepository() *Repository {
	return &Repository{
		users:     make(map[uuid.UUID]*User),
		questions: []Question{},
		responses: []Response{},
	}
}

func (r *Repository) SaveResponse(response Response) (Response, error) {
	r.responses = append(r.responses, response)
	return response, nil
}

func (r *Repository) calculateScore(answers []Answer) int {
	score := 0
	for _, ans := range answers {
		if ans.Response == "correct" {
			score += 10
		}
	}
	return score
}

//

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
	return user, nil
}

func (r *Repository) UpdateUser(user *User) error {
	_, ok := r.users[user.ID]
	if !ok {
		return fmt.Errorf("user not found")
	}
	r.users[user.ID] = user
	return nil
}
