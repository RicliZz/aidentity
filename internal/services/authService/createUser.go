package authService

import (
	"github.com/RicliZz/aidentity/internal/models"
	"github.com/RicliZz/aidentity/pkg/passw"
	"github.com/gin-gonic/gin"
)

func (s *AuthenticationService) Register(c *gin.Context) {
	var newUser models.RegisterModel
	if err := c.ShouldBind(&newUser); err != nil {
		c.JSON(400, models.ErrorModel{"Ошибка в запросе"})
		return
	}
	hashPass, err := passw.CreateHash(newUser.Password)
	if err != nil {
		c.JSON(500, models.ErrorModel{err.Error()})
		return
	}
	newUser.Password = hashPass

	if newUser.Telegram[0] != '@' {
		tg := "@" + newUser.Telegram
		newUser.Telegram = tg
	}
	createdUser, err := s.AuthenticationRepository.CreateUser(newUser)
	if err != nil {
		c.JSON(500, models.ErrorModel{err.Error()})
		return
	}
	c.JSON(201, createdUser)
}
