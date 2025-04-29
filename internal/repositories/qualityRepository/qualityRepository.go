package qualityRepository

import (
	"context"
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
