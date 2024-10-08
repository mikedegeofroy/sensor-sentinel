package http

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"sensor-sentinel/cmd/app/config"
	"sensor-sentinel/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/tj/go-spin"

	"golang.org/x/sync/errgroup"
)

const shutdownDuration = 1500 * time.Millisecond

type Server struct {
	HttpServer http.Server
	Router     *gin.Engine
}

func NewServer(router *gin.Engine, services services.Services) *Server {
	s := &Server{
		Router: router,
		HttpServer: http.Server{
			Addr:    fmt.Sprintf(":%d", config.C.Server.Port),
			Handler: router,
		},
	}

	setupRouter(s, services)

	return s
}

func (s *Server) Run(ctx context.Context) error {
	eg := errgroup.Group{}

	eg.Go(func() error {
		return s.HttpServer.ListenAndServe()
	})

	<-ctx.Done()
	err := s.HttpServer.Shutdown(ctx)
	err = errors.Join(eg.Wait(), err)
	shutdownWait()
	return err
}

func shutdownWait() {
	spinner := spin.New()
	const spinIterations = 20
	for i := 0; i < spinIterations; i++ {
		fmt.Printf("\rgraceful shutdown %s ", spinner.Next())
		time.Sleep(shutdownDuration / spinIterations)
	}
}
