package transactions

import "github.com/google/uuid"

type Repository interface {
	Save(entity *Entity) (err error)
	FindById(id uuid.UUID) (entity Entity, err error)
}
