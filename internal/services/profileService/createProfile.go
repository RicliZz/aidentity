package profileService

import (
	"github.com/RicliZz/aidentity/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (s *ProfileService) CreateProfile(c *gin.Context) {
	role := c.Param("role")
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(500, models.ErrorModel{"Проблема с JWT"})
		return
	}
	uuidUserID := userID.(uuid.UUID)

	if role == "student" {
		var newStudent models.StudentModel
		if err := c.ShouldBind(&newStudent); err != nil {
			c.JSON(400, models.ErrorModel{"Ошибка в запросе"})
			return
		}

		err := s.ProfileRepository.CreateStudentProfile(uuidUserID, newStudent)
		if err != nil {
			c.JSON(400, models.ErrorModel{err.Error()})
			return
		}
	} else if role == "schoolboy" {
	}
}
