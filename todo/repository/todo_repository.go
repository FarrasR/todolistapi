package repository

import (
	"todolistapi/domain"
	"todolistapi/entity/model"

	"gorm.io/gorm"
)

type todoRepository struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) domain.TodoRepository {
	return &todoRepository{
		db: db,
	}
}

func (r *todoRepository) GetSingleTodo(TodoID uint) (model.Todo, error) {
	var todo model.Todo

	result := r.db.Where(model.Todo{TodoID: TodoID}).First(&todo)

	if result.Error != nil {
		return model.Todo{}, result.Error
	}
	return todo, nil
}

func (r *todoRepository) GetAllTodos() ([]model.Todo, error) {
	var todos []model.Todo

	result := r.db.Find(&todos)

	if result.Error != nil {
		return []model.Todo{}, result.Error
	}
	return todos, nil
}

func (r *todoRepository) GetAllTodosByActivityId(ActivityID uint) ([]model.Todo, error) {
	var todos []model.Todo

	result := r.db.Where("activity_group_id = ?", ActivityID).Find(&todos)

	if result.Error != nil {
		return []model.Todo{}, result.Error
	}
	return todos, nil
}

func (r *todoRepository) CreateTodo(todo model.Todo) (model.Todo, error) {

	result := r.db.Create(&todo)

	if result.Error != nil {
		return model.Todo{}, result.Error
	}

	return todo, nil
}

func (r *todoRepository) DeleteTodo(todoID uint) error {
	result := r.db.Delete(&model.Todo{}, todoID)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *todoRepository) UpdateTodo(todo model.Todo) (model.Todo, error) {
	result := r.db.Save(&todo)
	if result.Error != nil {
		return model.Todo{}, result.Error
	}
	return todo, nil
}
