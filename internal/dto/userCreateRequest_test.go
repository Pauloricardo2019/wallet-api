package dto_test

import (
	"github.com/jaswdr/faker"
	"github.com/stretchr/testify/assert"
	"testing"
	"wallet-api/internal/dto"
)

func TestUserCreateRequest_GenerateUserVO(t *testing.T) {
	fakerGen := faker.New()
	fakePerson := fakerGen.Person()

	userCreateRequest := dto.UserCreateRequest{
		FirstName: fakePerson.FirstName(),
		LastName:  fakePerson.LastName(),
		Email:     fakePerson.Contact().Email,
		Username:  fakerGen.Internet().User(),
		Password:  fakerGen.Internet().Password(),
	}

	user := userCreateRequest.GenerateUserVO()

	assert.Equal(t, userCreateRequest.FirstName, user.FirstName)
	assert.Equal(t, userCreateRequest.LastName, user.LastName)
	assert.Equal(t, userCreateRequest.Email, user.Email)
	assert.Equal(t, userCreateRequest.Username, user.Username)
	assert.Equal(t, userCreateRequest.Password, user.Password)
}
