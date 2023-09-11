package interfaces

import (
	"Teeverse/pkg/domain"
	"Teeverse/pkg/utils/models"
)

type InventoryUseCase interface {
	AddInventory(inventory models.Inventory, image string) (models.InventoryResponse, error)
	UpdateInventory(invID int, invData models.UpdateInventory) (models.Inventory, error)
	DeleteInventory(id string) error

	ShowIndividualProducts(s string) (models.InventoryDetails, error)
	ListProducts(page int, limit int) ([]models.InventoryList, error)
	SearchProducts(key string, page, limit int, sortBY string) ([]models.InventoryList, error)
	GetCategories(page, limit int) ([]domain.Category, error)
	GetCategoryProducts(catID int, page, limit int) ([]models.InventoryList, error)

	AddImage(product_id int, imageURL string) (models.InventoryResponse, error)
	DeleteImage(product_id, image_id int) error
}
