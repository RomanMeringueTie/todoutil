package service

import (
	"reflect"
	"testing"
	"todo/internal/model"
)

type TodoRepositoryMock struct {
	todos []model.Todo
}

func (repository *TodoRepositoryMock) GetAll() []model.Todo {
	return repository.todos
}

var mockTodos []model.Todo = []model.Todo{
	*model.NewTodo("1", model.Open, "global"),
	*model.NewTodo("2", model.Closed, "internal/service.go"),
	*model.NewTodo("3", model.InProgress, "global"),
	*model.NewTodo("4", model.InProgress, "internal/repository.go"),
	*model.NewTodo("5", model.InProgress, "global"),
}

var mockRepository = TodoRepositoryMock{
	todos: mockTodos,
}

var testableService TodoService = NewTodoServiceImpl(&mockRepository)

func TestTodoServiceImplGetAll(t *testing.T) {
	expected := mockTodos

	actual := testableService.GetAll()

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("expected: %v, actual: %v", expected, actual)
	}
}

func TestTodoServiceImplGetInProgress(t *testing.T) {
	expected := []model.Todo{*model.NewTodo("3", model.InProgress, "global"),
		*model.NewTodo("4", model.InProgress, "internal/repository.go"),
		*model.NewTodo("5", model.InProgress, "global"),
	}

	actual := testableService.GetInProgress()

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("expected: %v, actual: %v", expected, actual)
	}
}
