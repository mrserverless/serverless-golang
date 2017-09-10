package api

type CreateFunc func(entity Entity) error
type GetFunc func(id [16]byte) (Entity, error)
type ListFunc func(params map[string]string) ([]Entity, error)
type UpdateFunc func(v Entity) error
type DeleteFunc func(id [16]byte) error

// CRUDAPI is a generic interface for common CRUD operations
// either extend this or use your own for more complex requirements
type CRUDAPI interface {
	Create (entity Entity) error
	Get(id [16]byte) (Entity, error)
	List(params map[string]string) ([]Entity, error)
	Update(v Entity) error
	Delete(id [16]byte) error
}