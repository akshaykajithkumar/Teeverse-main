package repository

import (
	"Teeverse/pkg/domain"
	interfaces "Teeverse/pkg/repository/interface"
	"Teeverse/pkg/utils/models"
	"errors"
	"time"

	"gorm.io/gorm"
)

type orderRepository struct {
	DB *gorm.DB
}

func NewOrderRepository(db *gorm.DB) interfaces.OrderRepository {
	return &orderRepository{
		DB: db,
	}
}

func (o *orderRepository) GetOrders(id int) ([]domain.Order, error) {

	var orders []domain.Order

	if err := o.DB.Raw("select * from orders where user_id=?", id).Scan(&orders).Error; err != nil {
		return []domain.Order{}, err
	}
	return orders, nil
}

func (o *orderRepository) GetOrdersInRange(startDate, endDate time.Time) ([]domain.Order, error) {
	var orders []domain.Order

	// Execute the query to get orders within the specified time range
	if err := o.DB.Raw("SELECT * FROM orders WHERE  ordered_at BETWEEN ? AND ?", startDate, endDate).Scan(&orders).Error; err != nil {
		return []domain.Order{}, err
	}
	return orders, nil
}

func (o *orderRepository) GetCart(id int) ([]models.GetCart, error) {

	var cart []models.GetCart

	if err := o.DB.Raw("SELECT inventories.product_name,cart_products.quantity,cart_products.total_price AS Total FROM cart_products JOIN inventories ON cart_products.inventory_id=inventories.id WHERE user_id=?", id).Scan(&cart).Error; err != nil {
		return []models.GetCart{}, err
	}
	return cart, nil
}

func (o *orderRepository) GetProductNameFromID(id int) (string, error) {
	var product string

	if err := o.DB.Raw("SELECT product_name FROM inventories WHERE id=?", id).Scan(&product).Error; err != nil {
		return "", err
	}
	return product, nil
}

func (o *orderRepository) OrderItems(userid int, order models.Order, total float64) (int, error) {

	var id int
	query := `
    INSERT INTO orders (user_id,address_id ,price,payment_method_id,ordered_at)
    VALUES (?, ?, ?, ?, ?)
    RETURNING id
    `
	if err := o.DB.Raw(query, userid, order.AddressID, total, order.PaymentID, time.Now()).Scan(&id).Error; err != nil {
		return 0, err
	}

	return id, nil
}

func (o *orderRepository) AddOrderProducts(order_id int, cart []models.GetCart) error {

	query := `
    INSERT INTO order_items (order_id,inventory_id,quantity,total_price)
    VALUES (?, ?, ?, ?)
    `

	for _, v := range cart {
		var inv int
		if err := o.DB.Raw("select id from inventories where product_name=?", v.ProductName).Scan(&inv).Error; err != nil {
			return err
		}

		if err := o.DB.Exec(query, order_id, inv, v.Quantity, v.Total).Error; err != nil {
			return err
		}
	}

	return nil
}

func (o *orderRepository) CancelOrder(id int) error {

	if err := o.DB.Exec("update orders set order_status='CANCELED' where id=?", id).Error; err != nil {
		return err
	}

	return nil
}

func (o *orderRepository) EditOrderStatus(status string, id int) error {

	if err := o.DB.Exec("update orders set order_status=? where id=?", status, id).Error; err != nil {
		return err
	}

	return nil
}

func (o *orderRepository) AdminOrders(status string) ([]domain.OrderDetails, error) {

	var orders []domain.OrderDetails
	if err := o.DB.Raw("SELECT orders.id AS order_id, users.name AS username, CONCAT(addresses.house_name, ' ', addresses.street, ' ', addresses.city) AS address, payment_methods.payment_method AS payment_method, orders.price As total FROM orders JOIN users ON users.id = orders.user_id JOIN addresses ON orders.address_id = addresses.id JOIN payment_methods ON orders.payment_method_id=payment_methods.id WHERE order_status = ?", status).Scan(&orders).Error; err != nil {
		return []domain.OrderDetails{}, err
	}

	return orders, nil
}

func (o *orderRepository) CheckOrder(orderID string, userID int) error {

	var count int
	err := o.DB.Raw("select count(*) from orders where order_id = ?", orderID).Scan(&count).Error
	if err != nil {
		return err
	}
	if count < 0 {
		return errors.New("no such order exist")
	}
	var checkUser int
	err = o.DB.Raw("select user_id from orders where order_id = ?", orderID).Scan(&checkUser).Error
	if err != nil {
		return err
	}

	if userID != checkUser {
		return errors.New("the order is not did by this user")
	}

	return nil
}

func (o *orderRepository) GetOrderDetail(orderID string) (domain.Order, error) {

	var orderDetails domain.Order
	err := o.DB.Raw("select * from orders where order_id = ?", orderID).Scan(&orderDetails).Error
	if err != nil {
		return domain.Order{}, err
	}

	return orderDetails, nil
}

func (o *orderRepository) FindAmountFromOrderID(id int) (float64, error) {

	var amount float64
	err := o.DB.Raw("select final_price from orders where id = ?", id).Scan(&amount).Error
	if err != nil {
		return 0, err
	}

	return amount, nil
}

func (o *orderRepository) FindUserIdFromOrderID(id int) (int, error) {

	var user_id int
	err := o.DB.Raw("select user_id from orders where id = ?", id).Scan(&user_id).Error
	if err != nil {
		return 0, err
	}

	return user_id, nil
}
