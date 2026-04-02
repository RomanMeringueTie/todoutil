package model

type Todo struct {
	name    string
	status  Status
	context string
	// TODO: Add prioroty
}

func NewTodo(name string, status Status, context string) *Todo {
	return &Todo{name, status, context}
}

func (todo *Todo) GetName() string {
	return todo.name
}

func (todo *Todo) GetStatus() Status {
	return todo.status
}

func (todo *Todo) GetContext() string {
	return todo.context
}
