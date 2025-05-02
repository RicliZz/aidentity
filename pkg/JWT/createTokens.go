package JWT

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"os"
	"time"
)

func CreateTokens(userID uuid.UUID) (string, uuid.UUID) {
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"userID": userID,
		"exp":    time.Now().Add(time.Minute * 15).Unix(),
	})
	tokenString, err := accessToken.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		panic(err)
	}

	refreshToken := uuid.New()

	return tokenString, refreshToken
}
