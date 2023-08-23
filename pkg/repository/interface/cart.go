package interfaces

import (
	"Teeverse/pkg/domain"
)

type CartRepository interface {
	//GetCart(id int) ([]models.GetCart, error)
	GetAddresses(id int) ([]domain.Address, error)
	GetCartId(user_id int) (int, error)
	CreateNewCart(user_id int) (int, error)
	AddLineItems(cart_id, inventory_id int) error
}
