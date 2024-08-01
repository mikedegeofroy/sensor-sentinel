package health

import (
	"github.com/gin-gonic/gin"
)

func SetupHandlers(r *gin.RouterGroup) {
	userGroup := r.Group("health")
	userGroup.GET("", GetHealth())
}
