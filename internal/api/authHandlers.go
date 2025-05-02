package api

import (
	"github.com/RicliZz/aidentity/internal/services"
	"github.com/gin-gonic/gin"
)

type AuthenticationHandler struct {
	AuthenticationService services.AuthenticationServiceInterface
}

func NewAuthenticationHandler(AuthenticationService services.AuthenticationServiceInterface) *AuthenticationHandler {
	return &AuthenticationHandler{
		AuthenticationService: AuthenticationService,
	}
}

func (h *AuthenticationHandler) InitAuthenticationHandlers(router *gin.RouterGroup) {
	authRouter := router.Group("/auth")
	{
		authRouter.POST("/register", h.AuthenticationService.Register)
		authRouter.POST("/login", h.AuthenticationService.Login)
	}
}
