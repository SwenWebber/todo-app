package repository

import (
	"errors"
	"log"
	"time"

	"github.com/swenwebber/todo-app/internal/model"
)

type MemoryTaskRepo struct {
	tasks  map[int]model.Task
	nextID int
}

func NewMemoryTaskRepository() *MemoryTaskRepo {
	return &MemoryTaskRepo{
		tasks:  make(map[int]model.Task),
		nextID: 1,
	}
}

// Create method implementation
func (m *MemoryTaskRepo) Create(task model.Task) model.Task {

	task.ID = m.nextID
	task.CreatedAt = time.Now()
	task.UpdatedAt = time.Now()
	m.tasks[task.ID] = task

	m.nextID++

	return task

}

// GetById task method implementation
func (m *MemoryTaskRepo) GetById(id int) (model.Task, error) {

	task, exists := m.tasks[id]
	if !exists {
		return model.Task{}, errors.New("task not found")
	}
	return task, nil
}

// GetAll task  method implementation
func (m *MemoryTaskRepo) GetAll() ([]model.Task, error) {
	tasks := []model.Task{}

	for _, task := range m.tasks {
		tasks = append(tasks, task)
	}

	return tasks, nil
}

// Update task method implementation
func (m *MemoryTaskRepo) Update(task model.Task) (model.Task, error) {
	_, exists := m.tasks[task.ID]

	if !exists {
		return model.Task{}, errors.New("task not found")
	}
	originalTask := m.tasks[task.ID]
	task.CreatedAt = originalTask.CreatedAt
	task.UpdatedAt = time.Now()

	m.tasks[task.ID] = task
	return task, nil

}

// Delete task  method implementation
func (m *MemoryTaskRepo) Delete(id int) error {
	_, exists := m.tasks[id]

	if !exists {
		return errors.New("task not found")
	}

	delete(m.tasks, id)
	log.Printf("Task %d deleted", id) // check what task was deleted

	if len(m.tasks) == 0 {
		m.tasks = make(map[int]model.Task)
		m.nextID = 1
	}

	return nil
}
