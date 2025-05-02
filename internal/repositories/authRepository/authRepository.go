package authRepository

import (
	"context"
	"github.com/RicliZz/aidentity/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type AuthenticationRepository struct {
	db *pgxpool.Pool
}

func NewAuthenticationRepository(db *pgxpool.Pool) *AuthenticationRepository {
	return &AuthenticationRepository{
		db: db,
	}
}

func (r *AuthenticationRepository) CreateUser(registerRequest models.RegisterModel) (*models.User, error) {
	var newUser models.User
	sqlQuery := `INSERT INTO "user" (email, telegram, password) VALUES 
                                                         ($1, $2, $3) RETURNING "ID", email, telegram`
	err := r.db.QueryRow(context.Background(), sqlQuery,
		registerRequest.Email, registerRequest.Telegram, registerRequest.Password).Scan(&newUser.ID, &newUser.Email, &newUser.Telegram)
	if err != nil {
		return nil, err
	}
	return &newUser, nil
}

func (r *AuthenticationRepository) GetUserByEmail(email string) (string, error) {
	var password string
	sqlQuery := `SELECT password FROM "user" WHERE email = $1`
	err := r.db.QueryRow(context.Background(), sqlQuery, email).Scan(&password)
	if err != nil {
		return "", err
	}
	return password, nil
}

func (r *AuthenticationRepository) GetUserByTelegram(telegram string) (string, error) {
	var password string
	sqlQuery := `SELECT password FROM "user" WHERE telegram = $1`
	err := r.db.QueryRow(context.Background(), sqlQuery, telegram).Scan(&password)
	if err != nil {
		return "", err
	}
	return password, nil
}
