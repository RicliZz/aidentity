package models

import "github.com/google/uuid"

type StudentModel struct {
	University         string            `json:"university" binding:"required"`
	StudyYear          string            `json:"study_year" binding:"required"`
	Speciality         string            `json:"speciality" binding:"required"`
	SpecialityCode     string            `json:"speciality_code" binding:"required"`
	LikedProfession    []ProfessionModel `json:"liked_profession" binding:"required"`
	NotLikedProfession []ProfessionModel `json:"not_liked_profession" binding:"required"`
	ParentProfession   []ProfessionModel `json:"parent_profession" binding:"required"`
	DreamProfession    []ProfessionModel `json:"dream_profession" binding:"required"`
}

type ProfessionModel struct {
	ID   uuid.UUID `json:"id" binding:"required"`
	Name string    `json:"name"`
}
