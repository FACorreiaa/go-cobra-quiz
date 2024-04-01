//go:build ignore
// +build ignore

package internal

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockUserRepository is a mock implementation of UserRepository
// MockUserRepository TODO
type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) getUser(ctx context.Context, user User) error {
	args := m.Called(ctx, user)
	return args.Error(0)
}

func (m *MockUserRepository) getQuestionByID(ctx context.Context, id int) *MultipleChoiceQuestion {
	args := m.Called(ctx, id)
	return args.Get(0).(*MultipleChoiceQuestion)
}

func (m *MockUserRepository) createUserAnswers(ctx context.Context, userAnswers map[string]string, user *User) (int, int, error) {
	args := m.Called(ctx, userAnswers, user)
	return args.Int(0), args.Int(1), args.Error(2)
}

func (m *MockUserRepository) createUser(ctx context.Context, user User) error {
	args := m.Called(ctx, user)
	return args.Error(0)
}

func (m *MockUserRepository) createUserID(ctx context.Context, user User) (User, error) {
	args := m.Called(ctx, user)
	return args.Get(0).(User), args.Error(1)
}

func (m *MockUserRepository) createSessionID(ctx context.Context, session Session) (Session, error) {
	args := m.Called(ctx, session)
	return args.Get(0).(Session), args.Error(1)
}

func (m *MockUserRepository) getUserByID(ctx context.Context, id uuid.UUID) (*User, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*User), args.Error(1)
}

func (m *MockUserRepository) getUsersResults(ctx context.Context) ([]User, error) {
	args := m.Called(ctx)
	return args.Get(0).([]User), args.Error(1)
}

func (m *MockUserRepository) updateUser(ctx context.Context, user *User) error {
	args := m.Called(ctx, user)
	return args.Error(0)
}

func TestUserService_CreateUser(t *testing.T) {
	repoMock := new(MockUserRepository)
	ctx := context.Background()
	user := User{ID: uuid.New(), Name: "testuser"}
	repoMock.On("createUser", user).Return(nil)

	userService := &Service{
		repo: &RepositoryStore{
			User: repoMock,
		},
	}

	err := userService.getUser(ctx, user)

	assert.NoError(t, err)
	repoMock.AssertExpectations(t)
}
