package api

import (
	"github.com/RicliZz/aidentity/internal/services"
	"github.com/gin-gonic/gin"
)

type QualityHandlers struct {
	QualityService services.QualityServiceInterface
}

func NewQualityHandlers(QualityService services.QualityServiceInterface) *QualityHandlers {
	return &QualityHandlers{
		QualityService: QualityService,
	}
}

func (h *QualityHandlers) InitQualityHandlers(router *gin.RouterGroup) {
	qualityRouter := router.Group("/quality")
	{
		qualityRouter.POST("/create")
	}
}
