package api

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockUserRepository is a mock implementation of UserRepository
type MockUserRepository struct {
	mock.Mock
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
	sessionID, err := m.Called(ctx, session).Get(0).(Session), m.Called(session).Error(1)
	return sessionID, err
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

	user := User{ID: uuid.New(), Name: "testuser"}
	repoMock.On("createUser", user).Return(nil)

	userService := &Service{
		repo: &RepositoryStore{
			User: repoMock,
		},
	}

	err := userService.createUser(context.Background(), user)

	assert.NoError(t, err)
	repoMock.AssertExpectations(t)
}

func TestUpdateUserName_ValidName(t *testing.T) {
	repoMock := new(MockUserRepository)

	userID := uuid.New()
	user := &User{ID: userID}
	repoMock.On("getUserByID", userID).Return(user, nil)

	repoMock.On("updateUser", user).Return(nil)

	userService := &Service{
		repo: &RepositoryStore{
			User: repoMock,
		},
	}

	newName := "newName"
	err := userService.updateUserName(context.Background(), userID, newName)

	assert.NoError(t, err)

	assert.Equal(t, newName, user.Name)
	assert.NotEqualValues(t, 123, user.Name)

	repoMock.AssertCalled(t, "updateUser", user)
}
