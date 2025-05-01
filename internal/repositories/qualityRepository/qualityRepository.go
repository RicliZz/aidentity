package qualityRepository

import (
	"context"
	"github.com/RicliZz/aidentity/internal/models"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
)

type QualityRepository struct {
	db *pgxpool.Pool
}

func NewQualityRepository(db *pgxpool.Pool) *QualityRepository {
	return &QualityRepository{
		db: db,
	}
}

func (r *QualityRepository) CreateQuality(nameNewQuality string) error {
	sqlQuery := "INSERT INTO quality (\"nameQuality\") VALUES ($1)"
	_, err := r.db.Exec(context.Background(), sqlQuery, nameNewQuality)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func (r *QualityRepository) DeleteQuality(qualityUUID uuid.UUID) (*models.Quality, error) {
	var deletedQuality models.Quality
	sqlQuery := "DELETE FROM quality WHERE \"ID\" = $1 RETURNING \"ID\", \"nameQuality\" "
	err := r.db.QueryRow(context.Background(), sqlQuery, qualityUUID).Scan(&deletedQuality.ID, &deletedQuality.Name)
	if err != nil {
		return nil, err
	}
	return &deletedQuality, nil
}
