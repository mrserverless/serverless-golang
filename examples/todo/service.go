package todo

import (
	"github.com/satori/go.uuid"
	"fmt"
	"github.com/yunspace/serverless-golang/api"
)

type MockTodoService struct {
	todoMap map[uuid.UUID]*Todo
}

var _ api.CRUDAPI = (*MockTodoService)(nil)

func NewMockTodoService() *MockTodoService {
	return &MockTodoService{todoMap: make(map[uuid.UUID]*Todo)}
}

func (s *MockTodoService) Create(v api.Entity) error {
	t := v.(*Todo)
	s.todoMap[t.ID()] = t
	return nil
}

func (s *MockTodoService) Get(id [16]byte) (api.Entity, error) {
	var t api.Entity
	t = &Todo{id, "hello world"}
	return t, nil
}

func (s *MockTodoService) List(_ map[string]string) ([]api.Entity, error) {
	ts := make([]api.Entity, len(s.todoMap))
	i := 0
	for _, t := range s.todoMap {
		ts[i] = t
		i++
	}
	return ts, nil
}

func (s *MockTodoService) Update(t api.Entity) error {
	if s.todoMap[t.ID()] == nil {
		return fmt.Errorf("Not found: %s", t.ID())
	}
	s.todoMap[t.ID()] = t.(*Todo)
	return nil
}

func (s *MockTodoService) Delete(id [16]byte) error {
	if s.todoMap[id] == nil {
		return fmt.Errorf("Not found: %s", id)
	}
	delete(s.todoMap, id)
	return nil
}