package water

import (
	"github.com/gin-gonic/gin"
)

func SetupHandlers(r *gin.RouterGroup) {
	userGroup := r.Group("water")
	userGroup.GET("status", GetWaterStatus())
}
