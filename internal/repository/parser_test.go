package repository

import (
	"reflect"
	"testing"
	"todo/internal/model"
)

func TestHasTodoSuffix(t *testing.T) {
	var hasSuffix = hasTodoSuffix("parser.go", ".go")
	if !hasSuffix {
		t.Errorf("expected: %t, actual: %t", true, hasSuffix)
	}

	hasSuffix = hasTodoSuffix("parser.todo", ".go")
	if !hasSuffix {
		t.Errorf("expected: %t, actual: %t", true, hasSuffix)
	}

	hasSuffix = hasTodoSuffix("main.c", ".go")
	if hasSuffix {
		t.Errorf("expected: %t, actual: %t", false, hasSuffix)
	}
}

func TestParseLine(t *testing.T) {
	var expected = model.NewTodo("Add Test for parseLine", model.InProgress, "parser.go")
	var actual = parseLine("parser.go", "// TASK: Add Test for parseLine [*]", "TASK: ")
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("expected: %v, actual: %v", expected, actual)
	}

	expected = model.NewTodo("Add Test for parseLine", model.Open, "parser.go")
	actual = parseLine("parser.go", "// TASK: Add Test for parseLine [o]", "TASK: ")
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("expected: %v, actual: %v", expected, actual)
	}

	expected = model.NewTodo("Add Test for parseLine", model.Closed, "parser.go")
	actual = parseLine("parser.go", "// TASK: Add Test for parseLine [x]", "TASK: ")
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("expected: %v, actual: %v", expected, actual)
	}

	expected = model.NewTodo("Add Test for parseLine", model.Open, "parser.go")
	actual = parseLine("parser.go", "// TASK: Add Test for parseLine", "TASK: ")
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("expected: %v, actual: %v", expected, actual)
	}

	expected = nil
	actual = parseLine("parser.go", "func getTodoNameAndStatus(line string) (string, model.Status) {", "TASK: ")
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("expected: %v, actual: %v", expected, actual)
	}
}
