package dto_test

import (
	"github.com/jaswdr/faker"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
	"wallet-api/internal/dto"
	"wallet-api/internal/model"
)

func TestUser_ParseFromVO(t *testing.T) {
	fakerGen := faker.New()
	var urlmage = fakerGen.Internet().URL()

	usermodel := &model.User{
		ID:        fakerGen.UInt64(),
		UrlImage:  &urlmage,
		FirstName: fakerGen.Person().FirstName(),
		LastName:  fakerGen.Person().LastName(),
		Email:     fakerGen.Internet().Email(),
		DDD:       "18",
		Phone:     fakerGen.Person().Contact().Phone,
		Username:  fakerGen.Internet().User(),
		Birth:     time.Now(),
		Biography: fakerGen.Lorem().Paragraph(1),
		Password:  fakerGen.Internet().Password(),
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}

	userdto := &dto.User{}

	userdto.ParseFromVO(usermodel)

	assert.Equal(t, usermodel.ID, userdto.ID)
	assert.Equal(t, usermodel.UrlImage, userdto.UrlImage)
	assert.Equal(t, usermodel.FirstName, userdto.FirstName)
	assert.Equal(t, usermodel.LastName, userdto.LastName)
	assert.Equal(t, usermodel.Email, userdto.Email)
	assert.Equal(t, usermodel.DDD, userdto.DDD)
	assert.Equal(t, usermodel.Phone, userdto.Phone)
	assert.Equal(t, usermodel.Username, userdto.Username)
	assert.Equal(t, usermodel.Birth, userdto.Birth)
	assert.Equal(t, usermodel.Biography, userdto.Biography)
	assert.Equal(t, usermodel.Password, userdto.Password)
	assert.Equal(t, usermodel.CreatedAt, userdto.CreatedAt)
	assert.Equal(t, usermodel.UpdatedAt, userdto.UpdatedAt)
}

func TestUser_ConvertToVO(t *testing.T) {
	fakerGen := faker.New()
	var urlmage = fakerGen.Internet().URL()

	userdto := dto.User{
		ID:        fakerGen.UInt64(),
		UrlImage:  &urlmage,
		FirstName: fakerGen.Person().FirstName(),
		LastName:  fakerGen.Person().LastName(),
		Email:     fakerGen.Internet().Email(),
		DDD:       "18",
		Phone:     fakerGen.Person().Contact().Phone,
		Username:  fakerGen.Internet().User(),
		Birth:     time.Now(),
		Biography: fakerGen.Lorem().Paragraph(1),
		Password:  fakerGen.Internet().Password(),
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}

	usermodel := userdto.ConvertToVO()

	assert.Equal(t, usermodel.ID, userdto.ID)
	assert.Equal(t, usermodel.UrlImage, userdto.UrlImage)
	assert.Equal(t, usermodel.FirstName, userdto.FirstName)
	assert.Equal(t, usermodel.LastName, userdto.LastName)
	assert.Equal(t, usermodel.Email, userdto.Email)
	assert.Equal(t, usermodel.DDD, userdto.DDD)
	assert.Equal(t, usermodel.Phone, userdto.Phone)
	assert.Equal(t, usermodel.Username, userdto.Username)
	assert.Equal(t, usermodel.Birth, userdto.Birth)
	assert.Equal(t, usermodel.Biography, userdto.Biography)
	assert.Equal(t, usermodel.Password, userdto.Password)
	assert.Equal(t, usermodel.CreatedAt, userdto.CreatedAt)
	assert.Equal(t, usermodel.UpdatedAt, userdto.UpdatedAt)

}

func TestUser_ParseFromArrayVO(t *testing.T) {
	fakerGen := faker.New()
	var urlmage = fakerGen.Internet().URL()

	user1 := model.User{
		ID:        fakerGen.UInt64(),
		UrlImage:  &urlmage,
		FirstName: fakerGen.Person().FirstName(),
		LastName:  fakerGen.Person().LastName(),
		Email:     fakerGen.Internet().Email(),
		DDD:       "43",
		Phone:     fakerGen.Person().Contact().Phone,
		Username:  fakerGen.Internet().User(),
		Birth:     time.Now(),
		Biography: fakerGen.Lorem().Paragraph(1),
		Password:  fakerGen.Internet().Password(),
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}

	user2 := model.User{
		ID:        fakerGen.UInt64(),
		UrlImage:  &urlmage,
		FirstName: fakerGen.Person().FirstName(),
		LastName:  fakerGen.Person().LastName(),
		Email:     fakerGen.Internet().Email(),
		DDD:       "18",
		Phone:     fakerGen.Person().Contact().Phone,
		Username:  fakerGen.Internet().User(),
		Birth:     time.Now(),
		Biography: fakerGen.Lorem().Paragraph(1),
		Password:  fakerGen.Internet().Password(),
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}

	user3 := model.User{
		ID:        fakerGen.UInt64(),
		UrlImage:  &urlmage,
		FirstName: fakerGen.Person().FirstName(),
		LastName:  fakerGen.Person().LastName(),
		Email:     fakerGen.Internet().Email(),
		DDD:       "64",
		Phone:     fakerGen.Person().Contact().Phone,
		Username:  fakerGen.Internet().User(),
		Birth:     time.Now(),
		Biography: fakerGen.Lorem().Paragraph(1),
		Password:  fakerGen.Internet().Password(),
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}

	usersBeforeFunc := []model.User{
		user1,
		user2,
		user3,
	}

	users := []model.User{
		user1,
		user2,
		user3,
	}

	dtoUsers := &dto.User{}

	dtoUsers.ParseFromArrayVO(users)

	for i := 0; i < len(usersBeforeFunc); i++ {
		assert.Equal(t, users[i].ID, usersBeforeFunc[i].ID)
		assert.Equal(t, users[i].UrlImage, usersBeforeFunc[i].UrlImage)
		assert.Equal(t, users[i].FirstName, usersBeforeFunc[i].FirstName)
		assert.Equal(t, users[i].LastName, usersBeforeFunc[i].LastName)
		assert.Equal(t, users[i].Email, usersBeforeFunc[i].Email)
		assert.Equal(t, users[i].DDD, usersBeforeFunc[i].DDD)
		assert.Equal(t, users[i].Phone, usersBeforeFunc[i].Phone)
		assert.Equal(t, users[i].Username, usersBeforeFunc[i].Username)
		assert.Equal(t, users[i].Birth, usersBeforeFunc[i].Birth)
		assert.Equal(t, users[i].Biography, usersBeforeFunc[i].Biography)
		assert.Equal(t, users[i].Password, usersBeforeFunc[i].Password)
		assert.Equal(t, users[i].CreatedAt, usersBeforeFunc[i].CreatedAt)
		assert.Equal(t, users[i].UpdatedAt, usersBeforeFunc[i].UpdatedAt)
	}
}
