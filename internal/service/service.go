package service

import (
	"todo/internal/model"
	"todo/internal/repository"
)

type TodoService interface {
	GetAll() []model.Todo
	GetInProgress() []model.Todo
	GetOpen() []model.Todo
	GetClosed() []model.Todo
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
	inProgressTodos := service.getTodoWithStatus(model.InProgress)
	return inProgressTodos
}

func (service *TodoServiceImpl) GetOpen() []model.Todo {
	openTodos := service.getTodoWithStatus(model.Open)
	return openTodos
}

func (service *TodoServiceImpl) GetClosed() []model.Todo {
	closedTodos := service.getTodoWithStatus(model.Closed)
	return closedTodos
}

func (service *TodoServiceImpl) getTodoWithStatus(status model.Status) []model.Todo {
	allTodos := service.repository.GetAll()

	todosWithStatus := make([]model.Todo, 0)
	for _, todo := range allTodos {
		if todo.GetStatus() == status {
			todosWithStatus = append(todosWithStatus, todo)
		}
	}

	return todosWithStatus
}
