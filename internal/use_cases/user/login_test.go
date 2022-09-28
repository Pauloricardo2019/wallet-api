package user_test

/*
func TestLoginRequest_OK(t *testing.T) {
	loginRequest := &model.LoginRequest{
		Email:    "example@gmail.com",
		Password: "Ex@mple123",
	}

	user := &model.User{
		ID: 123,
	}

	getEmail := func(email string) (bool, *model.User, error) {
		user := &model.User{
			ID:             123,
			FirstName:      "Paulo",
			LastName:       "Ricardo",
			Email:          "example@gmail.com",
			DDD:            "18",
			Phone:          "911223344",
			Username:       "paulo123",
			City:           "Los Angeles",
			Biography:      "Like to walk and to practice sports",
			HashedPassword: "ebd81e2a400b8a0775217e59eed74257d4238b198b00ef592d5f7162372edb5d",
		}
		return true, user, nil
	}

	found, token, err := userUseCase.Login(loginRequest, getEmail)
	require.NoError(t, err)

	assert.True(t, found)
	assert.Equal(t, user.ID, token.UserID)
}

func TestLoginRequest_NotFound(t *testing.T) {
	loginRequest := &model.LoginRequest{
		Email:    "example@gmail.com",
		Password: "Ex@mple123",
	}

	getEmail := func(email string) (bool, *model.User, error) {
		user := &model.User{}
		return true, user, nil
	}

	found, _, err := loginUseCase.ConnectUser(loginRequest, getEmail)
	require.Error(t, err)

	assert.False(t, found)

}
*/
