package internal

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"time"

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

type UserServiceRepository interface {
	createUserID(ctx context.Context, user User) (User, error)
	createSessionID(ctx context.Context, session Session) (Session, error)
	getUserByID(ctx context.Context, id uuid.UUID) (*User, error)
	getUsersResults(ctx context.Context) ([]User, error)
	updateUser(ctx context.Context, user *User) error
	getUser(ctx context.Context, user User) error
	getQuestionByID(ctx context.Context, id int) *MultipleChoiceQuestion
	createUserAnswers(ctx context.Context, userAnswers map[string]string, user *User) (int, int, error)
}

type RepositoryStore struct {
	User UserServiceRepository
}

func NewRepositoryStore() *RepositoryStore {
	return &RepositoryStore{
		User: NewUserRepository(),
	}
}

func (r *UserRepository) createUserID(ctx context.Context, user User) (User, error) {
	// Save the user in the repository
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	user.ID = uuid.New()
	return user, nil
}

func (r *UserRepository) createSessionID(ctx context.Context, session Session) (Session, error) {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	session.ID = uuid.New()
	return session, nil
}

func (r *UserRepository) getUserByID(ctx context.Context, id uuid.UUID) (*User, error) {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	user, ok := r.users[id]
	if !ok {
		return nil, fmt.Errorf("user not found")
	}
	userResponse := User{
		ID:   user.ID,
		Name: user.Name,
	}
	return &userResponse, nil
}

func (r *UserRepository) getUsersResults(ctx context.Context) ([]User, error) {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	var users []User
	for _, u := range r.users {
		users = append(users, *u)
	}
	return users, nil
}

func (r *UserRepository) updateUser(ctx context.Context, user *User) error {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	_, ok := r.users[user.ID]
	if !ok {
		return fmt.Errorf("user not found")
	}
	r.users[user.ID] = user
	return nil
}

func (r *UserRepository) getUser(ctx context.Context, user User) error {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	_, ok := r.users[user.ID]
	if ok {
		return fmt.Errorf("user already exists")
	}

	// Add the user to the repository
	r.users[user.ID] = &user

	return nil
}

func (r *UserRepository) createUserAnswers(ctx context.Context, userAnswers map[string]string, user *User) (int, int, error) {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	var score, correctAnswers int
	for id, answer := range userAnswers {
		questionID, err := strconv.Atoi(id)
		if err != nil {
			return 0, 0, fmt.Errorf("invalid question ID: %v", err)
		}
		question := r.getQuestionByID(ctx, questionID)
		if question == nil {
			return 0, 0, errors.New("question not found")
		}
		if user.hasAnswered(question.ID) {
			return 0, 0, errors.New("user has already answered this question")
		}
		if answer == question.CorrectAns {
			correctAnswers++
		}
		user.Answers = append(user.Answers, Answer{QuestionID: question.ID, Answer: answer})
	}

	score = correctAnswers * 10
	return score, correctAnswers, nil
}

func (r *UserRepository) getQuestionByID(ctx context.Context, id int) *MultipleChoiceQuestion {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	for _, q := range MultipleChoiceQuestions {
		if q.ID == id {
			return &q
		}
	}
	return nil
}
