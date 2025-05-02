package api

import (
	"github.com/RicliZz/aidentity/internal/services"
	"github.com/gin-gonic/gin"
)

type ProfileHandler struct {
	ProfileService services.ProfileServiceInterface
}

func NewProfileHandler(ProfileService services.ProfileServiceInterface) *ProfileHandler {
	return &ProfileHandler{
		ProfileService: ProfileService,
	}
}

func (h *ProfileHandler) InitProfileHandlers(router *gin.RouterGroup) {
	profileRouter := router.Group("/profile")
	{
		profileRouter.POST("/new/:role", h.ProfileService.CreateProfile)
	}
}
