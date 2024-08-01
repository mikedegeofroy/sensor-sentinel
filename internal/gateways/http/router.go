package http

import (
	"sensor-sentinel/docs"
	"sensor-sentinel/internal/gateways/http/health"
	"sensor-sentinel/internal/gateways/http/water"
	"sensor-sentinel/internal/services"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func setupRouter(s *Server, services services.Services) {
	s.Router.HandleMethodNotAllowed = true
	s.Router.Use(allowOriginMiddleware())

	v1 := s.Router.Group("/api/v1")
	{
		health.SetupHandlers(v1)
		water.SetupHandlers(v1, services)
	}

	docs.SwaggerInfo.BasePath = "/api/v1"
	v1.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
