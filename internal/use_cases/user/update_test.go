package user_test

import (
	"github.com/google/uuid"
	"testing"
	"time"
	"wallet-api/internal/model"
	"wallet-api/internal/use_cases/user"

	"github.com/stretchr/testify/require"
)

func Test_Update_OK(t *testing.T) {

	const userID uint64 = 1

	newUser := &model.User{
		ID:        1,
		FirstName: "Paulo",
		LastName:  "Ricardo",
		Email:     "t1334afafa56@gmail.com",
		DDD:       "18",
		Phone:     "911223344",
		Username:  "paulo123",
		Birth:     time.Now(),
		Biography: "Like to walk and to practice sports",
		Password:  "P@ulo123",
	}

	updateUser := func(*model.User) error {
		newUser.ID = 1
		newUser.FirstName = "ChangeName"
		newUser.LastName = "changeeLast"
		newUser.Email = uuid.New().String() + "@gmail.com"
		newUser.DDD = "21"
		newUser.Phone = "999886622"
		newUser.Username = "update123"
		newUser.Biography = "Like to drink and to run"

		return nil
	}

	getByEmail := func(email string) (bool, *model.User, error) {
		newUser := &model.User{}
		return true, newUser, nil
	}

	err := user.Update(userID, newUser, updateUser, getByEmail)
	require.NoError(t, err)

}
