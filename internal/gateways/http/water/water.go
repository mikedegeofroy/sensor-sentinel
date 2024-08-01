package water

import (
	"sensor-sentinel/internal/services"

	"github.com/gin-gonic/gin"
)

func SetupHandlers(r *gin.RouterGroup, services services.Services) {
	userGroup := r.Group("water")
	userGroup.GET("status", GetWaterStatus(services))
}
