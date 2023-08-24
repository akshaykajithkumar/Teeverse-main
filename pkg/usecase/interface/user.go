package interfaces

import (
	"Teeverse/pkg/domain"
	"Teeverse/pkg/utils/models"
)

type UserUseCase interface {
	Login(user models.UserLogin) (models.TokenUser, error)
	SignUp(user models.UserDetails) (models.TokenUser, error)
	AddAddress(id int, address models.AddAddress) error
	GetAddresses(id int) ([]domain.Address, error)
	GetUserDetails(id int) (models.UserResponse, error)

	ChangePassword(id int, old string, password string, repassword string) error
	EditName(id int, name string) error
	EditEmail(id int, email string) error
	EditPhone(id int, phone string) error

	GetCartID(userID int) (int, error)
	GetCart(id, page, limit int) ([]models.GetCart, error)
	RemoveFromCart(id int, inventoryID int) error
	ClearCart(cartID int) error
	UpdateQuantityAdd(id, inv_id int) error
	UpdateQuantityLess(id, inv_id int) error
}
