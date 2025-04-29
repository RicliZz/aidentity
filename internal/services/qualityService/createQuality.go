package qualityService

import (
	"github.com/RicliZz/aidentity/internal/models"
	"github.com/gin-gonic/gin"
)

func (s *QualityService) CreateQuality(c *gin.Context) error {
	var newQuality models.CreateQualityModel
	if err := c.ShouldBindJSON(&newQuality); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return err
	}
	err := s.QualityRepository.CreateQuality(newQuality.Name)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return err
	}
	return nil
}
