package user

import (
	"testing"
	"time"
	"wallet-api/internal/model"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_Create_OK(t *testing.T) {

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

	createUser := func(user *model.User) (*model.User, error) {
		user.ID = 123
		return user, nil
	}

	getByLogin := func(login string) (bool, *model.User, error) {
		newUser := &model.User{}
		return true, newUser, nil
	}

	getByEmail := func(email string) (bool, *model.User, error) {
		newUser := &model.User{}
		return true, newUser, nil
	}
	userDTO, err := Create(user, createUser, getByEmail, getByLogin)
	require.NoError(t, err)

	assert.Equal(t, "Paulo", userDTO.FirstName)
	assert.Equal(t, "Ricardo", userDTO.LastName)
	assert.Equal(t, birthToTime, user.Birth)
	assert.Equal(t, "paulotest@gmail.com", userDTO.Email)
}

func Test_Create_EmailAlreadyExists(t *testing.T) {
	birthDate := "29-01-2000"
	birthToTime, err := time.Parse("02-01-2006", birthDate)
	require.NoError(t, err)

	user := &model.User{
		ID:        123,
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

	getByLogin := func(login string) (bool, *model.User, error) {
		newUser := &model.User{}
		return true, newUser, nil
	}

	getByEmail := func(email string) (bool, *model.User, error) {
		user.Email = "paulotest@gmail.com"
		return true, user, nil
	}

	createUser := func(*model.User) (*model.User, error) {
		user.ID = 123
		return user, nil
	}

	_, err = Create(user, createUser, getByEmail, getByLogin)
	require.Error(t, err)

	assert.Equal(t, "invalid user. email already exists", err.Error())
}

func Test_Create_LoginAlreadyExists(t *testing.T) {
	birthDate := "29-01-2000"
	birthToTime, err := time.Parse("02-01-2006", birthDate)
	require.NoError(t, err)

	user := &model.User{
		ID:        123,
		FirstName: "Paulo",
		LastName:  "Ricardo",
		Email:     "paulotest@gmail.com",
		DDD:       "21",
		Phone:     "911223344",
		Username:  "paulo123",
		Birth:     birthToTime,
		Biography: "Like to walk and to practice sports",
		Password:  "P@ulo123",
	}

	getByLogin := func(login string) (bool, *model.User, error) {
		user.Username = "paulo123"
		return true, user, nil
	}

	getByEmail := func(email string) (bool, *model.User, error) {
		newUser := &model.User{}
		return true, newUser, nil
	}

	createUser := func(*model.User) (*model.User, error) {
		user.ID = 123
		return user, nil
	}

	_, err = Create(user, createUser, getByEmail, getByLogin)
	require.Error(t, err)

	assert.Equal(t, "invalid user. auth already exists", err.Error())

}

func Test_UserValidation_OK(t *testing.T) {
	birthDate := "29-01-2000"
	birthToTime, err := time.Parse("02-01-2006", birthDate)
	require.NoError(t, err)

	newUser := &model.User{
		ID:        123,
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

	err = validateUser(newUser)
	require.NoError(t, err)
}

func Test_UserValidation_InvalidName(t *testing.T) {
	tests := []struct {
		TestName             string
		User                 *model.User
		ExpectedErrorMessage string
	}{
		{
			TestName: "Invalid_Name",
			User: &model.User{
				FirstName: "",
				LastName:  "Ricardo",
				Email:     "paulotest@gmail.com",
				DDD:       "18",
				Phone:     "911223344",
				Username:  "paulo123",
				Birth:     time.Now(),
				Biography: "Like to walk and to practice sports",
				Password:  "P@ulo123",
				CreatedAt: time.Time{},
				UpdatedAt: time.Time{},
			},
			ExpectedErrorMessage: "invalid user. field FirstName cannot to be empty",
		},
		{
			TestName: "Invalid_Name_Less_Than_4_Characters",
			User: &model.User{
				FirstName: "tes",
				LastName:  "Ricardo",
				Email:     "paulotest@gmail.com",
				DDD:       "18",
				Phone:     "911223344",
				Username:  "paulo123",
				Birth:     time.Now(),
				Biography: "Like to walk and to practice sports",
				Password:  "P@ulo123",
				CreatedAt: time.Time{},
				UpdatedAt: time.Time{},
			},
			ExpectedErrorMessage: "invalid user. FirstName cannot be less than 4 characters",
		},
		{
			TestName: "Invalid_Name_Longer_Than_100_Characters",
			User: &model.User{
				FirstName: "Paulo Paulo Paulo Paulo Paulo Paulo Paulo Paulo Paulo Paulo Paulo Paulo Paulo Paulo Paulo Paulo Paulo Paulo Paulo Paulo Paulo Paulo Paulo Paulo Paulo Paulo Paulo Paulo Paulo Paulo Paulo",
				LastName:  "Ricardo",
				Email:     "paulotest@gmail.com",
				DDD:       "18",
				Phone:     "911223344",
				Username:  "paulo123",
				Birth:     time.Now(),
				Biography: "Like to walk and to practice sports",
				Password:  "P@ulo123",
				CreatedAt: time.Time{},
				UpdatedAt: time.Time{},
			},
			ExpectedErrorMessage: "invalid user. FirstName cannot be longer than 100 characters",
		},
		{
			TestName: "Invalid_Name",
			User: &model.User{
				FirstName: "Paulo",
				LastName:  "",
				Email:     "paulotest@gmail.com",
				DDD:       "18",
				Phone:     "911223344",
				Username:  "paulo123",
				Birth:     time.Now(),
				Biography: "Like to walk and to practice sports",
				Password:  "P@ulo123",
				CreatedAt: time.Time{},
				UpdatedAt: time.Time{},
			},
			ExpectedErrorMessage: "invalid user. field LastName cannot to be empty",
		},
		{
			TestName: "Invalid_Name_Less_Than_4_Characters",
			User: &model.User{
				FirstName: "Paulo",
				LastName:  "Ric",
				Email:     "paulotest@gmail.com",
				DDD:       "18",
				Phone:     "911223344",
				Username:  "paulo123",
				Birth:     time.Now(),
				Biography: "Like to walk and to practice sports",
				Password:  "P@ulo123",
				CreatedAt: time.Time{},
				UpdatedAt: time.Time{},
			},
			ExpectedErrorMessage: "invalid user. LastName cannot be less than 4 characters",
		},
		{
			TestName: "Invalid_Name_Longer_Than_100_Characters",
			User: &model.User{
				FirstName: "Paulo",
				LastName:  "Ricardo Ricardo Ricardo Ricardo Ricardo Ricardo Ricardo Ricardo Ricardo Ricardo Ricardo Ricardo Ricardo Ricardo Ricardo Ricardo Ricardo Ricardo Ricardo Ricardo Ricardo Ricardo Ricardo Ricardo Ricardo Ricardo",
				Email:     "paulotest@gmail.com",
				DDD:       "18",
				Phone:     "911223344",
				Username:  "paulo123",
				Birth:     time.Now(),
				Biography: "Like to walk and to practice sports",
				Password:  "P@ulo123",
				CreatedAt: time.Time{},
				UpdatedAt: time.Time{},
			},
			ExpectedErrorMessage: "invalid user. LastName cannot be longer than 100 characters",
		},
		{
			TestName: "Invalid_Email",
			User: &model.User{
				FirstName: "Paulo",
				LastName:  "Ricardo",
				Email:     "",
				DDD:       "18",
				Phone:     "911223344",
				Username:  "paulo123",
				Birth:     time.Now(),
				Biography: "Like to walk and to practice sports",
				Password:  "P@ulo123",
				CreatedAt: time.Time{},
				UpdatedAt: time.Time{},
			},
			ExpectedErrorMessage: "invalid user. field email cannot to be empty",
		},
		{
			TestName: "Invalid_Login",
			User: &model.User{
				FirstName: "Paulo",
				LastName:  "Ricardo",
				Email:     "paulotest@gmail.com",
				DDD:       "18",
				Phone:     "911223344",
				Username:  "",
				Birth:     time.Now(),
				Biography: "Like to walk and to practice sports",
				Password:  "P@ulo123",
				CreatedAt: time.Time{},
				UpdatedAt: time.Time{},
			},
			ExpectedErrorMessage: "invalid user. field auth cannot to be empty",
		},
		{
			TestName: "Invalid_Login_Less_Than_4_Characters",
			User: &model.User{
				FirstName: "Paulo",
				LastName:  "Ricardo",
				Email:     "paulotest@gmail.com",
				DDD:       "18",
				Phone:     "911223344",
				Username:  "p",
				Birth:     time.Now(),
				Biography: "Like to walk and to practice sports",
				Password:  "P@ulo123",
				CreatedAt: time.Time{},
				UpdatedAt: time.Time{},
			},
			ExpectedErrorMessage: "invalid user. auth cannot be less than 4 characters",
		},
		{
			TestName: "Invalid_Login_Longer_Than_50_Characters",
			User: &model.User{
				FirstName: "Paulo",
				LastName:  "Ricardo",
				Email:     "paulotest@gmail.com",
				DDD:       "18",
				Phone:     "911223344",
				Username:  "paulo123 paulo123 paulo123 paulo123 paulo123 paulo123 paulo123 paulo123 paulo123 paulo123 paulo123 paulo123 paulo123 paulo123 paulo123",
				Birth:     time.Now(),
				Biography: "Like to walk and to practice sports",
				Password:  "P@ulo123",
				CreatedAt: time.Time{},
				UpdatedAt: time.Time{},
			},
			ExpectedErrorMessage: "invalid user. auth cannot be longer than 50 characters",
		},
	}

	for _, tt := range tests {
		t.Run(tt.TestName, func(t *testing.T) {
			err := validateUser(tt.User)
			require.NotNil(t, err)
			assert.Equal(t, tt.ExpectedErrorMessage, err.Error())
		})
	}
}

func Test_PasswordValidation_InvalidPassword(t *testing.T) {
	tests := []struct {
		TestName             string
		Password             string
		ExpectedErrorMessage string
	}{
		{
			TestName:             "Invalid_Password_Empty",
			Password:             "",
			ExpectedErrorMessage: "invalid user. password cannot to be empty",
		},
		{
			TestName:             "Invalid_Password_Be_Less_Than_7_Characters",
			Password:             "teste",
			ExpectedErrorMessage: "invalid user. this password must be 7 characters long",
		},
		{
			TestName:             "Invalid_Password_1_letter",
			Password:             "987454575",
			ExpectedErrorMessage: "invalid user. this password must be at least 1 letter",
		},
		{
			TestName:             "Invalid_Password_1_Uppercase_letter",
			Password:             "teste1234566",
			ExpectedErrorMessage: "invalid user. this password must be at least 1 uppercase letter",
		},
		{
			TestName:             "Invalid_Password_1_Number",
			Password:             "testeTEST&",
			ExpectedErrorMessage: "invalid user. this password must be at least 1 number",
		},
		{
			TestName:             "Invalid_Password_1_Special_Character",
			Password:             "Teste123456",
			ExpectedErrorMessage: "invalid user. this password must be at least 1 special character",
		},
	}

	for _, tt := range tests {
		t.Run(tt.TestName, func(t *testing.T) {
			err := validatePassword(tt.Password)
			require.NotNil(t, err)
			assert.Equal(t, tt.ExpectedErrorMessage, err.Error())
		})
	}
}
