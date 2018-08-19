package Domain

import (
	"github.com/satori/go.uuid"
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

type Category struct {
	Id  uuid.UUID
	Name string
	Language string
	Mostra bool
}

type CategoryRepository interface {
	Add(category *Category) error
}

type SiteRepository interface {
	AddSite(site *Site) error
}

type SiteCatRepository interface {
	AddSiteCategory(site *Site, category *Category) error
}