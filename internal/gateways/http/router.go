package http

import (
	"sensor-sentinel/docs"
	"sensor-sentinel/internal/gateways/http/health"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func setupRouter(s *Server) {
	s.Router.HandleMethodNotAllowed = true
	s.Router.Use(allowOriginMiddleware())

	v1 := s.Router.Group("/api/v1")
	{
		health.SetupHandlers(v1)
	}

	docs.SwaggerInfo.BasePath = "/api/v1"
	v1.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
