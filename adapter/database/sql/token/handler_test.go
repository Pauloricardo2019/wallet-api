//go:build integration_test
// +build integration_test

package token_test

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
	tokenRepo "wallet-api/adapter/database/sql/token"
	userRepo "wallet-api/adapter/database/sql/user"
	"wallet-api/internal/model"
)

func TestCreateToken_OK(t *testing.T) {
	birthDate := "29-01-2000"
	birthToTime, err := time.Parse("02-01-2006", birthDate)
	require.NoError(t, err)

	user := &model.User{
		FirstName: "Paulo",
		LastName:  "Ricardo",
		Email:     "paulotest@gmail.com",
		DDD:       "18",
		Phone:     "911223344",
		Username:  "paulo123",
		Birth:     birthToTime,
		Biography: "Like to walk and to practice sports",
		Password:  "P@ulo123",
	}
	userCreated, err := userRepo.Create(user)
	require.NoError(t, err)

	token := &model.Token{
		Value:     uuid.New().String(),
		UserID:    userCreated.ID,
		CreatedAt: time.Now(),
		ExpiresAt: time.Now().Add(time.Hour * 24),
	}

	tokenCreated, err := tokenRepo.CreateToken(token)
	require.NoError(t, err)

	assert.True(t, tokenCreated.UserID > 0)
	assert.Equal(t, userCreated.ID, tokenCreated.UserID)
	assert.Equal(t, token.Value, tokenCreated.Value)

}

func TestCreateToken_NotCreated(t *testing.T) {
	token := &model.Token{}

	tokenCreated, err := tokenRepo.CreateToken(token)
	require.Error(t, err)

	assert.Nil(t, tokenCreated)
}

func TestGetTokenByValue_OK(t *testing.T) {
	birthDate := "29-01-2000"
	birthToTime, err := time.Parse("02-01-2006", birthDate)
	require.NoError(t, err)

	user := &model.User{
		FirstName: "Paulo",
		LastName:  "Ricardo",
		Email:     "paulotest@gmail.com",
		DDD:       "18",
		Phone:     "911223344",
		Username:  "paulo123",
		Birth:     birthToTime,
		Biography: "Like to walk and to practice sports",
		Password:  "P@ulo123",
	}

	userCreated, err := userRepo.Create(user)
	require.NoError(t, err)

	token := &model.Token{
		Value:     uuid.New().String(),
		UserID:    userCreated.ID,
		CreatedAt: time.Now(),
		ExpiresAt: time.Now().Add(time.Hour * 24),
	}

	tokenCreated, err := tokenRepo.CreateToken(token)
	require.NoError(t, err)

	found, _, err := tokenRepo.GetTokenByValue(tokenCreated.Value)
	require.NoError(t, err)

	assert.True(t, found)
}

func TestGetTokenByValue_NotFound(t *testing.T) {
	found, _, err := tokenRepo.GetTokenByValue(uuid.New().String())
	require.NoError(t, err)

	assert.False(t, found)
}
