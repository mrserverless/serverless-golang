package todo

import "github.com/satori/go.uuid"

type Todo struct {
	ID      uuid.UUID `json:"id"`
	Message string    `json:"message"`
}
