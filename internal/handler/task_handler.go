package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/swenwebber/todo-app/internal/model"
	"github.com/swenwebber/todo-app/internal/service"
)

type TaskHandler struct {
	service *service.TaskService
}

type CreateTaskRequest struct {
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

type UpdateTaskRequest struct {
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

type TaskResponse struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Status    bool   `json:"status"`
	CreatedAt string `json:"CreatedAt"`
	UpdatedAt string `json:"UpdatedAt"`
}

func NewTaskHandler(s *service.TaskService) *TaskHandler {
	return &TaskHandler{
		service: s,
	}
}

func (h *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	var req CreateTaskRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if req.Title == "" || strings.TrimSpace(req.Title) == "" {
		http.Error(w, "title cannot be empty", http.StatusBadRequest)
		return
	}

	task, err := h.service.CreateTask(req.Title, req.Status)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response := TaskResponse{
		ID:        task.ID,
		Title:     task.Title,
		Status:    task.Status,
		CreatedAt: task.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: task.UpdatedAt.Format("2006-01-02 15:04:05"),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)

}

func (h *TaskHandler) GetTask(w http.ResponseWriter, r *http.Request) {
	path := strings.Split(r.URL.Path, "/")
	id := path[len(path)-1]

	converted_id, err := strconv.Atoi(id)

	if err != nil {
		http.Error(w, "id must be a number", http.StatusBadRequest)
	}

	if converted_id <= 0 {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return

	}

	task, err := h.service.GetTask(converted_id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(TaskResponse{
		ID:        task.ID,
		Title:     task.Title,
		Status:    task.Status,
		CreatedAt: task.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: task.UpdatedAt.Format("2006-01-02 15:04:05"),
	})

}

func (h *TaskHandler) GetAllTasks(w http.ResponseWriter, r *http.Request) {
	allTasks, err := h.service.GetAllTasks()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	allTasksSlice := make([]TaskResponse, 0)

	for _, task := range allTasks {
		allTasksSlice = append(allTasksSlice, TaskResponse{
			ID:        task.ID,
			Title:     task.Title,
			Status:    task.Status,
			CreatedAt: task.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: task.UpdatedAt.Format("2006-01-02 15:04:05"),
		})

	}

	response := struct {
		Tasks []TaskResponse `json:"tasks"`
	}{
		Tasks: allTasksSlice,
	}

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(response)

}

func (h *TaskHandler) UpdateTask(w http.ResponseWriter, r *http.Request) {
	path := strings.Split(r.URL.Path, "/")
	id := path[len(path)-1]
	var req UpdateTaskRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	if req.Title == "" || strings.TrimSpace(req.Title) == "" {
		http.Error(w, "title cannot be empty", http.StatusBadRequest)
		return
	}

	convertedID, err := strconv.Atoi(id)
	if err != nil || convertedID <= 0 {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	task := model.Task{
		ID:     convertedID,
		Title:  req.Title,
		Status: req.Status,
	}

	updatedTask, err := h.service.UpdateTask(task)

	response := TaskResponse{
		ID:        updatedTask.ID,
		Title:     updatedTask.Title,
		Status:    updatedTask.Status,
		CreatedAt: updatedTask.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: updatedTask.UpdatedAt.Format("2006-01-02 15:04:05"),
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

func (h *TaskHandler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	path := strings.Split(r.URL.Path, "/")

	id := path[len(path)-1]

	convertedID, err := strconv.Atoi(id)

	if err != nil || convertedID <= 0 {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	newErr := h.service.DeleteTask(convertedID)

	if newErr != nil {
		http.Error(w, newErr.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
