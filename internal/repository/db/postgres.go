package db

import (
	"errors"

	"github.com/swenwebber/todo-app/internal/model"
	"gorm.io/gorm"
)

type DBTaskRepo struct {
	db *gorm.DB
}

func NewDBTaskRepo(db *gorm.DB) (*DBTaskRepo, error) {
	if db == nil {
		return nil, errors.New("Database connection is nil")
	}
	return &DBTaskRepo{db: db}, nil
}

func (db *DBTaskRepo) Create(task model.Task) model.Task {
	result := db.db.Create(&task)
	if result.Error != nil {
		return model.Task{}
	}

	return task
}

func (db *DBTaskRepo) GetById(id int) (model.Task, error) {
	var task model.Task
	result := db.db.First(&task, id)

	if result.Error != nil {
		return model.Task{}, result.Error
	}
	return task, nil

}

func (db *DBTaskRepo) GetAll() ([]model.Task, error) {
	var tasks []model.Task

	result := db.db.Find(&tasks)

	if result.Error != nil {
		return nil, result.Error
	}

	return tasks, nil

}

func (db *DBTaskRepo) Update(task model.Task) (model.Task, error) {
	var existingTask model.Task

	result := db.db.First(&existingTask, task.ID)

	if result.Error != nil {
		return model.Task{}, result.Error
	}

	result = db.db.Model(&existingTask).
		Updates(model.Task{Title: task.Title, Status: task.Status})

	if result.Error != nil {
		return model.Task{}, result.Error
	}

	var updatedTask model.Task

	if err := db.db.First(&updatedTask, task.ID).Error; err != nil {
		return model.Task{}, err
	}

	return updatedTask, nil
}

func (db *DBTaskRepo) Delete(id int) error {

	result := db.db.Delete(&model.Task{}, id)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("task not found")
	}

	if result := db.db.Find(&model.Task{}); result.RowsAffected == 0 {
		db.db.Exec("ALTER SEQUENCE tasks_id_seq RESTART WITH 1")

	}
	return nil
}
