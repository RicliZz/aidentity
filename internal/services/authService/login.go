package authService

import (
	"fmt"
	"github.com/RicliZz/aidentity/internal/models"
	"github.com/RicliZz/aidentity/pkg/JWT"
	"github.com/RicliZz/aidentity/pkg/passw"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (s *AuthenticationService) Login(c *gin.Context) {
	var loginRequest models.LoginModel
	var hashedPass string
	var userID uuid.UUID
	if err := c.ShouldBind(&loginRequest); err != nil {
		c.JSON(400, models.ErrorModel{"Ошибка в запросе"})
		return
	}
	if loginRequest.Telegram == nil && loginRequest.Email != nil {
		id, hashPass, err := s.AuthenticationRepository.GetUserByEmail(*loginRequest.Email)
		if err != nil {
			c.JSON(400, models.ErrorModel{fmt.Sprintf(
				"Не удалось найти пользоваля с Email %s", *loginRequest.Email)})
			return
		}
		hashedPass = hashPass
		userID = id
	} else {
		tg := *loginRequest.Telegram
		if tg[0] != '@' {
			tg = "@" + *loginRequest.Telegram
		}
		*loginRequest.Telegram = tg
		id, hashPass, err := s.AuthenticationRepository.GetUserByTelegram(*loginRequest.Email)
		if err != nil {
			c.JSON(400, models.ErrorModel{fmt.Sprintf(
				"Не удалось найти пользоваля с Telegram %s", *loginRequest.Telegram)})
			return
		}
		hashedPass = hashPass
		userID = id
	}
	match, err := passw.CompareHashAndPassword(hashedPass, *loginRequest.Password)
	if err != nil {
		c.JSON(400, models.ErrorModel{err.Error()})
		return
	}
	if !match {
		c.JSON(401, models.ErrorModel{"Пароли не совпадают"})
		return
	}
	accessToken, refreshToken := JWT.CreateTokens(userID)
	newSession := models.CreateSessionModel{
		UserID:       userID,
		RefreshToken: refreshToken,
		Ua:           c.GetHeader("User-Agent"),
		Fingerprint:  loginRequest.Fingerprint,
		IP:           c.ClientIP(),
		ExpiresAt:    30 * 24 * 60 * 60,
	}
	err = s.AuthenticationRepository.CreateSession(newSession)
	if err != nil {
		c.JSON(500, models.ErrorModel{err.Error()})
		return
	}
	c.JSON(200, models.TokensModel{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	})
}
