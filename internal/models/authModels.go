package models

import (
	"github.com/google/uuid"
)

type RegisterModel struct {
	Email    string `json:"email" binding:"required,email"`
	Telegram string `json:"telegram" binding:"required"`
	Password string `json:"password" binding:"required,min=8"`
}

type LoginModel struct {
	Email       *string `json:"email"`
	Telegram    *string `json:"telegram"`
	Password    *string `json:"password" binding:"required"`
	Fingerprint string  `json:"fingerptint" binding:"required"`
}

type TokensModel struct {
	AccessToken  string
	RefreshToken uuid.UUID
}

type CreateSessionModel struct {
	UserID       uuid.UUID
	RefreshToken uuid.UUID
	Ua           string
	Fingerprint  string
	IP           string
	ExpiresAt    int64
}
