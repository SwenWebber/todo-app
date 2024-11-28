package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/swenwebber/todo-app/internal/handler"
	"github.com/swenwebber/todo-app/internal/server/routes"
)

type Config struct {
	Port         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	IdleTimeout  time.Duration
}

type Server struct {
	httpServer *http.Server
	handler    *handler.TaskHandler
	config     *Config
}

func NewServer(handler *handler.TaskHandler, config *Config) *Server {
	if config == nil {
		config = &Config{
			Port:         "8080",
			ReadTimeout:  15 * time.Second,
			WriteTimeout: 15 * time.Second,
			IdleTimeout:  60 * time.Second,
		}
	}

	router := routes.NewRouter(handler)

	httpServer := &http.Server{
		Addr:         ":" + config.Port,
		Handler:      router,
		ReadTimeout:  config.ReadTimeout,
		WriteTimeout: config.WriteTimeout,
		IdleTimeout:  config.IdleTimeout,
	}

	return &Server{
		httpServer: httpServer,
		handler:    handler,
		config:     config,
	}
}

func (s *Server) Run() error {
	//Creating error channel
	errChan := make(chan error, 1)

	//Creating shutdown signal

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	//starting server in go routine
	go func() {
		log.Printf("Starting server on port %s", s.config.Port)

		if err := s.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			errChan <- err
		}

	}()

	//Waiting for signals and errors

	select {
	case <-quit:
		//  got shutdown signal
		log.Println("Server is shutting down...")
		return s.shutdown()
	case err := <-errChan:
		// got error
		return fmt.Errorf("server error:%w", err)
	}

}

func (s *Server) shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	err := s.httpServer.Shutdown(ctx)

	if err != nil {
		log.Printf("Error during server shutdown: %v", err)
		return fmt.Errorf("server shutdown error:%w", err)
	}

	log.Println("Server shutdown completed")
	return nil
}
