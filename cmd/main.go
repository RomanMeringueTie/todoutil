package main

import (
	"todo/internal/di"
	"todo/internal/model"
	"todo/internal/presentation"
)

func main() {
	diContainer := di.NewTodoDiContainerImpl()
	service := diContainer.GetService()
	inputFlag := presentation.ParseFlags()
	var todos []model.Todo

	switch inputFlag {
	case presentation.All:
		todos = service.GetAll()
	case presentation.InProgress:
		todos = service.GetInProgress()
	case presentation.Open:
		todos = service.GetOpen()
	case presentation.Closed:
		todos = service.GetClosed()
	}

	presentation.DisplayTodos(todos)
}
