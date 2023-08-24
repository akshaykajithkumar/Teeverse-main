package interfaces

import (
	"Teeverse/pkg/domain"
	"Teeverse/pkg/utils/models"
	"time"
)

type OrderRepository interface {
	GetOrders(id, page, limit int) ([]domain.Order, error)
	GetProductsQuantity() ([]domain.ProductReport, error)
	GetOrdersInRange(startDate, endDate time.Time) ([]domain.Order, error)
	GetProductNameFromID(id int) (string, error)
	GetCart(userid int) ([]models.GetCart, error)
	OrderItems(userid int, order models.Order, total float64) (int, error)
	AddOrderProducts(order_id int, cart []models.GetCart) error
	CancelOrder(id int) error
	EditOrderStatus(status string, id int) error
	AdminOrders(status string) ([]domain.OrderDetails, error)

	CheckOrder(orderID string, userID int) error
	GetOrderDetail(orderID string) (domain.Order, error)
	FindUserIdFromOrderID(id int) (int, error)
	FindAmountFromOrderID(id int) (float64, error)
	ReturnOrder(id int) error
	CheckIfTheOrderIsAlreadyReturned(id int) (string, error)
}
