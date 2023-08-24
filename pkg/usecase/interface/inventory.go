package interfaces

import (
	"Teeverse/pkg/utils/models"
)

type InventoryUseCase interface {
	AddInventory(inventory models.Inventory, image string) (models.InventoryResponse, error)
	UpdateInventory(invID int, invData models.UpdateInventory) (models.Inventory, error)
	DeleteInventory(id string) error

	ShowIndividualProducts(s string) (models.Inventory, error)
	ListProducts(page int, limit int) ([]models.Inventory, error)
	SearchProducts(key string, page, limit int) ([]models.Inventory, error)
	GetCategoryProducts(catID int, page, limit int) ([]models.Inventory, error)
}
