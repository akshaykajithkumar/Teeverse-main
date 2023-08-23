package usecase

import (
	"Teeverse/pkg/domain"
	interfaces "Teeverse/pkg/repository/interface"
	services "Teeverse/pkg/usecase/interface"
	"Teeverse/pkg/utils/models"
	"errors"
	"strconv"

	"github.com/razorpay/razorpay-go"
)

type paymentUseCase struct {
	paymentRepository interfaces.PaymentRepository
	userRepository    interfaces.UserRepository
}

func NewPaymentUseCase(paymentRepository interfaces.PaymentRepository, userRepository interfaces.UserRepository) services.PaymentUseCase {
	return &paymentUseCase{
		paymentRepository: paymentRepository,
		userRepository:    userRepository,
	}
}

func (p *paymentUseCase) AddNewPaymentMethod(paymentMethod string) error {
	if paymentMethod == " " {
		return errors.New("enter method name")
	}
	if err := p.paymentRepository.AddNewPaymentMethod(paymentMethod); err != nil {
		return err
	}
	return nil
}

func (p *paymentUseCase) RemovePaymentMethod(paymentMethodID int) error {
	if paymentMethodID == 0 {
		return errors.New("enter method id")
	}
	if err := p.paymentRepository.RemovePaymentMethod(paymentMethodID); err != nil {
		return err
	}
	return nil
}

func (p *paymentUseCase) GetPaymentMethods() ([]domain.PaymentMethod, error) {
	paymentMethods, err := p.paymentRepository.GetPaymentMethods()
	if err != nil {
		return []domain.PaymentMethod{}, err
	}
	return paymentMethods, nil
}

func (p *paymentUseCase) MakePaymentRazorPay(orderID string, userID int) (models.OrderPaymentDetails, error) {
	var orderDetails models.OrderPaymentDetails
	//get orderid
	newid, err := strconv.Atoi(orderID)
	if err != nil {
		return models.OrderPaymentDetails{}, err
	}
	orderDetails.OrderID = newid

	orderDetails.UserID = userID

	//get username
	username, err := p.paymentRepository.FindUsername(userID)
	if err != nil {
		return models.OrderPaymentDetails{}, err
	}

	orderDetails.Username = username

	//get total
	newfinal, err := p.paymentRepository.FindPrice(newid)
	if err != nil {
		return models.OrderPaymentDetails{}, err
	}

	orderDetails.FinalPrice = newfinal

	client := razorpay.NewClient("rzp_test_hfgxceEmSupD3T", "9UrqCcAM6F5JMXqTFmJBQBGQ")

	data := map[string]interface{}{
		"amount":   int(orderDetails.FinalPrice) * 100,
		"currency": "INR",
		"receipt":  "some_receipt_id",
	}

	body, err := client.Order.Create(data, nil)
	if err != nil {
		return models.OrderPaymentDetails{}, nil
	}

	razorPayOrderID := body["id"].(string)

	orderDetails.Razor_id = razorPayOrderID

	return orderDetails, nil
}

func (p *paymentUseCase) VerifyPayment(paymentID string, razorID string, orderID string) error {

	err := p.paymentRepository.UpdatePaymentDetails(orderID, paymentID, razorID)
	if err != nil {
		return err
	}

	//clearcart

	orderIDint, err := strconv.Atoi(orderID)
	//fmt.Println("====orderID", orderID)
	if err != nil {
		return err
	}
	userID, err := p.userRepository.FindUserIDByOrderID(orderIDint)
	if err != nil {
		return err
	}
	cartID, err := p.userRepository.GetCartID(userID)
	//fmt.Println("CartID=======", cartID)

	if err != nil {
		return err
	}
	p.userRepository.ClearCart(cartID)
	if err != nil {
		return err
	}

	return nil

}
