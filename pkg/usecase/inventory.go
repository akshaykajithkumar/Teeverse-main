package usecase

import (
	interfaces "Teeverse/pkg/repository/interface"
	services "Teeverse/pkg/usecase/interface"
	"Teeverse/pkg/utils/models"
	"errors"
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

func (i *inventoryUseCase) UpdateInventory(pid int, stock int) (models.InventoryResponse, error) {

	result, err := i.repository.CheckInventory(pid)
	if err != nil {

		return models.InventoryResponse{}, err
	}

	if !result {
		return models.InventoryResponse{}, errors.New("there is no inventory as you mentioned")
	}

	newcat, err := i.repository.UpdateInventory(pid, stock)
	if err != nil {
		return models.InventoryResponse{}, err
	}

	return newcat, err
}

func (i *inventoryUseCase) DeleteInventory(inventoryID string) error {

	err := i.repository.DeleteInventory(inventoryID)
	if err != nil {
		return err
	}
	return nil

}

func (i *inventoryUseCase) ShowIndividualProducts(id string) (models.Inventory, error) {

	product, err := i.repository.ShowIndividualProducts(id)
	if err != nil {
		return models.Inventory{}, err
	}

	return product, nil

}

func (i *inventoryUseCase) ListProducts(page int, limit int) ([]models.Inventory, error) {

	productDetails, err := i.repository.ListProducts(page, limit)
	if err != nil {
		return []models.Inventory{}, err
	}
	return productDetails, nil

}

func (i *inventoryUseCase) SearchProducts(key string, page, limit int) ([]models.Inventory, error) {

	productDetails, err := i.repository.SearchProducts(key, page, limit)
	if err != nil {
		return []models.Inventory{}, err
	}

	return productDetails, nil

}
