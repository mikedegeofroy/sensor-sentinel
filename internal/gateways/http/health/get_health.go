package health

import (
	"github.com/gin-gonic/gin"
)

// GetHealth godoc
// @Summary Get health
// @Description Get the health of the application
// @Tags health
// @Schemes http https
// @Success 200
// @Failure 500
// @Router /health [get]
func GetHealth() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Status(200)
	}
}
