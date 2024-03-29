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

func (s *Service) addUser(user User) error {
	return s.repo.addUser(user)
}

func (s *Service) generateUserID(user User) (User, error) {
	return s.repo.generateUserID(user)
}

func (s *Service) generateSessionID(session Session) (Session, error) {
	return s.repo.generateSessionID(session)
}

func (s *Service) getUserByID(id uuid.UUID) (*User, error) {
	return s.repo.getUserByID(id)
}

func (s *Service) updateUserName(userID uuid.UUID, newName string) error {
	user, err := s.repo.getUserByID(userID)
	if err != nil {
		return fmt.Errorf("failed to retrieve user: %v", err)
	}

	user.Name = newName

	err = s.repo.updateUser(user)
	if err != nil {
		return fmt.Errorf("failed to update user: %v", err)
	}

	return nil
}

func (s *Service) getUsers() ([]User, error) {
	return s.repo.getUsers()
}

func (s *Service) updateUserScore(user *User, score int) error {
	// Update the user's score
	user.Score = score

	// Call the repository method to save the updated user information
	if err := s.repo.updateUser(user); err != nil {
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

func (s *Service) calculateUserPercent(user []User, score int) float64 {
	var higherScores int
	for _, u := range user {
		if u.Score > score {
			higherScores++
		}
	}

	totalUsers := len(user)
	percentile := (float64(higherScores) / float64(totalUsers)) * 100
	return percentile
}

func (u *User) hasAnswered(questionID int) bool {
	for _, ans := range u.Answers {
		if ans.QuestionID == questionID {
			return true
		}
	}
	return false
}
