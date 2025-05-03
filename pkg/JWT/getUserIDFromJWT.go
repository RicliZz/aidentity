package JWT

import (
	"github.com/RicliZz/aidentity/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"strings"
)

func GetUserID(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			c.AbortWithStatusJSON(401, models.ErrorModel{Message: "Invalid Header"})
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return []byte(secret), nil
		})

		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(401, models.ErrorModel{Message: "Invalid token"})
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.AbortWithStatusJSON(401, models.ErrorModel{Message: "Invalid token"})
			return
		}
		userID, ok := claims["userID"].(string)
		if !ok {
			c.AbortWithStatusJSON(401, models.ErrorModel{Message: "Unauthorized"})
			return
		}
		c.Set("userID", userID)
		c.Next()
	}
}
