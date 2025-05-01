package repositories

import (
	"github.com/RicliZz/aidentity/internal/models"
	"github.com/google/uuid"
)

type QualityRepositoryInterface interface {
	CreateQuality(nameNewQuality string) error
	DeleteQuality(qualityUUID uuid.UUID) (*models.Quality, error)
}
