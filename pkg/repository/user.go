package repository

import (
	"errors"
	"fmt"

	"Teeverse/pkg/domain"
	interfaces "Teeverse/pkg/repository/interface"
	"Teeverse/pkg/utils/models"

	"gorm.io/gorm"
)

type userDatabase struct {
	DB *gorm.DB
}

func NewUserRepository(DB *gorm.DB) interfaces.UserRepository {
	return &userDatabase{DB}
}

func (c *userDatabase) CheckUserAvailability(email string) bool {

	var count int
	query := fmt.Sprintf("select count(*) from users where email='%s'", email)
	if err := c.DB.Raw(query).Scan(&count).Error; err != nil {
		return false
	}
	// if count is greater than 0 that means the user already exist
	return count > 0
}

func (cr *userDatabase) UserBlockStatus(email string) (bool, error) {
	var permission bool
	err := cr.DB.Raw("select permission from users where email = ?", email).Scan(&permission).Error
	if err != nil {
		return false, err
	}
	return permission, nil
}

func (c *userDatabase) FindUserByEmail(user models.UserLogin) (models.UserResponse, error) {

	var user_details models.UserResponse

	err := c.DB.Raw(`
		SELECT *
		FROM users where email = ? and permission = true
		`, user.Email).Scan(&user_details).Error

	if err != nil {
		return models.UserResponse{}, errors.New("error checking user details")
	}

	return user_details, nil
}

func (c *userDatabase) FindUserIDByOrderID(orderID int) (int, error) {

	var userID int

	err := c.DB.Raw(`
		SELECT user_id
		FROM orders where id = ? 
		`, orderID).Scan(&userID).Error

	if err != nil {
		return 0, errors.New("error checking user details")
	}

	return userID, nil
}

func (c *userDatabase) SignUp(user models.UserDetails) (models.UserResponse, error) {

	var userDetails models.UserResponse
	err := c.DB.Raw("INSERT INTO users (name, email, password, phone, username) VALUES (?, ?, ?, ?,?) RETURNING id, name, email, phone", user.Name, user.Email, user.Password, user.Phone, user.Username).Scan(&userDetails).Error

	if err != nil {
		return models.UserResponse{}, err
	}

	return userDetails, nil
}

func (i *userDatabase) AddAddress(id int, address models.AddAddress, result bool) error {
	err := i.DB.Exec(`
		INSERT INTO addresses (user_id, name, house_name, street, city, state, pin,"default")
		VALUES ($1, $2, $3, $4, $5, $6, $7,$8 )
		RETURNING id`,
		id, address.Name, address.HouseName, address.Street, address.City, address.State, address.Pin, result).Error
	if err != nil {
		return err
	}

	return nil
}

func (c *userDatabase) CheckIfFirstAddress(id int) bool {

	var count int
	// query := fmt.Sprintf("select count(*) from addresses where user_id='%s'", id)
	if err := c.DB.Raw("select count(*) from addresses where user_id=?", id).Scan(&count).Error; err != nil {
		return false
	}
	// if count is greater than 0 that means the user already exist
	return count > 0

}

func (ad *userDatabase) GetAddresses(id int) ([]domain.Address, error) {

	var addresses []domain.Address

	if err := ad.DB.Raw("select * from addresses where user_id=?", id).Scan(&addresses).Error; err != nil {
		return []domain.Address{}, err
	}

	return addresses, nil

}

func (ad *userDatabase) GetUserDetails(id int) (models.UserResponse, error) {

	var details models.UserResponse

	if err := ad.DB.Raw("select id,name,username,email,phone from users where id=?", id).Scan(&details).Error; err != nil {
		return models.UserResponse{}, err
	}

	return details, nil

}

func (i *userDatabase) ChangePassword(id int, password string) error {

	err := i.DB.Exec("UPDATE users SET password=? WHERE id=?", password, id).Error
	if err != nil {
		return err
	}

	return nil

}

func (i *userDatabase) GetPassword(id int) (string, error) {

	var userPassword string
	err := i.DB.Raw("select password from users where id = ?", id).Scan(&userPassword).Error
	if err != nil {
		return "", err
	}
	return userPassword, nil

}

func (ad *userDatabase) FindIdFromPhone(phone string) (int, error) {

	var id int

	if err := ad.DB.Raw("select id from users where phone=?", phone).Scan(&id).Error; err != nil {
		return id, err
	}

	return id, nil

}

func (i *userDatabase) EditName(id int, name string) error {
	err := i.DB.Exec(`update users set name=? where id=?`, name, id).Error
	if err != nil {
		return err
	}

	return nil
}

func (i *userDatabase) EditEmail(id int, email string) error {
	err := i.DB.Exec(`update users set email=? where id=?`, email, id).Error
	if err != nil {
		return err
	}

	return nil
}

func (i *userDatabase) EditPhone(id int, phone string) error {
	err := i.DB.Exec(`update users set phone=? where id=?`, phone, id).Error
	if err != nil {
		return err
	}

	return nil
}

func (ad *userDatabase) RemoveFromCart(cartID int, inventoryID int) error {

	if err := ad.DB.Exec(`delete from line_items where cart_id=? and inventory_id=?`, cartID, inventoryID).Error; err != nil {
		return err
	}

	return nil

}

func (ad *userDatabase) ClearCart(cartID int) error {

	if err := ad.DB.Exec(`delete from line_items where cart_id=?`, cartID).Error; err != nil {
		return err
	}

	return nil

}

func (ad *userDatabase) UpdateQuantityAdd(id, inv_id int) error {
	query := `
		UPDATE line_items
		SET quantity = quantity + 1
		WHERE cart_id=? AND inventory_id=?
	`

	result := ad.DB.Exec(query, id, inv_id)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (ad *userDatabase) UpdateQuantityLess(id, inv_id int) error {
	query := `
		UPDATE line_items
		SET quantity = quantity - 1
		WHERE cart_id=? AND inventory_id=?
	`

	result := ad.DB.Exec(query, id, inv_id)
	if result.Error != nil {
		return result.Error
	}

	return nil

}

func (cr *userDatabase) FindUserByOrderID(orderId string) (domain.User, error) {

	var userDetails domain.User
	err := cr.DB.Raw("select users.name,users.email,users.phone from users inner join orders on orders.user_id = users.id where order_id = ?", orderId).Scan(&userDetails).Error
	if err != nil {
		return domain.User{}, err
	}

	return userDetails, nil
}

func (ad *userDatabase) GetCartID(id int) (int, error) {

	var cart_id int

	if err := ad.DB.Raw("select id from carts where user_id=?", id).Scan(&cart_id).Error; err != nil {
		return 0, err
	}

	return cart_id, nil

}

func (ad *userDatabase) GetProductsInCart(cart_id int) ([]int, error) {

	var cart_products []int

	if err := ad.DB.Raw("select inventory_id from line_items where cart_id=?", cart_id).Scan(&cart_products).Error; err != nil {
		return []int{}, err
	}

	return cart_products, nil

}

func (ad *userDatabase) FindProductNames(inventory_id int) (string, error) {

	var product_name string

	if err := ad.DB.Raw("select product_name from inventories where id=?", inventory_id).Scan(&product_name).Error; err != nil {
		return "", err
	}

	return product_name, nil

}

func (ad *userDatabase) FindCartQuantity(cart_id, inventory_id int) (int, error) {

	var quantity int

	if err := ad.DB.Raw("select quantity from line_items where cart_id=? and inventory_id=?", cart_id, inventory_id).Scan(&quantity).Error; err != nil {
		return 0, err
	}

	return quantity, nil

}

func (ad *userDatabase) FindPrice(inventory_id int) (float64, error) {

	var price float64

	if err := ad.DB.Raw("select price from inventories where id=?", inventory_id).Scan(&price).Error; err != nil {
		return 0, err
	}

	return price, nil

}

func (ad *userDatabase) FindCategory(inventory_id int) (int, error) {

	var category int

	if err := ad.DB.Raw("select category_id from inventories where id=?", inventory_id).Scan(&category).Error; err != nil {
		return 0, err
	}

	return category, nil
}
