package Domain

import (
	"github.com/satori/go.uuid"
)

type Category struct {
	Id  uuid.UUID
	Name string
	Language string
	Mostra bool
	Position string
}

type CategoryRepository interface {
	Add(category *Category) error
}