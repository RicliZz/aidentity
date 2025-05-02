package repositories

import (
	"github.com/RicliZz/aidentity/internal/models"
	"github.com/google/uuid"
)

type QualityRepositoryInterface interface {
	CreateQuality(nameNewQuality string) error
	DeleteQuality(qualityUUID uuid.UUID) (*models.Quality, error)
}

type AuthenticationRepositoryInterface interface {
	CreateUser(newUser models.RegisterModel) (*models.User, error)
	GetUserByEmail(email string) (uuid.UUID, string, error)
	GetUserByTelegram(telegram string) (uuid.UUID, string, error)
	CreateSession(model models.CreateSessionModel) error
}

type ProfileRepositoryInterface interface {
	CreateStudentProfile(userID uuid.UUID, newStudent models.StudentModel) error
}
