package api

import (
	"context"
	"errors"
	"fmt"
	"sort"

	"github.com/google/uuid"
)

type Service struct {
	repo *RepositoryStore
}

func NewService(repo *RepositoryStore) *Service {
	return &Service{repo: repo}
}

type SessionService interface {
	createSessionID(ctx context.Context, session Session) (Session, error)
}

type QuizService interface {
	getQuestionByID(ctx context.Context, id int) *MultipleChoiceQuestion
	createUserAnswers(ctx context.Context, userAnswers map[string]string, user *User) (int, int, error)
}

type RankingService interface {
	getUsersResults(ctx context.Context) ([]User, error)
}

type UserService interface {
	getUser(ctx context.Context, user User) error
	createUserID(ctx context.Context, user User) (User, error)
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

func (s *Service) getUser(ctx context.Context, user User) error {
	if user.ID.String() == "" {
		return errors.New("user id wasn't generated")
	}
	return s.repo.User.getUser(ctx, user)
}

func (s *Service) createUserID(ctx context.Context, user User) (User, error) {
	if user.ID.String() == "" {
		return User{}, errors.New("no id found")
	}
	return s.repo.User.createUserID(ctx, user)
}

func (s *Service) createSessionID(ctx context.Context, session Session) (Session, error) {
	if session.ID.String() == "" {
		return Session{}, errors.New("session id wasn't generated")
	}
	return s.repo.User.createSessionID(ctx, session)
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

func (s *Service) getQuestionByID(ctx context.Context, id int) *MultipleChoiceQuestion {
	return s.repo.User.getQuestionByID(ctx, id)
}

func (s *Service) calculateUserPercent(ctx context.Context, user []User, score int) float64 {
	var betterUsers int

	sort.Slice(user, func(i, j int) bool {
		return user[i].Score > user[j].Score
	})

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

func (s *Service) createUserAnswers(ctx context.Context, userAnswers map[string]string, user *User) (int, int, error) {
	if len(userAnswers) == 0 {
		return 0, 0, nil
	}
	return s.repo.User.createUserAnswers(ctx, userAnswers, user)
}

func (u *User) hasAnswered(questionID int) bool {
	for _, ans := range u.Answers {
		if ans.QuestionID == questionID {
			return true
		}
	}
	return false
}
