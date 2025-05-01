package models

import "github.com/google/uuid"

type Quality struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type CreateQualityModel struct {
	Name string `json:"name" binding:"required"`
}
