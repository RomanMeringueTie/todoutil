package di

import (
	"todo/internal/repository"
	"todo/internal/service"
)

type TodoDiContainer interface {
	GetRepository() repository.TodoRepository
	GetService() service.TodoService
}

type TodoDiContainerImpl struct {
	repository repository.TodoRepository
	service    service.TodoService
}

func NewTodoDiContainerImpl() *TodoDiContainerImpl {
	repository := repository.NewTodoRepositoryImpl(".", ".go")
	service := service.NewTodoServiceImpl(repository)
	return &TodoDiContainerImpl{repository: repository, service: service}
}

func (container TodoDiContainerImpl) GetRepository() repository.TodoRepository {
	return container.repository
}

func (container TodoDiContainerImpl) GetService() service.TodoService {
	return container.service
}
