package main

import (
	"fmt"
	"os"
	"text/tabwriter"
	"todo/internal/model"
)

func displayTodos(todos []model.Todo) {
	writer := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)
	fmt.Fprintln(writer, "Context\tName\tStatus\t")
	fmt.Fprintln(writer, "-------\t----\t------\t")

	for _, todo := range todos {
		displayTodo(writer, todo)
	}
	writer.Flush()
}

func displayTodo(writer *tabwriter.Writer, todo model.Todo) {
	fmt.Fprintf(writer, "%s\t%s\t%s\t\n", todo.GetContext(), todo.GetName(), todo.GetStatus().String())
}
