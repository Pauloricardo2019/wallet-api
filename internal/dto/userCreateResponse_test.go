package dto_test

import (
	"github.com/jaswdr/faker"
	"github.com/stretchr/testify/assert"
	"testing"
	"wallet-api/internal/dto"
	"wallet-api/internal/model"
)

func TestUserCreateResponse_ParseFromUserVO(t *testing.T) {
	fakerGen := faker.New()
	fakePerson := fakerGen.Person()

	usermodel := &model.User{
		FirstName: fakePerson.FirstName(),
		LastName:  fakePerson.LastName(),
		Email:     fakePerson.Contact().Email,
		Username:  fakerGen.Internet().User(),
	}

	userCreateResponseDTO := &dto.UserCreateResponse{}

	userCreateResponseDTO.ParseFromUserVO(usermodel)

	assert.Equal(t, usermodel.FirstName, userCreateResponseDTO.FirstName)
	assert.Equal(t, usermodel.LastName, userCreateResponseDTO.LastName)
	assert.Equal(t, usermodel.Email, userCreateResponseDTO.Email)
	assert.Equal(t, usermodel.Username, userCreateResponseDTO.Username)

}
