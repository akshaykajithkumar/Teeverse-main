package models

type InventoryResponse struct {
	ProductID int
	Stock     int
}

type StockUpdate struct {
	Productid int `json:"product_id"`
	Stock     int `json:"stock"`
}

type Inventory struct {
	ID          uint    `json:"id"`
	CategoryID  int     `json:"category_id"`
	Image       string  `json:"image"`
	ProductName string  `json:"productName"`
	Description string  `json:"description"`
	Stock       int     `json:"stock"`
	Price       float64 `json:"price"`
}

type NewInventory struct {
	CategoryID  int     `json:"category_id"`
	ProductName string  `json:"productName"`
	Description string  `json:"description"`
	Stock       int     `json:"stock"`
	Price       float64 `json:"price"`
}