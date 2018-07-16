package repository

import "awesomeProject2/Domain"

type CategoryRepository interface {
	Add(category *Domain.Category) error
}