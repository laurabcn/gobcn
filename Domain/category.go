package Domain

import (
	"github.com/satori/go.uuid"
)

type Category struct {
	Id  uuid.UUID
	Name string
	Language string
	Mostra bool
}