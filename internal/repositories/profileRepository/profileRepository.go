package profileRepository

import (
	"context"
	"github.com/RicliZz/aidentity/internal/models"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ProfileRepository struct {
	db *pgxpool.Pool
}

func NewProfileRepository(db *pgxpool.Pool) *ProfileRepository {
	return &ProfileRepository{
		db: db,
	}
}

func (r *ProfileRepository) CreateStudentProfile(userID uuid.UUID, newStudent models.StudentModel) error {
	var universityID uuid.UUID
	var specialityID uuid.UUID
	tx, err := r.db.Begin(context.Background())
	if err != nil {
		return err
	}
	defer tx.Rollback(context.Background())
	err = tx.QueryRow(context.Background(),
		`INSERT INTO university (name) VALUES ($1) ON CONFLICT DO NOTHING RETURNING "ID"`,
		newStudent.University).Scan(&universityID)
	if err != nil {
		return err
	}
	err = tx.QueryRow(context.Background(),
		`INSERT INTO speciality ("universityID", name, code) VALUES ($1, $2, $3) RETURNING "ID"`,
		universityID, newStudent.Speciality, newStudent.SpecialityCode).Scan(&specialityID)
	if err != nil {
		return err
	}
	_, err = tx.Exec(context.Background(),
		`INSERT INTO student ("userID", "studyYear", "specialityID") VALUES ($1, $2, $3)`,
		userID, newStudent.StudyYear, specialityID)
	insertProfession := func(profs []models.ProfessionModel, relation string) error {
		for _, prof := range profs {
			_, err := tx.Exec(context.Background(),
				`INSERT INTO user_profession ("userID", "professionID", preference) VALUES ($1, $2, $3)`,
				userID, prof.ID, relation)
			if err != nil {
				return err
			}
		}
		return nil
	}

	if err = insertProfession(newStudent.LikedProfession, "liked"); err != nil {
		return err
	}
	if err = insertProfession(newStudent.NotLikedProfession, "disliked"); err != nil {
		return err
	}
	if err = insertProfession(newStudent.ParentProfession, "parent"); err != nil {
		return err
	}
	if err = insertProfession(newStudent.DreamProfession, "dream"); err != nil {
		return err
	}
	err = tx.Commit(context.Background())
	if err != nil {
		return err
	}
	return nil
}
