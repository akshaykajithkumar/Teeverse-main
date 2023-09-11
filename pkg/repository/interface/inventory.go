package interfaces

import (
	"Teeverse/pkg/domain"
	"Teeverse/pkg/utils/models"
)

type InventoryRepository interface {
	AddInventory(inventory models.Inventory, url string) (models.InventoryResponse, error)
	UpdateImage(invID int, url string) (models.Inventory, error)
	CheckInventory(pid int) (bool, error)
	UpdateInventory(pid int, invData models.UpdateInventory) (models.Inventory, error)
	DeleteInventory(id string) error
	ShowIndividualProducts(id string) (models.InventoryList, error)
	ListProducts(page int, limit int) ([]models.InventoryList, error)
	CheckStock(inventory_id int) (int, error)
	CheckPrice(inventory_id int) (float64, error)
	GetCategories(page, limit int) ([]domain.Category, error)
	SearchProducts(key string, page, limit int, sortBY string) ([]models.InventoryList, error)
	GetCategoryProducts(catID int, page, limit int) ([]models.InventoryList, error)
	AddImage(product_id int, imageURL string) (models.InventoryResponse, error)
	DeleteImage(product_id, imageID int) error
	GetImagesFromInventoryID(product_id int) ([]models.ImageInfo, error)
}
