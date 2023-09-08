package usecase

import (
	interfaces "Teeverse/pkg/repository/interface"
	services "Teeverse/pkg/usecase/interface"
	"errors"

	"Teeverse/pkg/utils/models"
)

type cartUseCase struct {
	repo                interfaces.CartRepository
	inventoryRepository interfaces.InventoryRepository
	userUseCase         services.UserUseCase
	paymentUseCase      services.PaymentUseCase
}

func NewCartUseCase(repo interfaces.CartRepository, inventoryRepo interfaces.InventoryRepository, userUseCase services.UserUseCase, paymentUseCase services.PaymentUseCase) *cartUseCase {
	return &cartUseCase{
		repo:                repo,
		inventoryRepository: inventoryRepo,
		userUseCase:         userUseCase,
		paymentUseCase:      paymentUseCase,
	}
}

func (i *cartUseCase) AddToCart(user_id, inventory_id int) error {
	//check if the desired product has quantity available
	stock, err := i.inventoryRepository.CheckStock(inventory_id)
	if err != nil {
		return err
	}
	//if available then call userRepository
	if stock <= 0 {
		return errors.New("out of stock")
	}

	//find user cart id
	cart_id, err := i.repo.GetCartId(user_id)
	if err != nil {
		return errors.New("some error in geting user cart")
	}
	//if user has no existing cart create new cart
	if cart_id == 0 {
		cart_id, err = i.repo.CreateNewCart(user_id)
		if err != nil {
			return errors.New("cannot create cart from user")
		}
	}
	//check if already added
	if i.repo.CheckIfInvAdded(inventory_id, cart_id) {
		err := i.repo.AddQuantity(inventory_id, cart_id)
		if err != nil {
			return err
		}
		return nil
	}

	//add product to line items
	if err := i.repo.AddLineItems(cart_id, inventory_id); err != nil {
		return errors.New("error in adding products")
	}

	return nil
}

func (i *cartUseCase) CheckOut(id int) (models.CheckOut, error) {

	address, err := i.repo.GetAddresses(id)
	if err != nil {
		return models.CheckOut{}, err
	}
	//cartchange
	products, err := i.userUseCase.GetCart(id, 0, 0)
	if err != nil {
		return models.CheckOut{}, err
	}

	paymentmethods, err := i.paymentUseCase.GetPaymentMethods()
	if err != nil {
		return models.CheckOut{}, err
	}

	var price float64
	for _, v := range products {
		price = price + v.DiscountedPrice
	}

	var checkout models.CheckOut

	checkout.Addresses = address
	checkout.Products = products
	checkout.PaymentMethods = paymentmethods
	checkout.TotalPrice = price

	return checkout, err
}
