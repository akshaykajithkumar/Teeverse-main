package usecase

import (
	"Teeverse/pkg/domain"
	"Teeverse/pkg/helper"
	interfaces "Teeverse/pkg/repository/interface"
	services "Teeverse/pkg/usecase/interface"
	"Teeverse/pkg/utils/models"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type userUseCase struct {
	userRepo  interfaces.UserRepository
	offerRepo interfaces.OfferRepository
}

func NewUserUseCase(repo interfaces.UserRepository, offer interfaces.OfferRepository) services.UserUseCase {
	return &userUseCase{
		userRepo:  repo,
		offerRepo: offer,
	}
}

func (u *userUseCase) Login(user models.UserLogin) (models.TokenUser, error) {
	// checking if a username exist with this email address
	ok := u.userRepo.CheckUserAvailability(user.Email)
	if !ok {
		return models.TokenUser{}, errors.New("the user does not exist")
	}

	permission, err := u.userRepo.UserBlockStatus(user.Email)
	if err != nil {
		return models.TokenUser{}, err
	}

	if !permission {
		return models.TokenUser{}, errors.New("user is blocked by admin")
	}

	// Get the user details in order to check the password, in this case ( The same function can be reused in future )
	user_details, err := u.userRepo.FindUserByEmail(user)
	if err != nil {
		return models.TokenUser{}, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user_details.Password), []byte(user.Password))
	if err != nil {
		return models.TokenUser{}, errors.New("password incorrect")
	}

	tokenString, err := helper.GenerateTokenUser(user_details)
	if err != nil {
		return models.TokenUser{}, errors.New("could not create token")
	}
	return models.TokenUser{
		Username: user_details.Username,
		Token:    tokenString,
	}, nil

}

func (u *userUseCase) SignUp(user models.UserDetails) (models.TokenUser, error) {
	// Check whether the user already exist. If yes, show the error message, since this is signUp
	userExist := u.userRepo.CheckUserAvailability(user.Email)
	if userExist {
		return models.TokenUser{}, errors.New("user already exist, sign in")
	}
	if user.Password != user.ConfirmPassword {
		return models.TokenUser{}, errors.New("password does not match")
	}

	// Hash password since details are validated

	hashedPassword, err := helper.PasswordHashing(user.Password)
	if err != nil {
		return models.TokenUser{}, err
	}

	user.Password = hashedPassword

	// add user details to the database
	userData, err := u.userRepo.SignUp(user)
	if err != nil {
		return models.TokenUser{}, err
	}

	// crete a JWT token string for the user
	tokenString, err := helper.GenerateTokenUser(userData)
	if err != nil {
		return models.TokenUser{}, errors.New("could not create token due to some internal error")
	}

	return models.TokenUser{
		Username: user.Username,
		Token:    tokenString,
	}, nil
}

func (i *userUseCase) AddAddress(id int, address models.AddAddress) error {

	rslt := i.userRepo.CheckIfFirstAddress(id)
	var result bool

	if !rslt {
		result = true
	} else {
		result = false
	}

	err := i.userRepo.AddAddress(id, address, result)
	if err != nil {
		return err
	}

	return nil

}

func (i *userUseCase) GetAddresses(id int) ([]domain.Address, error) {

	addresses, err := i.userRepo.GetAddresses(id)
	if err != nil {
		return []domain.Address{}, err
	}

	return addresses, nil

}

func (i *userUseCase) GetUserDetails(id int) (models.UserResponse, error) {

	details, err := i.userRepo.GetUserDetails(id)
	if err != nil {
		return models.UserResponse{}, err
	}

	return details, nil

}

func (i *userUseCase) ChangePassword(id int, old string, password string, repassword string) error {

	userPassword, err := i.userRepo.GetPassword(id)
	if err != nil {
		return err
	}

	err = bcrypt.CompareHashAndPassword([]byte(userPassword), []byte(old))
	if err != nil {
		return errors.New("password incorrect")
	}

	if password != repassword {
		return errors.New("passwords does not match")
	}

	newpassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return errors.New("internal server error")
	}

	return i.userRepo.ChangePassword(id, string(newpassword))

}

func (i *userUseCase) EditName(id int, name string) error {

	err := i.userRepo.EditName(id, name)
	if err != nil {
		return err
	}

	return nil

}

func (i *userUseCase) EditEmail(id int, email string) error {

	err := i.userRepo.EditEmail(id, email)
	if err != nil {
		return err
	}

	return nil

}

func (i *userUseCase) EditPhone(id int, phone string) error {

	err := i.userRepo.EditPhone(id, phone)
	if err != nil {
		return err
	}

	return nil

}

func (u *userUseCase) GetCartID(userID int) (int, error) {
	cartID, err := u.userRepo.GetCartID(userID)
	if err != nil {
		return 0, err
	}
	return cartID, nil
}

func (u *userUseCase) GetCart(id, page, limit int) ([]models.GetCart, error) {

	//find cart id
	cart_id, err := u.userRepo.GetCartID(id)
	if err != nil {
		return []models.GetCart{}, err
	}
	//find products inside cart
	products, err := u.userRepo.GetProductsInCart(cart_id, page, limit)
	if err != nil {
		return []models.GetCart{}, err
	}
	//find product names
	var product_names []string
	for i := range products {
		product_name, err := u.userRepo.FindProductNames(products[i])
		if err != nil {
			return []models.GetCart{}, err
		}
		product_names = append(product_names, product_name)
	}

	//find quantity
	var quantity []int
	for i := range products {
		q, err := u.userRepo.FindCartQuantity(cart_id, products[i])
		if err != nil {
			return []models.GetCart{}, err
		}
		quantity = append(quantity, q)
	}

	var price []float64
	for i := range products {
		q, err := u.userRepo.FindPrice(products[i])
		if err != nil {
			return []models.GetCart{}, err
		}
		price = append(price, q)
	}

	var categories []int
	for i := range products {
		c, err := u.userRepo.FindCategory(products[i])
		if err != nil {
			return []models.GetCart{}, err
		}
		categories = append(categories, c)
	}

	var getcart []models.GetCart
	for i := range product_names {
		var get models.GetCart
		get.ProductName = product_names[i]
		get.Category_id = categories[i]
		get.Quantity = quantity[i]
		get.Total = price[i]

		getcart = append(getcart, get)
	}

	//find offers
	var offers []int
	for i := range categories {
		c, err := u.offerRepo.FindDiscountPercentage(categories[i])
		if err != nil {
			return []models.GetCart{}, err
		}
		offers = append(offers, c)
	}

	//find discounted price
	for i := range getcart {
		getcart[i].DiscountedPrice = (getcart[i].Total) - (getcart[i].Total * float64(offers[i]) / 100)
	}

	return getcart, nil

}

func (i *userUseCase) RemoveFromCart(cartID int, inventoryID int) error {

	err := i.userRepo.RemoveFromCart(cartID, inventoryID)
	if err != nil {
		return err
	}

	return nil

}
func (i *userUseCase) ClearCart(cartID int) error {

	err := i.userRepo.ClearCart(cartID)
	if err != nil {
		return err
	}

	return nil

}

func (i *userUseCase) UpdateQuantityAdd(id, inv_id int) error {

	err := i.userRepo.UpdateQuantityAdd(id, inv_id)
	if err != nil {
		return err
	}

	return nil

}

func (i *userUseCase) UpdateQuantityLess(id, inv_id int) error {

	err := i.userRepo.UpdateQuantityLess(id, inv_id)
	if err != nil {
		return err
	}

	return nil

}
