package interfaces

import "Teeverse/pkg/domain"

type CategoryRepository interface {
	AddCategory(category string) (domain.Category, error)
	CheckCategory(currrent string) (bool, error)
	UpdateCategory(current, new string) (domain.Category, error)
	DeleteCategory(categoryID string) error
	GetCategories(page, limit int) ([]domain.Category, error)
}
