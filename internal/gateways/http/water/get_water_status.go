package water

import "github.com/gin-gonic/gin"

// GetWaterStatus godoc
// @Summary Get water level
// @Description Get the status of the water level
// @Tags water
// @Schemes http https
// @Success 200
// @Failure 500
// @Router /water/status [get]
func GetWaterStatus() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Status(200)
	}
}
