package authRepository

import (
	"context"
	"github.com/RicliZz/aidentity/internal/models"
	"github.com/google/uuid"
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

func (r *AuthenticationRepository) GetUserByEmail(email string) (uuid.UUID, string, error) {
	var password string
	var userID uuid.UUID
	sqlQuery := `SELECT "ID",password FROM "user" WHERE email = $1`
	err := r.db.QueryRow(context.Background(), sqlQuery, email).Scan(&userID, &password)
	if err != nil {
		return uuid.Nil, "", err
	}
	return userID, password, nil
}

func (r *AuthenticationRepository) GetUserByTelegram(telegram string) (uuid.UUID, string, error) {
	var password string
	var userID uuid.UUID
	sqlQuery := `SELECT "ID", password FROM "user" WHERE telegram = $1`
	err := r.db.QueryRow(context.Background(), sqlQuery, telegram).Scan(&userID, &password)
	if err != nil {
		return uuid.Nil, "", err
	}
	return userID, password, nil
}

func (r *AuthenticationRepository) CreateSession(newSession models.CreateSessionModel) error {
	sqlQuery := `INSERT INTO "refreshSession" ("userID", "refreshToken", "ua", fingerprint, "IP", "exp")
				VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := r.db.Exec(context.Background(), sqlQuery,
		newSession.UserID, newSession.RefreshToken, newSession.Ua, newSession.Fingerprint, newSession.IP, newSession.ExpiresAt)
	if err != nil {
		return err
	}
	return nil
}
