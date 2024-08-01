package water

import (
	"net/http"
	"sensor-sentinel/internal/services"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Level int `json:"level"`
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
			Level: level,
		}
		c.JSON(http.StatusOK, res)
	}
}
