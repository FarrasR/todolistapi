package domain

import "todolistapi/entity/model"

type TodoRepository interface {
	GetSingleTodo(TodoID uint) (model.Todo, error)
	GetAllTodos() ([]model.Todo, error)
	GetAllTodosByActivityId(ActivityID uint) ([]model.Todo, error)
	CreateTodo(todo model.Todo) (model.Todo, error)
	DeleteTodo(todoID uint) error
	UpdateTodo(todo model.Todo) (model.Todo, error)
}
