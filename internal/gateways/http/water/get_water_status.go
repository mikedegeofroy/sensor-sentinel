package water

import (
	"sensor-sentinel/internal/services"

	"github.com/gin-gonic/gin"
)

type Response struct {
	level int
}

// GetWaterStatus godoc
// @Summary Get water level
// @Description Get the status of the water level
// @Tags water
// @Schemes http https
// @Success 200
// @Failure 500
// @Router /water/status [get]
func GetWaterStatus(services services.Services) gin.HandlerFunc {
	return func(c *gin.Context) {
		level, err := services.WaterService.GetWaterLevel()
		if err != nil {
			c.Status(500)
		}
		res := Response{
			level: level,
		}
		c.JSON(200, res)
	}
}
