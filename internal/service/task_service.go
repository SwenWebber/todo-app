package service

import (
	"errors"

	"github.com/swenwebber/todo-app/internal/model"
	"github.com/swenwebber/todo-app/internal/repository"
)

type TaskService struct {
	repo repository.TaskRepository
}

func NewTaskService(repo repository.TaskRepository) *TaskService {
	return &TaskService{
		repo: repo,
	}
}

func (s *TaskService) CreateTask(title string, status bool) (model.Task, error) {
	if title == "" {
		return model.Task{}, errors.New("Title cannot be empty")
	}

	if len(title) < 5 || len(title) > 55 {
		return model.Task{}, errors.New("Invalid title length")
	}

	task := model.Task{
		Title:  title,
		Status: status,
	}

	return s.repo.Create(task), nil
}

func (s *TaskService) GetTask(id int) (model.Task, error) {
	if id <= 0 {
		return model.Task{}, errors.New("id should be positive number and greater than 0")
	}

	task, err := s.repo.GetById(id)

	if err != nil {
		return model.Task{}, err
	}
	return task, nil
}

func (s *TaskService) GetAllTasks() ([]model.Task, error) {
	return s.repo.GetAll()
}

func (s *TaskService) UpdateTask(task model.Task) (model.Task, error) {
	if task.ID <= 0 {
		return model.Task{}, errors.New("Invalid id")
	}

	if task.Title == "" {
		return model.Task{}, errors.New("title cannot be empty")
	}

	updatedTask, err := s.repo.Update(task)

	if err != nil {
		return model.Task{}, err
	}

	return updatedTask, nil

}

func (s *TaskService) DeleteTask(id int) error {
	if id <= 0 {
		return errors.New("invalid id")
	}

	err := s.repo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
