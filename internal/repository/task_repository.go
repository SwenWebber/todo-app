package repository

import "github.com/swenwebber/todo-app/internal/model"

type TaskRepository interface {
	Create(task model.Task) model.Task          //Creating new task
	GetById(id int) (model.Task, error)         //Getting task by id(exact id)
	GetAll() ([]model.Task, error)              //Getting all tasks
	Update(task model.Task) (model.Task, error) //Updating task
	Delete(id int) error                        //Deleting concrete task
}
