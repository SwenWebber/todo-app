package repository

import "github.com/swenwebber/todo-app/internal/model"

type TaskRepository interface {
	Create(task model.Task) model.Task //Creating new task
	GetById(id int) model.Task         //Getting task by id(exact id)
	GetAll() []model.Task              //Getting all tasks
	Update(task model.Task) model.Task //Updating task
	Delete(id int)                     //Deleting concrete task
}
