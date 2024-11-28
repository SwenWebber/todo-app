package main

import (
	"log"
	"net/http"

	"github.com/swenwebber/todo-app/internal/handler"
	"github.com/swenwebber/todo-app/internal/repository"
	"github.com/swenwebber/todo-app/internal/server/routes"
	"github.com/swenwebber/todo-app/internal/service"
)

func main() {
	repo := repository.NewMemoryTaskRepository()
	service := service.NewTaskService(repo)
	handler := handler.NewTaskHandler(service)

	router := routes.NewRouter(handler)

	log.Println("Server  starting on :8080")

	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal("Server error:", err)
	}

}
