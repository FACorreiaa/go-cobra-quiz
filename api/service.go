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

// ISession I woulnd't usually name ISession but since im
// doing everything on this file, it gets easier to understand
type ISession interface {
	generateSessionID(session Session) (Session, error)
}

type IQuiz interface {
	findQuestionByID(id int) *MultipleChoiceQuestion
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

type ServiceUser struct {
	User    IUser
	Session ISession
	Quiz    IQuiz
	Ranking IRanking
}

func NewQuizService(repo *Repository) *ServiceUser {
	return &ServiceUser{
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
	for _, q := range MultipleChoiceQuestions {
		if q.ID == id {
			return &q
		}
	}
	return nil
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

func (u *User) hasAnswered(questionID int) bool {
	for _, ans := range u.Answers {
		if ans.QuestionID == questionID {
			return true
		}
	}
	return false
}
