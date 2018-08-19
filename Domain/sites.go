package Domain

import (
	"github.com/satory/go.uuid"
)

type Site struct {
	Id  uuid.UUID
	Name string
	Language string
	Mostra bool
	District string
	Phone string
	Web string
	Content string
	Excerpt string
	Latitude string
	Longitude string
	Type string
	Barri string
	Address string
	Position string
}

type SiteRepository interface {
	AddSite(site *Site) error
}