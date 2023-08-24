package interfaces

import (
	"Teeverse/pkg/domain"
	"Teeverse/pkg/utils/models"
)

type OrderUseCase interface {
	GetOrders(id, page, limit int) ([]domain.Order, error)
	OrderItemsFromCart(userid int, order models.Order) (string, error)
	CancelOrder(id int) error
	EditOrderStatus(status string, id int) error
	AdminOrders(page, limit int) (domain.AdminOrdersResponse, error)
	DailyOrders() (domain.SalesReport, error)
	WeeklyOrders() (domain.SalesReport, error)
	MonthlyOrders() (domain.SalesReport, error)
	AnnualOrders() (domain.SalesReport, error)
	CustomDateOrders(dates models.CustomDates) (domain.SalesReport, error)
	ReturnOrder(id int) error
}
