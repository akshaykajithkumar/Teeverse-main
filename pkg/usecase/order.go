package usecase

import (
	domain "Teeverse/pkg/domain"
	interfaces "Teeverse/pkg/repository/interface"
	services "Teeverse/pkg/usecase/interface"
	"Teeverse/pkg/utils/models"
	"fmt"
)

type orderUseCase struct {
	orderRepository interfaces.OrderRepository
	userUseCase     services.UserUseCase
}

func NewOrderUseCase(repo interfaces.OrderRepository, userUseCase services.UserUseCase) *orderUseCase {
	return &orderUseCase{
		orderRepository: repo,
		userUseCase:     userUseCase,
	}
}

func (i *orderUseCase) GetOrders(id int) ([]domain.Order, error) {

	orders, err := i.orderRepository.GetOrders(id)
	if err != nil {
		return []domain.Order{}, err
	}

	return orders, nil

}

func (i *orderUseCase) OrderItemsFromCart(userid int, order models.Order) (string, error) {

	cart, err := i.userUseCase.GetCart(userid)
	if err != nil {
		return "", err
	}

	var total float64
	for _, v := range cart {
		total = total + v.Total
	}

	//COD
	if order.PaymentID == 1 {
		order_id, err := i.orderRepository.OrderItems(userid, order, total)
		if err != nil {
			return "", err
		}

		if err := i.orderRepository.AddOrderProducts(order_id, cart); err != nil {
			return "", err
		}
		cartID, _ := i.userUseCase.GetCartID(userid)
		if err := i.userUseCase.ClearCart(cartID); err != nil {
			return "", err
		}

	} else if order.PaymentID == 2 {
		// razorpay
		order_id, err := i.orderRepository.OrderItems(userid, order, total)
		if err != nil {
			return "", err
		}

		if err := i.orderRepository.AddOrderProducts(order_id, cart); err != nil {
			return "", err
		}
		link := fmt.Sprintf("http://localhost:1243/users/payment/razorpay?id=%d", order_id)

		return link, err
	}

	//wallet

	return "", nil

}

func (i *orderUseCase) CancelOrder(id int) error {

	err := i.orderRepository.CancelOrder(id)
	if err != nil {
		return err
	}
	return nil

}

func (i *orderUseCase) EditOrderStatus(status string, id int) error {

	err := i.orderRepository.EditOrderStatus(status, id)
	if err != nil {
		return err
	}
	return nil

}

func (i *orderUseCase) AdminOrders() (domain.AdminOrdersResponse, error) {

	var response domain.AdminOrdersResponse

	pending, err := i.orderRepository.AdminOrders("PENDING")
	if err != nil {
		return domain.AdminOrdersResponse{}, err
	}

	shipped, err := i.orderRepository.AdminOrders("SHIPPED")
	if err != nil {
		return domain.AdminOrdersResponse{}, err
	}

	delivered, err := i.orderRepository.AdminOrders("DELIVERED")
	if err != nil {
		return domain.AdminOrdersResponse{}, err
	}

	canceled, err := i.orderRepository.AdminOrders("CANCELED")
	if err != nil {
		return domain.AdminOrdersResponse{}, err
	}

	response.Canceled = canceled
	response.Pending = pending
	response.Shipped = shipped
	response.Delivered = delivered
	return response, nil

}
