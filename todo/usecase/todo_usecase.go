package usecase

import (
	"fmt"
	"todolistapi/domain"
	"todolistapi/entity/model"

	"gorm.io/gorm"
)

type todoUsecase struct {
	TodoRepository domain.TodoRepository
}

func NewTodoUsecase(todoRepository domain.TodoRepository) domain.TodoUsecase {
	return &todoUsecase{
		TodoRepository: todoRepository,
	}
}

func (u *todoUsecase) GetAllTodos() ([]model.Todo, error) {
	return u.TodoRepository.GetAllTodos()
}

func (u *todoUsecase) GetAllTodosByActivityId(ActivityID uint) ([]model.Todo, error) {
	return u.TodoRepository.GetAllTodosByActivityId(ActivityID)
}

func (u *todoUsecase) GetOneTodo(TodoID uint) (model.Todo, error) {
	return u.TodoRepository.GetSingleTodo(TodoID)
}

// the api is weird, so this is the best i could come up with without having any assumption
func (u *todoUsecase) CreateTodo(ActivityID uint, title string) (model.Todo, error) {
	todo := model.Todo{
		ActivityGroupID: ActivityID,
		Title:           title,
		Priority:        "very-high",
		IsActive:        false,
	}

	return u.TodoRepository.CreateTodo(todo)
}

func (u *todoUsecase) DeleteTodo(TodoID uint) error {
	_, err := u.TodoRepository.GetSingleTodo(TodoID)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("not found")
		}
		return err
	}

	err = u.TodoRepository.DeleteTodo(TodoID)
	if err != nil {
		return err
	}

	return nil
}

func (u *todoUsecase) UpdateTodo(todoID uint, title string) (model.Todo, error) {
	todo, err := u.TodoRepository.GetSingleTodo(todoID)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return model.Todo{}, fmt.Errorf("not found")
		}
		return model.Todo{}, err
	}
	todo.Title = title

	return u.TodoRepository.UpdateTodo(todo)
}
