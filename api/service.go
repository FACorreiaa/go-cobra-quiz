package api

import (
	"fmt"

	"github.com/google/uuid"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) AddUser(user User) error {
	return s.repo.AddUser(user)
}

func (s *Service) GenerateUserID(user User) (User, error) {
	return s.repo.GenerateUserID(user)
}

func (s *Service) GenerateSessionID(session Session) (Session, error) {
	return s.repo.GenerateSessionID(session)
}

func (s *Service) GetUserByID(id uuid.UUID) (*User, error) {
	return s.repo.GetUserByID(id)
}

func (s *Service) UpdateUserName(userID uuid.UUID, newName string) error {
	user, err := s.repo.GetUserByID(userID)
	if err != nil {
		return fmt.Errorf("failed to retrieve user: %v", err)
	}

	user.Name = newName

	err = s.repo.UpdateUser(user)
	if err != nil {
		return fmt.Errorf("failed to update user: %v", err)
	}

	return nil
}

func (s *Service) GetUsers() ([]User, error) {
	return s.repo.GetUsers()
}

func (s *Service) UpdateUserScore(user *User, score int) error {
	// Update the user's score
	user.Score = score

	// Call the repository method to save the updated user information
	if err := s.repo.UpdateUser(user); err != nil {
		return err
	}

	return nil
}

func (s *Service) isValidAnswer(answer string, options []string) bool {
	for _, opt := range options {
		if answer == opt {
			return true
		}
	}
	return false
}

func (s *Service) findQuestionByID(id int) *MultipleChoiceQuestion {
	for _, q := range MultipleChoiceQuestions {
		if q.ID == id {
			return &q
		}
	}
	return nil
}

func (u *User) HasAnswered(questionID int) bool {
	for _, ans := range u.Answers {
		if ans.QuestionID == questionID {
			return true
		}
	}
	return false
}
