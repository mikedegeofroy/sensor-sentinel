package server

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tj/go-spin"
	"golang.org/x/sync/errgroup"

	"sensor-sentinel/internal/gateways/http"
)

const shutdownDuration = 1500 * time.Millisecond

type Server struct {
	HttpServer *http.Server
}

func NewServer() *Server {
	r := gin.Default()

	s := &Server{
		HttpServer: http.NewServer(r),
	}

	return s
}

func (s *Server) Run(ctx context.Context) error {
	eg := errgroup.Group{}

	eg.Go(func() error {
		return s.HttpServer.HttpServer.ListenAndServe()
	})

	<-ctx.Done()
	err := s.HttpServer.HttpServer.Shutdown(ctx)
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
