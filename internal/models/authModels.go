package models

type RegisterModel struct {
	Email    string `json:"email" binding:"required,email"`
	Telegram string `json:"telegram" binding:"required"`
	Password string `json:"password" binding:"required,min=8"`
}

type LoginModel struct {
	Email    *string `json:"email"`
	Telegram *string `json:"telegram"`
	Password *string `json:"password" binding:"required"`
}
