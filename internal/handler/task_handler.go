package handler

import "github.com/swenwebber/todo-app/internal/service"

type TaskHandler struct {
	service *service.TaskService
}
