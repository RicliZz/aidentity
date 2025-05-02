package authService

import (
	"fmt"
	"github.com/RicliZz/aidentity/internal/models"
	"github.com/RicliZz/aidentity/pkg/passw"
	"github.com/gin-gonic/gin"
)

func (s *AuthenticationService) Login(c *gin.Context) {
	var loginRequest models.LoginModel
	var hashedPass string
	if err := c.ShouldBind(&loginRequest); err != nil {
		c.JSON(400, models.ErrorModel{"Ошибка в запросе"})
		return
	}
	if loginRequest.Telegram == nil && loginRequest.Email != nil {
		hashPass, err := s.AuthenticationRepository.GetUserByEmail(*loginRequest.Email)
		if err != nil {
			c.JSON(400, models.ErrorModel{fmt.Sprintf(
				"Не удалось найти пользоваля с Email %s", *loginRequest.Email)})
			return
		}
		hashedPass = hashPass
	} else {
		tg := *loginRequest.Telegram
		if tg[0] != '@' {
			tg = "@" + *loginRequest.Telegram
		}
		*loginRequest.Telegram = tg
		hashPass, err := s.AuthenticationRepository.GetUserByTelegram(*loginRequest.Email)
		if err != nil {
			c.JSON(400, models.ErrorModel{fmt.Sprintf(
				"Не удалось найти пользоваля с Telegram %s", *loginRequest.Telegram)})
			return
		}
		hashedPass = hashPass
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
	accessToken, refreshToken, err :=
}
