package api

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockUserRepository is a mock implementation of UserRepository
type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) createUser(user User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserRepository) generateUserID(user User) (User, error) {
	args := m.Called(user)
	return args.Get(0).(User), args.Error(1)
}

func (m *MockUserRepository) generateSessionID(session Session) (Session, error) {
	args := m.Called(session)
	return args.Get(0).(Session), args.Error(1)
}

func (m *MockUserRepository) getUserByID(id uuid.UUID) (*User, error) {
	args := m.Called(id)
	return args.Get(0).(*User), args.Error(1)
}

func (m *MockUserRepository) getUsersResults() ([]User, error) {
	args := m.Called()
	return args.Get(0).([]User), args.Error(1)
}

func (m *MockUserRepository) updateUser(user *User) error {
	args := m.Called(user)
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

	err := userService.createUser(user)

	assert.NoError(t, err)
	repoMock.AssertExpectations(t)
}
