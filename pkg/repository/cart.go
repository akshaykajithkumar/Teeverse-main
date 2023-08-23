package repository

import (
	"Teeverse/pkg/domain"
	interfaces "Teeverse/pkg/repository/interface"

	"gorm.io/gorm"
)

type cartRepository struct {
	DB *gorm.DB
}

func NewCartRepository(db *gorm.DB) interfaces.CartRepository {
	return &cartRepository{
		DB: db,
	}
}

func (ad *cartRepository) GetAddresses(id int) ([]domain.Address, error) {

	var addresses []domain.Address

	if err := ad.DB.Raw("SELECT * FROM addresses WHERE user_id=?", id).Scan(&addresses).Error; err != nil {
		return []domain.Address{}, err
	}

	return addresses, nil

}

func (ad *cartRepository) GetCartId(user_id int) (int, error) {

	var id int

	if err := ad.DB.Raw("SELECT id FROM carts WHERE user_id=?", user_id).Scan(&id).Error; err != nil {
		return 0, err
	}

	return id, nil

}

func (i *cartRepository) CreateNewCart(user_id int) (int, error) {
	var id int
	err := i.DB.Exec(`
		INSERT INTO carts (user_id)
		VALUES (?)`, user_id).Error
	if err != nil {
		return 0, err
	}

	if err := i.DB.Raw("select id from carts where user_id=?", user_id).Scan(&id).Error; err != nil {
		return 0, err
	}

	return id, nil
}

func (i *cartRepository) AddLineItems(cart_id, inventory_id int) error {

	err := i.DB.Exec(`
		INSERT INTO line_items (cart_id,inventory_id)
		VALUES (?,?)`, cart_id, inventory_id).Error
	if err != nil {
		return err
	}

	return nil
}
