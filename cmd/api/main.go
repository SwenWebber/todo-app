package main

import (
	"log"
	"time"

	"github.com/swenwebber/todo-app/internal/handler"
	"github.com/swenwebber/todo-app/internal/repository"
	"github.com/swenwebber/todo-app/internal/server"

	"github.com/swenwebber/todo-app/internal/service"
)

func main() {
	repo := repository.NewMemoryTaskRepository()
	service := service.NewTaskService(repo)
	api_handler := handler.NewTaskHandler(service)
	templateHandler := handler.NewTemplateHandler(service)

	config := &server.Config{
		Port:         "8080",
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}

	srv := server.NewServer(api_handler, templateHandler, config)

	if err := srv.Run(); err != nil {
		log.Fatalf("Server error: %v", err)
	}

}
