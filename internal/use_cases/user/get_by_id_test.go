package user_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"time"
	"wallet-api/internal/model"
	"wallet-api/internal/use_cases/user"

	"testing"
)

func Test_GetByID_OK(t *testing.T) {

	const userID uint64 = 1

	getByID := func(ID uint64) (bool, *model.User, error) {
		birthDate := "29-01-2000"
		birthToTime, err := time.Parse("02-01-2006", birthDate)
		require.NoError(t, err)

		assert.Equal(t, userID, ID)
		result := &model.User{
			ID:        1,
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
		return true, result, nil
	}

	condition, userCreate, err := user.GetByID(userID, getByID)
	require.NoError(t, err)

	assert.Equal(t, true, condition)
	assert.Equal(t, userID, userCreate.ID)
	assert.Equal(t, "Paulo", userCreate.FirstName)
	assert.Equal(t, "paulotest@gmail.com", userCreate.Email)
	assert.Equal(t, "paulo123", userCreate.Username)
}
