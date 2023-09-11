package usecase

import (
	"Teeverse/pkg/domain"
	interfaces "Teeverse/pkg/repository/interface"
	services "Teeverse/pkg/usecase/interface"
	"Teeverse/pkg/utils/models"
	"errors"
	"strconv"
)

type inventoryUseCase struct {
	repository interfaces.InventoryRepository
}

func NewInventoryUseCase(repo interfaces.InventoryRepository) services.InventoryUseCase {
	return &inventoryUseCase{
		repository: repo,
	}
}

func (i *inventoryUseCase) AddInventory(inventory models.Inventory, image string) (models.InventoryResponse, error) {

	//send the url and save it in database
	InventoryResponse, err := i.repository.AddInventory(inventory, image)
	if err != nil {
		return models.InventoryResponse{}, err
	}

	return InventoryResponse, nil

}

func (i *inventoryUseCase) UpdateInventory(invID int, invData models.UpdateInventory) (models.Inventory, error) {

	result, err := i.repository.CheckInventory(invID)
	if err != nil {

		return models.Inventory{}, err
	}

	if !result {
		return models.Inventory{}, errors.New("there is no inventory as you mentioned")
	}

	newinv, err := i.repository.UpdateInventory(invID, invData)
	if err != nil {
		return models.Inventory{}, err
	}

	return newinv, err
}

func (i *inventoryUseCase) DeleteInventory(inventoryID string) error {

	err := i.repository.DeleteInventory(inventoryID)
	if err != nil {
		return err
	}
	return nil

}

func (i *inventoryUseCase) ShowIndividualProducts(id string) (models.InventoryDetails, error) {

	product, err := i.repository.ShowIndividualProducts(id)
	if err != nil {
		return models.InventoryDetails{}, err
	}
	productID, err := strconv.Atoi(id)
	if err != nil {
		return models.InventoryDetails{}, err
	}
	var AdditionalImages []models.ImageInfo
	AdditionalImages, err = i.repository.GetImagesFromInventoryID(productID)
	if err != nil {
		return models.InventoryDetails{}, err
	}
	InvDetails := models.InventoryDetails{Inventory: product, AdditionalImages: AdditionalImages}

	return InvDetails, nil

}

func (i *inventoryUseCase) ListProducts(page int, limit int) ([]models.InventoryList, error) {

	productDetails, err := i.repository.ListProducts(page, limit)
	if err != nil {
		return []models.InventoryList{}, err
	}
	return productDetails, nil

}

func (i *inventoryUseCase) SearchProducts(key string, page, limit int, sortBY string) ([]models.InventoryList, error) {

	productDetails, err := i.repository.SearchProducts(key, page, limit, sortBY)
	if err != nil {
		return []models.InventoryList{}, err
	}

	return productDetails, nil

}

func (i *inventoryUseCase) GetCategoryProducts(catID int, page, limit int) ([]models.InventoryList, error) {

	productDetails, err := i.repository.GetCategoryProducts(catID, page, limit)
	if err != nil {
		return []models.InventoryList{}, err
	}

	return productDetails, nil

}
func (i *inventoryUseCase) GetCategories(page, limit int) ([]domain.Category, error) {
	categories, err := i.repository.GetCategories(page, limit)
	if err != nil {
		return []domain.Category{}, err
	}
	return categories, nil
}

func (i *inventoryUseCase) AddImage(product_id int, imageURL string) (models.InventoryResponse, error) {

	InventoryResponse, err := i.repository.AddImage(product_id, imageURL)
	if err != nil {
		return models.InventoryResponse{}, err
	}

	return InventoryResponse, nil

}

func (i *inventoryUseCase) DeleteImage(product_id int, image_id int) error {

	err := i.repository.DeleteImage(product_id, image_id)
	if err != nil {
		return err
	}

	return nil

}
