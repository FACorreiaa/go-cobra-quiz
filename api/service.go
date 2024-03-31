package api

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

type Service struct {
	repo *RepositoryStore
}

func NewService(repo *RepositoryStore) *Service {
	return &Service{repo: repo}
}

type SessionService interface {
	generateSessionID(ctx context.Context, session Session) (Session, error)
}

type QuizService interface {
	findQuestionByID(ctx context.Context, id int) *MultipleChoiceQuestion
	processUserAnswers(ctx context.Context, userAnswers map[string]string, user *User) (int, int, error)
}

type RankingService interface {
	getUsersResults(ctx context.Context) ([]User, error)
}

type UserService interface {
	createUser(ctx context.Context, user User) error
	generateUserID(ctx context.Context, user User) (User, error)
	getUserByID(ctx context.Context, id uuid.UUID) (*User, error)
	updateUserName(ctx context.Context, userID uuid.UUID, newName string) error
	updateUserScore(ctx context.Context, user *User, score int) error
	calculateUserPercent(ctx context.Context, user []User, score int) float64
}

type ServiceStore struct {
	User    UserService
	Session SessionService
	Quiz    QuizService
	Ranking RankingService
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

func (s *Service) createUser(ctx context.Context, user User) error {
	return s.repo.User.createUser(ctx, user)
}

func (s *Service) generateUserID(ctx context.Context, user User) (User, error) {
	return s.repo.User.generateUserID(ctx, user)
}

func (s *Service) generateSessionID(ctx context.Context, session Session) (Session, error) {
	return s.repo.User.generateSessionID(ctx, session)
}

func (s *Service) getUserByID(ctx context.Context, id uuid.UUID) (*User, error) {
	return s.repo.User.getUserByID(ctx, id)
}

func (s *Service) updateUserName(ctx context.Context, userID uuid.UUID, newName string) error {
	user, err := s.repo.User.getUserByID(ctx, userID)
	if err != nil {
		return fmt.Errorf("failed to retrieve user: %v", err)
	}

	user.Name = newName

	err = s.repo.User.updateUser(ctx, user)
	if err != nil {
		return fmt.Errorf("failed to update user: %v", err)
	}

	return nil
}

func (s *Service) getUsersResults(ctx context.Context) ([]User, error) {
	return s.repo.User.getUsersResults(ctx)
}

func (s *Service) updateUserScore(ctx context.Context, user *User, score int) error {
	// Update the user's score
	user.Score = score

	// Call the repository method to save the updated user information
	if err := s.repo.User.updateUser(ctx, user); err != nil {
		return err
	}

	return nil
}

func (s *Service) findQuestionByID(ctx context.Context, id int) *MultipleChoiceQuestion {
	return s.repo.User.findQuestionByID(ctx, id)
}

func (s *Service) calculateUserPercent(ctx context.Context, user []User, score int) float64 {
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

func (s *Service) processUserAnswers(ctx context.Context, userAnswers map[string]string, user *User) (int, int, error) {
	return s.repo.User.processUserAnswers(ctx, userAnswers, user)
}

func (u *User) hasAnswered(ctx context.Context, questionID int) bool {
	for _, ans := range u.Answers {
		if ans.QuestionID == questionID {
			return true
		}
	}
	return false
}
