package services

import "github.com/gin-gonic/gin"

type QualityServiceInterface interface {
	CreateQuality(c *gin.Context)
	DeleteQuality(c *gin.Context)
}

type AuthenticationServiceInterface interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
}

type ProfileServiceInterface interface {
	CreateProfile(c *gin.Context)
}
