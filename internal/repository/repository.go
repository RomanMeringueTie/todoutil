package repository

import (
	"todo/internal/model"
)

type TodoRepository interface {
	GetAll() []model.Todo
}

type TodoRepositoryImpl struct {
	todos []model.Todo
}

func NewTodoRepositoryImpl() *TodoRepositoryImpl {
	todos := parseDirs()

	return &TodoRepositoryImpl{todos: todos}
}

func (repository *TodoRepositoryImpl) GetAll() []model.Todo {
	return repository.todos
}
