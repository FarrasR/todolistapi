package domain

import "todolistapi/entity/model"

type TodoUsecase interface {
	GetAllTodos() ([]model.Todo, error)
	GetAllTodosByActivityId(ActivityID uint) ([]model.Todo, error)
	GetOneTodo(TodoID uint) (model.Todo, error)
	CreateTodo(ActivityID uint, title string) (model.Todo, error)
	DeleteTodo(TodoID uint) error
	UpdateTodo(todoID uint, title string) (model.Todo, error)
}
