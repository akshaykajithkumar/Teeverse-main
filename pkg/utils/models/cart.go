package models

import "Teeverse/pkg/domain"

// type AddToCart struct {
// 	UserID      int `json:"user_id"`
// 	InventoryID int `json:"inventory_id"`
// }

type GetCart struct {
	ProductName string  `json:"product_name"`
	Category_id int     `json:"category_id"`
	Quantity    int     `json:"quantity"`
	Total       float64 `json:"total_price"`
}

type CheckOut struct {
	Addresses      []domain.Address
	Products       []GetCart
	PaymentMethods []domain.PaymentMethod
	TotalPrice     float64
}

type Order struct {
	AddressID int `json:"address_id"`
	PaymentID int `json:"paymentID"`
}
