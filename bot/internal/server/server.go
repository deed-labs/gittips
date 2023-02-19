package server

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"strconv"
	"time"

	"go.uber.org/zap"
)

const (
	readTimeout  = 30 * time.Second
	writeTimeout = 15 * time.Second
)

type Server struct {
	srv *http.Server
}

func New(handler http.Handler, port int) *Server {
	address := net.JoinHostPort("0.0.0.0", strconv.Itoa(port))

	srv := &http.Server{
		Addr:         address,
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
		Handler:      handler,
	}

	return &Server{
		srv: srv,
	}
}

func (s *Server) Run(ctx context.Context) error {
	go func() {
		err := s.srv.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			err = fmt.Errorf("listen and serve: %w", err)
			zap.Error(err)
		}
	}()

	<-ctx.Done()
	downCtx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	if err := s.srv.Shutdown(downCtx); err != nil {
		return fmt.Errorf("shutdown: %w", err)
	}

	return nil
}
