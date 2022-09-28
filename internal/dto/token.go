package dto

import (
	"time"
	"wallet-api/internal/model"
)

type Token struct {
	Token     string
	UserID    uint64
	CreatedAt time.Time
	ExpiresAt time.Time
} // @name Token

func (dto *Token) ParseFromVO(token *model.Token) {
	dto.Token = token.Value
	dto.UserID = token.UserID
	dto.CreatedAt = token.CreatedAt
	dto.ExpiresAt = token.ExpiresAt
}
