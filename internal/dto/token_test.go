package dto_test

import (
	"github.com/jaswdr/faker"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
	"wallet-api/internal/dto"
	"wallet-api/internal/model"
)

func TestToken_ParseFromVO(t *testing.T) {
	fakerGen := faker.New()

	usertoken := &model.Token{
		Value:     fakerGen.RandomStringWithLength(10),
		UserID:    fakerGen.UInt64(),
		CreatedAt: time.Time{},
		ExpiresAt: time.Time{},
	}

	tokenDTO := &dto.Token{}

	tokenDTO.ParseFromVO(usertoken)

	assert.Equal(t, usertoken.Value, tokenDTO.Token)
	assert.Equal(t, usertoken.UserID, tokenDTO.UserID)
	assert.Equal(t, usertoken.CreatedAt, tokenDTO.CreatedAt)
	assert.Equal(t, usertoken.ExpiresAt, tokenDTO.ExpiresAt)

}
