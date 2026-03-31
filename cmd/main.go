package main

import (
	"todo/internal/di"
)

func main() {
	// TODO: Read flags
	diContainer := di.NewTodoDiContainerImpl()
	service := diContainer.GetService()
	todos := service.GetAll()
	// todos := service.GetInProgress()
	displayTodos(todos)
}
