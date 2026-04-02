package main

import (
	"todo/internal/di"
	"todo/internal/model"
	"todo/internal/presentation"
	"todo/internal/repository"
)

func main() {
	inputFlag := presentation.ParseFlags()
	if inputFlag == presentation.Init {
		repository.CreateConfig()
		return
	}

	diContainer := di.NewTodoDiContainerImpl()
	service := diContainer.GetService()

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
