package api

// Entity represents a basic domain model with UUID
type Entity interface {
	ID() [16]byte
	SetID([16]byte)
}