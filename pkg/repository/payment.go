package repository

import (
	"Teeverse/pkg/domain"
	interfaces "Teeverse/pkg/repository/interface"

	"gorm.io/gorm"
)

type paymentRepository struct {
	DB *gorm.DB
}

func NewPaymentRepository(DB *gorm.DB) interfaces.PaymentRepository {
	return &paymentRepository{
		DB: DB,
	}
}

func (p *paymentRepository) AddNewPaymentMethod(paymentMethod string) error {
	query := `INSERT INTO payment_methods(payment_method) VALUES(?)`
	if err := p.DB.Exec(query, paymentMethod); err != nil {
		return err.Error
	}

	return nil
}

func (p *paymentRepository) RemovePaymentMethod(paymentMethodID int) error {
	query := `DELETE FROM payment_methods WHERE id=?`
	if err := p.DB.Exec(query, paymentMethodID); err != nil {
		return err.Error
	}

	return nil
}

func (p *paymentRepository) GetPaymentMethods() ([]domain.PaymentMethod, error) {
	var paymentMethods []domain.PaymentMethod
	if err := p.DB.Raw(`SELECT * FROM payment_methods`).Scan(&paymentMethods).Error; err != nil {
		return []domain.PaymentMethod{}, err
	}

	return paymentMethods, nil

}

func (p *paymentRepository) FindUsername(user_id int) (string, error) {
	var name string
	if err := p.DB.Raw("SELECT name FROM users WHERE id=?", user_id).Scan(&name).Error; err != nil {
		return "", err
	}

	return name, nil
}

func (p *paymentRepository) FindPrice(order_id int) (float64, error) {
	var price float64
	if err := p.DB.Raw("SELECT price FROM orders WHERE id=?", order_id).Scan(&price).Error; err != nil {
		return 0, err
	}

	return price, nil
}

func (p *paymentRepository) UpdatePaymentDetails(orderID, paymentID, razorID string) error {
	status := "PAID"
	if err := p.DB.Exec(`UPDATE orders SET payment_status = $1 , payment_id=$3 WHERE id = $2`, status, orderID, paymentID).Error; err != nil {
		return err
	}
	return nil
}
