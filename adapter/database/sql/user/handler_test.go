//go:build integration_test
// +build integration_test

package user_test

import (
	"github.com/google/uuid"
	"math"
	"testing"
	"time"
	"wallet-api/adapter/database/sql/album"
	userRepo "wallet-api/adapter/database/sql/user"
	"wallet-api/internal/model"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetByID_OK(t *testing.T) {
	birthDate := "29-01-2000"
	birthToTime, err := time.Parse("02-01-2006", birthDate)
	require.NoError(t, err)
	user := &model.User{
		UrlImage:  nil,
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

	createdUser, err := userRepo.Create(user)
	require.NoError(t, err)

	found, User, err := userRepo.GetByID(createdUser.ID)
	require.NoError(t, err)

	assert.True(t, found)
	assert.NotNil(t, User)
}

func TestGetByID_NoUser(t *testing.T) {
	found, _, err := userRepo.GetByID(321)
	require.NoError(t, err)
	assert.False(t, found)
}

func TestGetByLogin_OK(t *testing.T) {
	birthDate := "29-01-2000"
	birthToTime, err := time.Parse("02-01-2006", birthDate)
	require.NoError(t, err)
	user := &model.User{
		UrlImage:  nil,
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

	found, existingUser, err := userRepo.GetByLogin(userCreated.Username)
	require.NoError(t, err)
	assert.True(t, found)
	assert.Equal(t, "paulo123", existingUser.Username)
}

func TestGetByLogin_NotFound(t *testing.T) {
	username := uuid.New().String()
	found, _, err := userRepo.GetByLogin(username)
	require.NoError(t, err)
	assert.False(t, found)
}

func TestGetByEmail_OK(t *testing.T) {
	birthDate := "29-01-2000"
	birthToTime, err := time.Parse("02-01-2006", birthDate)
	require.NoError(t, err)

	user := &model.User{
		UrlImage:  nil,
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

	user, err = userRepo.Create(user)
	require.NoError(t, err)

	found, existingUser, err := userRepo.GetByEmail(user.Email)
	require.NoError(t, err)

	assert.True(t, found)
	assert.Equal(t, "paulotest@gmail.com", existingUser.Email)
}

func TestGetByEmail_NotFound(t *testing.T) {
	email := uuid.New().String() + "@digitalsys.com.br"
	found, _, err := userRepo.GetByEmail(email)
	require.NoError(t, err)

	assert.False(t, found)
}

func TestCreateUser_OK(t *testing.T) {
	birthDate := "29-01-2000"
	birthToTime, err := time.Parse("02-01-2006", birthDate)
	require.NoError(t, err)
	user := &model.User{
		UrlImage:  nil,
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

	createdUser, err := userRepo.Create(user)
	require.NoError(t, err)

	assert.NotNil(t, createdUser)
	assert.True(t, createdUser.ID > 0)
	assert.Equal(t, user.FirstName, createdUser.FirstName)
	assert.Equal(t, user.LastName, createdUser.LastName)
	assert.Equal(t, user.Email, createdUser.Email)
	assert.Equal(t, user.DDD, createdUser.DDD)
	assert.Equal(t, user.Phone, createdUser.Phone)
	assert.Equal(t, user.Username, createdUser.Username)
	assert.Equal(t, birthToTime, user.Birth)
	assert.Equal(t, user.Biography, createdUser.Biography)
}

func TestUpdateUser_OK(t *testing.T) {
	birthDate := "29-01-2000"
	birthToTime, err := time.Parse("02-01-2006", birthDate)
	require.NoError(t, err)

	user := &model.User{
		UrlImage:  nil,
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

	user, err = userRepo.Create(user)
	require.NoError(t, err)

	birthDate = "30-05-2001"
	birthToTime, err = time.Parse("02-01-2006", birthDate)
	require.NoError(t, err)
	user.Birth = birthToTime
	user.FirstName = "Changed FirstName"
	user.Email = "user_changed@digitalsys.com.br"

	err = userRepo.Update(user)
	require.NoError(t, err)

	_, user, err = userRepo.GetByID(user.ID)
	require.NoError(t, err)

	assert.Equal(t, "Changed FirstName", user.FirstName)
	assert.Equal(t, "user_changed@digitalsys.com.br", user.Email)
}

func TestUpdateUser_NotFound(t *testing.T) {

	birthDate := "30-05-2001"
	birthToTime, err := time.Parse("02-01-2006", birthDate)
	require.NoError(t, err)

	user := &model.User{
		UrlImage:  nil,
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

	user, err = userRepo.Create(user)
	require.NoError(t, err)

	user.FirstName = "Changed Name"
	user.Email = "user_changed@digitalsys.com.br"

	err = userRepo.Update(user)
	require.NoError(t, err)

	found, _, err := userRepo.GetByID(math.MaxUint32)
	require.NoError(t, err)

	assert.False(t, found)
}

func TestDeleteUser_OK(t *testing.T) {

	birthDate := "30-05-2001"
	birthToTime, err := time.Parse("02-01-2006", birthDate)
	require.NoError(t, err)

	user := &model.User{
		UrlImage:  nil,
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

	user, err = userRepo.Create(user)
	require.NoError(t, err)

	err = userRepo.Delete(user.ID)
	require.NoError(t, err)

	found, _, err := userRepo.GetByID(user.ID)
	require.NoError(t, err)
	require.False(t, found)
}

func TestDeleteUser_NotDeleted(t *testing.T) {

	birthDate := "30-05-2001"
	birthToTime, err := time.Parse("02-01-2006", birthDate)
	require.NoError(t, err)

	user := &model.User{
		UrlImage:  nil,
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

	user, err = userRepo.Create(user)
	require.NoError(t, err)

	err = album.Delete(math.MaxInt32)
	require.NoError(t, err)
}
