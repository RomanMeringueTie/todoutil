package repository

import "todo/internal/model"

type TodoRepository interface {
	GetAll() []model.Todo
}
