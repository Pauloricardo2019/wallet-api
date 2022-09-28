package dto_test

import (
	"github.com/jaswdr/faker"
	"github.com/stretchr/testify/assert"
	"testing"
	"wallet-api/internal/dto"
)

func TestLoginRequest_ConvertToVO(t *testing.T) {
	fakerGen := faker.New()

	loginRequest := dto.LoginRequest{
		Email:    fakerGen.Internet().Email(),
		Password: fakerGen.Internet().Password(),
	}

	login := loginRequest.ConvertToVO()

	assert.Equal(t, loginRequest.Email, login.Email)
	assert.Equal(t, loginRequest.Password, login.Password)
}
