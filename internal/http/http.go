package http

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/fenek-dev/llm-craft/internal/entity"
)

type Server struct {
	s   Service
	mux *http.ServeMux
}

type Service interface {
	Generate(ctx context.Context, el1, el2 string) (entity.Element, bool, error)
}

func New(s Service) *Server {
	return &Server{
		s:   s,
		mux: http.NewServeMux(),
	}
}

func (s *Server) InitRoutes() {
	s.mux.HandleFunc("GET /pair", s.HandleGenerate)
}

func (s *Server) Run(ctx context.Context, addr string) error {
	srv := &http.Server{
		Addr:    addr,
		Handler: s.mux,
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	go func() {
		log.Println("Starting server on", addr)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Could not listen on %s: %v\n", addr, err)
		}
	}()

	<-stop
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		return err
	}

	log.Println("Server gracefully stopped")
	return nil
}
