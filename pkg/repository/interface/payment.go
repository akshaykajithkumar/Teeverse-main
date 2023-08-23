package interfaces

import "Teeverse/pkg/domain"

type PaymentRepository interface {
	AddNewPaymentMethod(paymentMethod string) error
	RemovePaymentMethod(paymentMethodID int) error
	GetPaymentMethods() ([]domain.PaymentMethod, error)
	FindUsername(user_id int) (string, error)
	FindPrice(order_id int) (float64, error)
	UpdatePaymentDetails(orderID, paymentID, razorID string) error
}
