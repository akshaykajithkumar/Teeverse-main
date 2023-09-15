package repository

import (
	"Teeverse/pkg/domain"
	interfaces "Teeverse/pkg/repository/interface"
	"Teeverse/pkg/utils/models"
	"errors"
	"strconv"

	"gorm.io/gorm"
)

type inventoryRepository struct {
	DB *gorm.DB
}

func NewInventoryRepository(DB *gorm.DB) interfaces.InventoryRepository {
	return &inventoryRepository{
		DB: DB,
	}
}

// func (i *inventoryRepository) AddInventory(inventory models.Inventory, url string) (models.InventoryResponse, error) {
// 	var inventoryResponse models.InventoryResponse
// 	query := `
//     INSERT INTO inventories (category_id,category product_name,description, stock, price, image)
//     VALUES (?, ?, ?, ?, ?,?) RETURNING id
// 	`
// 	i.DB.Raw(query, inventory.CategoryID, inventory.ProductName, inventory.Description, inventory.Stock, inventory.Price, url).Scan(&inventoryResponse.ProductID)

// 	return inventoryResponse, nil

// }

func (i *inventoryRepository) AddInventory(inventory models.Inventory, url string) (models.InventoryResponse, error) {

	query := `
    INSERT INTO inventories (category_id, product_name, description, stock, price, image)
    VALUES (?, ?, ?, ?, ?, ?);
    `
	i.DB.Exec(query, inventory.CategoryID, inventory.ProductName, inventory.Description, inventory.Stock, inventory.Price, url)

	var inventoryResponse models.InventoryResponse

	return inventoryResponse, nil

}

func (i *inventoryRepository) UpdateImage(invID int, url string) (models.Inventory, error) {
	// Check the database connection
	if i.DB == nil {
		return models.Inventory{}, errors.New("database connection is nil")
	}

	if err := i.DB.Exec("UPDATE inventories SET image = ? WHERE id= ?", url, invID).Error; err != nil {
		return models.Inventory{}, err
	}
	// Retrieve the update
	var inventory models.Inventory
	if err := i.DB.Raw("SELECT * FROM inventories WHERE id=?", invID).Scan(&inventory).Error; err != nil {
		return models.Inventory{}, err
	}

	return inventory, nil
}

func (i *inventoryRepository) CheckInventory(pid int) (bool, error) {
	var k int
	err := i.DB.Raw("SELECT COUNT(*) FROM inventories WHERE id=?", pid).Scan(&k).Error
	if err != nil {
		return false, err
	}

	if k == 0 {
		return false, err
	}

	return true, err
}

func (i *inventoryRepository) UpdateInventory(pid int, invData models.UpdateInventory) (models.Inventory, error) {

	// Check the database connection
	if i.DB == nil {
		return models.Inventory{}, errors.New("database connection is nil")
	}

	if invData.CategoryID != 0 {
		if err := i.DB.Exec("UPDATE inventories SET category_id = ? WHERE id= ?", invData.CategoryID, pid).Error; err != nil {
			return models.Inventory{}, err
		}
	}
	if invData.ProductName != "" && invData.ProductName != "string" {
		if err := i.DB.Exec("UPDATE inventories SET product_name = ? WHERE id= ?", invData.ProductName, pid).Error; err != nil {
			return models.Inventory{}, err
		}
	}
	if invData.Description != "" && invData.Description != "string" {
		if err := i.DB.Exec("UPDATE inventories SET description = ? WHERE id= ?", invData.Description, pid).Error; err != nil {
			return models.Inventory{}, err
		}
	}
	if invData.Stock != 0 {
		if err := i.DB.Exec("UPDATE inventories SET stock =  ? WHERE id= ?", invData.Stock, pid).Error; err != nil {
			return models.Inventory{}, err
		}
	}

	if invData.Price != 0 {
		if err := i.DB.Exec("UPDATE inventories SET price =  ? WHERE id= ?", invData.Price, pid).Error; err != nil {
			return models.Inventory{}, err
		}
	}
	// Retrieve the update
	var inventory models.Inventory
	if err := i.DB.Raw("SELECT * FROM inventories WHERE id=?", pid).Scan(&inventory).Error; err != nil {
		return models.Inventory{}, err
	}

	return inventory, nil
}

func (i *inventoryRepository) DeleteInventory(inventoryID string) error {
	id, err := strconv.Atoi(inventoryID)
	if err != nil {
		return errors.New("converting into integer not happened")
	}

	result := i.DB.Exec("DELETE FROM inventories WHERE id = ?", id)

	if result.RowsAffected < 1 {
		return errors.New("no records with that ID exist")
	}

	return nil
}

func (i *inventoryRepository) ShowIndividualProducts(id string) (models.InventoryList, error) {
	pid, error := strconv.Atoi(id)
	if error != nil {
		return models.InventoryList{}, errors.New("convertion not happened")
	}
	var product models.InventoryList
	err := i.DB.Raw(`
	SELECT inventories.id, inventories.product_name, inventories.description, inventories.stock, inventories.price, inventories.image, categories.category AS category FROM inventories JOIN categories ON inventories.category_id = categories.id
		WHERE
			inventories.id = ?
			`, pid).Scan(&product).Error
	//SELECT inventories.id, inventories.product_name, inventories.description, inventories.stock, inventories.price, inventories.image, categories.category AS category FROM inventories JOIN categories ON inventories.category_id = categories.id LIMIT ? OFFSET ?", limit, offset

	if err != nil {
		return models.InventoryList{}, errors.New("error retrieved record")
	}
	return product, nil

}

func (ad *inventoryRepository) ListProducts(page int, limit int) ([]models.InventoryList, error) {
	// pagination purpose -
	if page == 0 {
		page = 1
	}
	if limit == 0 {
		limit = 10
	}
	offset := (page - 1) * limit
	var productDetails []models.InventoryList

	if err := ad.DB.Raw("SELECT inventories.id, inventories.product_name, inventories.description, inventories.stock, inventories.price, inventories.image, categories.category AS category FROM inventories JOIN categories ON inventories.category_id = categories.id LIMIT ? OFFSET ?", limit, offset).Scan(&productDetails).Error; err != nil {
		//"SELECT inventories.product_name,cart_products.quantity,cart_products.total_price AS Total FROM cart_products JOIN inventories ON cart_products.inventory_id=inventories.id WHERE user_id=?"

		return []models.InventoryList{}, err
	}

	return productDetails, nil

}

func (i *inventoryRepository) CheckStock(pid int) (int, error) {
	var k int
	if err := i.DB.Raw("SELECT stock FROM inventories WHERE id=?", pid).Scan(&k).Error; err != nil {
		return 0, err
	}
	return k, nil
}

func (i *inventoryRepository) CheckPrice(pid int) (float64, error) {
	var k float64
	err := i.DB.Raw("SELECT price FROM inventories WHERE id=?", pid).Scan(&k).Error
	if err != nil {
		return 0, err
	}

	return k, nil
}

// func (ad *inventoryRepository) SearchProducts(key string, page, limit int) ([]models.Inventory, error) {
// 	if page == 0 {
// 		page = 1
// 	}
// 	if limit == 0 {
// 		limit = 10
// 	}
// 	offset := (page - 1) * limit
// 	var productDetails []models.Inventory

// 	query := `
// 		SELECT *
// 		FROM inventories
// 		WHERE product_name ILIKE '%' || ? || '%'
// 		OR description ILIKE '%' || ? || '%'
// 		limit ? offset ?
// 	`

// 	if err := ad.DB.Raw(query, key, key, limit, offset).Scan(&productDetails).Error; err != nil {
// 		return []models.Inventory{}, err
// 	}

//		return productDetails, nil
//	}
func (ad *inventoryRepository) SearchProducts(key string, page, limit int, sortBy string) ([]models.InventoryList, error) {
	if page == 0 {
		page = 1
	}
	if limit == 0 {
		limit = 10
	}
	offset := (page - 1) * limit
	var productDetails []models.InventoryList

	// Define the default sorting order (ASC)
	sortOrder := "ASC"

	// Validate and set the sorting order based on sortBy input
	if sortBy == "desc" {
		sortOrder = "DESC"
	} // Add more sorting options as needed

	// query := `
	// 	SELECT *
	// 	FROM inventories
	// 	WHERE product_name ILIKE '%' || ? || '%'
	// 	OR description ILIKE '%' || ? || '%'
	// 	ORDER BY product_name ` + sortOrder + `
	// 	LIMIT ? OFFSET ?
	// `
	query := `
	SELECT
		inventories.id,
		inventories.product_name,
		inventories.description,
		inventories.stock,
		inventories.price,
		inventories.image,
		categories.category AS category
	FROM
		inventories
	JOIN
		categories
	ON
		inventories.category_id = categories.id
	WHERE
	product_name ILIKE '%' || ? || '%'
	 	OR description ILIKE '%' || ? || '%'
	 	ORDER BY product_name ` + sortOrder + `
	 	LIMIT ? OFFSET ?
	 `
	if err := ad.DB.Raw(query, key, sortOrder, limit, offset).Scan(&productDetails).Error; err != nil {
		return []models.InventoryList{}, err
	}

	return productDetails, nil
}

func (ad *inventoryRepository) GetCategoryProducts(catID int, page, limit int) ([]models.InventoryList, error) {
	if page == 0 {
		page = 1
	}
	if limit == 0 {
		limit = 10
	}
	offset := (page - 1) * limit
	var productDetails []models.InventoryList

	query := `
	SELECT inventories.id, inventories.product_name, inventories.description, inventories.stock, inventories.price, inventories.image, categories.category AS category FROM inventories JOIN categories ON inventories.category_id = categories.id 
		WHERE inventories.category_id=?
		limit ? offset ?
	`

	if err := ad.DB.Raw(query, catID, limit, offset).Scan(&productDetails).Error; err != nil {
		//SELECT inventories.id, inventories.product_name, inventories.description, inventories.stock, inventories.price, inventories.image, categories.category AS category FROM inventories JOIN categories ON inventories.category_id = categories.id LIMIT ? OFFSET ?", limit, offset
		return []models.InventoryList{}, err
	}

	return productDetails, nil
}
func (ad *inventoryRepository) GetCategories(page, limit int) ([]domain.Category, error) {
	if page == 0 {
		page = 1
	}
	if limit == 0 {
		limit = 10
	}
	offset := (page - 1) * limit
	var categories []domain.Category

	if err := ad.DB.Raw("select id,category from categories limit ? offset ?", limit, offset).Scan(&categories).Error; err != nil {
		return []domain.Category{}, err
	}

	return categories, nil
}

func (ad *inventoryRepository) AddImage(product_id int, imageURL string) (models.InventoryResponse, error) {
	var inventoryResponse models.InventoryResponse
	query := `
    INSERT INTO images (inventory_id,image_url)
    VALUES (?, ?) RETURNING inventory_id
	`
	ad.DB.Raw(query, product_id, imageURL).Scan(&inventoryResponse.ProductID)

	return inventoryResponse, nil
}

func (ad *inventoryRepository) DeleteImage(product_id, imageID int) error {
	result := ad.DB.Exec("DELETE FROM images WHERE id = ?", imageID)

	if result.RowsAffected < 1 {
		return errors.New("no records with that ID exist")
	}

	return nil
}

func (ad *inventoryRepository) GetImagesFromInventoryID(product_id int) ([]models.ImageInfo, error) {
	var images []models.ImageInfo
	if err := ad.DB.Raw("select id,image_url from images where inventory_id=?", product_id).Scan(&images).Error; err != nil {
		return []models.ImageInfo{}, err
	}
	return images, nil
}
