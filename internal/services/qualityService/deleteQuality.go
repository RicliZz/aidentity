package qualityService

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (r *QualityService) DeleteQuality(c *gin.Context) {
	qualityID := c.Param("id")
	qualityUUID, err := uuid.Parse(qualityID)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	deletedQuality, err := r.QualityRepository.DeleteQuality(qualityUUID)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, deletedQuality)
}
