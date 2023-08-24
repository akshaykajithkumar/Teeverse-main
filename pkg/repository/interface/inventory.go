package interfaces

import (
	"Teeverse/pkg/utils/models"
)

type InventoryRepository interface {
	AddInventory(inventory models.Inventory, url string) (models.InventoryResponse, error)
	CheckInventory(pid int) (bool, error)
	UpdateInventory(pid int, invData models.UpdateInventory) (models.Inventory, error)
	DeleteInventory(id string) error
	ShowIndividualProducts(id string) (models.Inventory, error)
	ListProducts(page int, limit int) ([]models.Inventory, error)
	CheckStock(inventory_id int) (int, error)
	CheckPrice(inventory_id int) (float64, error)
	SearchProducts(key string, page, limit int) ([]models.Inventory, error)
	GetCategoryProducts(catID int, page, limit int) ([]models.Inventory, error)
}
