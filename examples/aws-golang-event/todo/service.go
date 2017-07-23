package todo

import (
	"github.com/satori/go.uuid"
)

func Create(t *Todo) error {
	return nil
}

func Get(id uuid.UUID) *Todo {
	t := &Todo{id, "hello world"}
	return t
}

func List() []*Todo {
	t := &Todo{uuid.NewV4(), "hello world"}
	todos := make([]*Todo, 1)
	todos[0] = t
	return todos
}

func Update(t *Todo) error {
	return nil
}

func Delete(id uuid.UUID) error {
	return nil
}