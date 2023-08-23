package interfaces

import (
	"Teeverse/pkg/domain"
	"Teeverse/pkg/utils/models"
)

type PaymentUseCase interface {
	AddNewPaymentMethod(paymentMethod string) error
	RemovePaymentMethod(paymentMethodID int) error
	GetPaymentMethods() ([]domain.PaymentMethod, error)
	MakePaymentRazorPay(orderID string, userID int) (models.OrderPaymentDetails, error)
	VerifyPayment(paymentID string, razorID string, orderID string) error
}
