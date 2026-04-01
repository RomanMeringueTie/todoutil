package service

import (
	"todo/internal/model"
	"todo/internal/repository"
)

type TodoService interface {
	GetAll() []model.Todo
	GetInProgress() []model.Todo
	// Add GetOpen
}

type TodoServiceImpl struct {
	repository repository.TodoRepository
}

func NewTodoServiceImpl(repository repository.TodoRepository) *TodoServiceImpl {
	return &TodoServiceImpl{repository: repository}
}

func (service *TodoServiceImpl) GetAll() []model.Todo {
	return service.repository.GetAll()
}

func (service *TodoServiceImpl) GetInProgress() []model.Todo {
	allTodos := service.repository.GetAll()

	inProgressTodos := make([]model.Todo, 0)
	for _, todo := range allTodos {
		if todo.GetStatus().String() == "InProgress" {
			inProgressTodos = append(inProgressTodos, todo)
		}
	}

	return inProgressTodos
}
