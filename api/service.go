package api

import (
	"fmt"

	"github.com/google/uuid"
)

type Service struct {
	repo *RepositoryStore
}

func NewService(repo *RepositoryStore) *Service {
	return &Service{repo: repo}
}

// ISession I woulnd't usually name ISession but since im
// doing everything on this file, it gets easier to understand
type ISession interface {
	generateSessionID(session Session) (Session, error)
}

type IQuiz interface {
	findQuestionByID(id int) *MultipleChoiceQuestion
	processUserAnswers(userAnswers map[string]string, user *User) (int, int, error)
}

type IRanking interface {
	getUsersResults() ([]User, error)
}

type IUser interface {
	createUser(user User) error
	generateUserID(user User) (User, error)
	getUserByID(id uuid.UUID) (*User, error)
	updateUserName(userID uuid.UUID, newName string) error
	updateUserScore(user *User, score int) error
	calculateUserPercent(user []User, score int) float64
}

type ServiceStore struct {
	User    IUser
	Session ISession
	Quiz    IQuiz
	Ranking IRanking
}

func NewServiceStore(repo *RepositoryStore) *ServiceStore {
	return &ServiceStore{
		User:    NewService(repo),
		Session: NewService(repo),
		Quiz:    NewService(repo),
		Ranking: NewService(repo),
	}
}

// Implementation would generally be on a different file

func (s *Service) createUser(user User) error {
	return s.repo.User.createUser(user)
}

func (s *Service) generateUserID(user User) (User, error) {
	return s.repo.User.generateUserID(user)
}

func (s *Service) generateSessionID(session Session) (Session, error) {
	return s.repo.User.generateSessionID(session)
}

func (s *Service) getUserByID(id uuid.UUID) (*User, error) {
	return s.repo.User.getUserByID(id)
}

func (s *Service) updateUserName(userID uuid.UUID, newName string) error {
	user, err := s.repo.User.getUserByID(userID)
	if err != nil {
		return fmt.Errorf("failed to retrieve user: %v", err)
	}

	user.Name = newName

	err = s.repo.User.updateUser(user)
	if err != nil {
		return fmt.Errorf("failed to update user: %v", err)
	}

	return nil
}

func (s *Service) getUsersResults() ([]User, error) {
	return s.repo.User.getUsersResults()
}

func (s *Service) updateUserScore(user *User, score int) error {
	// Update the user's score
	user.Score = score

	// Call the repository method to save the updated user information
	if err := s.repo.User.updateUser(user); err != nil {
		return err
	}

	return nil
}

func (s *Service) findQuestionByID(id int) *MultipleChoiceQuestion {
	return s.repo.Quiz.findQuestionByID(id)
}

func (s *Service) calculateUserPercent(user []User, score int) float64 {
	var betterUsers int
	for _, u := range user {
		if u.Score > score {
			betterUsers++
		}
	}

	totalUsers := len(user)

	if totalUsers == 1 {
		return 100
	}

	if betterUsers == 0 {
		return 100
	}

	percentile := float64(betterUsers) / float64(totalUsers) * 100

	return percentile
}

func (s *Service) processUserAnswers(userAnswers map[string]string, user *User) (int, int, error) {
	return s.repo.Quiz.processUserAnswers(userAnswers, user)
}
func (u *User) hasAnswered(questionID int) bool {
	for _, ans := range u.Answers {
		if ans.QuestionID == questionID {
			return true
		}
	}
	return false
}
