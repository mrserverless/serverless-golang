package todo

import (
	"github.com/satori/go.uuid"
	"fmt"
)

type TodoService interface {
	Create(t *Todo) error
	Get(id uuid.UUID) (*Todo, error)
	List() ([]*Todo, error)
	Update(t *Todo) error
	Delete(id uuid.UUID) error
}

type MockTodoService struct {
	todoMap map[uuid.UUID]*Todo
}

func NewMockTodoService() *MockTodoService {
	return &MockTodoService{todoMap: make(map[uuid.UUID]*Todo)}
}

func (s *MockTodoService) Create(t *Todo) error {
	s.todoMap[t.ID] = t
	return nil
}

func (s *MockTodoService) Get(id uuid.UUID) (*Todo, error) {
	t := &Todo{id, "hello world"}
	return t, nil
}

func (s *MockTodoService) List() ([]*Todo, error) {
	ts := make([]*Todo, len(s.todoMap))
	i := 0
	for _, t := range s.todoMap {
		ts[i] = t
		i++
	}
	return ts, nil
}

func (s *MockTodoService) Update(t *Todo) error {
	if s.todoMap[t.ID] == nil {
		return fmt.Errorf("Not found: %s", t.ID)
	}
	s.todoMap[t.ID] = t
	return nil
}

func (s *MockTodoService) Delete(id uuid.UUID) error {
	if s.todoMap[id] == nil {
		return fmt.Errorf("Not found: %s", id)
	}
	delete(s.todoMap, id)
	return nil
}