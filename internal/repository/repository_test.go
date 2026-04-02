package repository

import (
	"reflect"
	"testing"
	"todo/internal/model"
)

func TestTodoRepositoryImpl(t *testing.T) {
	var mockTodos []model.Todo = []model.Todo{
		*model.NewTodo("1", model.Open, "global"),
		*model.NewTodo("2", model.Closed, "internal/service.go"),
		*model.NewTodo("3", model.InProgress, "global"),
		*model.NewTodo("4", model.InProgress, "internal/repository.go"),
		*model.NewTodo("5", model.InProgress, "global"),
	}
	repository := TodoRepositoryImpl{todos: mockTodos}

	expected := mockTodos
	actual := repository.GetAll()

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("expected: %v, actual: %v", expected, actual)
	}
}
