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
	Phone int
	Web string
	Content string
	Excerpt string
	Latitude string
	Longitude string
	Type string
	Barri string
	Address string
}

type SiteRepository interface {
	AddSite(site *Site) error
}