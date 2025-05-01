package qualityService

import (
	"github.com/RicliZz/aidentity/internal/models"
	"github.com/gin-gonic/gin"
)

func (s *QualityService) CreateQuality(c *gin.Context) {
	var newQuality models.CreateQualityModel
	if err := c.ShouldBindJSON(&newQuality); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	err := s.QualityRepository.CreateQuality(newQuality.Name)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(201, "Successfully created quality")
}
