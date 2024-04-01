package internal

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestNewUserRepository_CreateSessionID(t *testing.T) {
	repo := NewUserRepository()
	session := Session{}

	session, err := repo.createSessionID(context.Background(), session)
	require.NoError(t, err)
	require.NotNil(t, session.ID)

}
func TestUserRepository_CreateUserID(t *testing.T) {
	repo := NewUserRepository()
	user := User{}
	user.Name = "TestUser"
	createdUser, err := repo.createUserID(context.Background(), user)

	require.NoError(t, err)
	require.NotNil(t, createdUser.ID)
}

func TestUserRepository_UserAlreadyAnswered(t *testing.T) {
	repo := NewUserRepository()

	user := &User{
		ID:      uuid.New(),
		Name:    "TestUser",
		Answers: []Answer{{QuestionID: 1, Answer: "A"}},
	}

	userAnswers := map[string]string{
		"1": "B",
	}

	_, _, err := repo.createUserAnswers(context.Background(), userAnswers, user)

	require.Error(t, err)
	require.Equal(t, "user has already answered this question", err.Error())
}
