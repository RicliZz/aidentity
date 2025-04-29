package models

type CreateQualityModel struct {
	Name string `json:"name" binding:"required"`
}
