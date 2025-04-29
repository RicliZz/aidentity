package services

import "github.com/gin-gonic/gin"

type QualityServiceInterface interface {
	CreateQuality(c *gin.Context) error
}
