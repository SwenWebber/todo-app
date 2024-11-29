package routes

import (
	"github.com/gorilla/mux"
	"github.com/swenwebber/todo-app/internal/handler"
)

func NewRouter(h *handler.TaskHandler, th *handler.TemplateHandler) *mux.Router {
	r := mux.NewRouter()

	//frontend
	r.HandleFunc("/", th.Home).Methods("GET")

	//API endpoints
	r.HandleFunc("/tasks", h.CreateTask).Methods("POST")
	r.HandleFunc("/tasks/{id}", h.GetTask).Methods("GET")
	r.HandleFunc("/tasks", h.GetAllTasks).Methods("GET")
	r.HandleFunc("/tasks/{id}", h.UpdateTask).Methods("PUT")
	r.HandleFunc("/tasks/{id}", h.DeleteTask).Methods("DELETE")

	return r
}
