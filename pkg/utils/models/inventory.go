package models

type InventoryResponse struct {
	ProductID int
	Stock     int
}

type InventoryUpdate struct {
	Productid int `json:"product_id"`
	Stock     int `json:"stock"`
}

type Inventory struct {
	ID          uint    `json:"id"`
	CategoryID  int     `json:"category_id"`
	Image       string  `json:"image"`
	ProductName string  `json:"product_name"`
	Stock       int     `json:"stock"`
	Price       float64 `json:"price"`
}
