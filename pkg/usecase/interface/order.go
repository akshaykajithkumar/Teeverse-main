package interfaces

import (
	"Teeverse/pkg/domain"
	"Teeverse/pkg/utils/models"
)

type OrderUseCase interface {
	GetOrders(id int) ([]domain.Order, error)
	OrderItemsFromCart(userid int, order models.Order) (string, error)
	CancelOrder(id int) error
	EditOrderStatus(status string, id int) error
	AdminOrders() (domain.AdminOrdersResponse, error)
}
