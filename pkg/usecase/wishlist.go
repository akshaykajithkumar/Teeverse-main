package usecase

import (
	interfaces "Teeverse/pkg/repository/interface"
	services "Teeverse/pkg/usecase/interface"
	"Teeverse/pkg/utils/models"
	"errors"
)

type wishlistUseCase struct {
	wishlistRepo interfaces.WishlistRepository
}

func NewWishlistUseCase(repo interfaces.WishlistRepository) services.WishlistUseCase {
	return &wishlistUseCase{
		wishlistRepo: repo,
	}
}

func (w *wishlistUseCase) GetWishlistID(userID int) (int, error) {
	wishlistID, err := w.wishlistRepo.GetWishlistID(userID)
	if err != nil {
		return 0, err
	}
	return wishlistID, nil
}

func (w *wishlistUseCase) GetWishlist(id int) ([]models.GetWishlist, error) {

	//find cart id
	wishlistID, err := w.wishlistRepo.GetWishlistID(id)
	if err != nil {
		return []models.GetWishlist{}, err
	}
	//find products inside cart
	products, err := w.wishlistRepo.GetProductsInWishlist(wishlistID)
	if err != nil {
		return []models.GetWishlist{}, err
	}
	//find product names
	var product_names []string
	for i := range products {
		product_name, err := w.wishlistRepo.FindProductNames(products[i])
		if err != nil {
			return []models.GetWishlist{}, err
		}
		product_names = append(product_names, product_name)
	}

	var price []float64
	for i := range products {
		q, err := w.wishlistRepo.FindPrice(products[i])
		if err != nil {
			return []models.GetWishlist{}, err
		}
		price = append(price, q)
	}

	var categories []string
	for i := range products {
		c, err := w.wishlistRepo.FindCategory(products[i])
		if err != nil {
			return []models.GetWishlist{}, err
		}
		categories = append(categories, c)
	}

	var getwishlist []models.GetWishlist
	for i := range product_names {
		var get models.GetWishlist
		get.ProductName = product_names[i]
		get.Category = categories[i]
		get.Price = price[i]

		getwishlist = append(getwishlist, get)
	}

	return getwishlist, nil

}

func (w *wishlistUseCase) RemoveFromWishlist(wishlistID int, inventoryID int) error {

	err := w.wishlistRepo.RemoveFromWishlist(wishlistID, inventoryID)
	if err != nil {
		return err
	}

	return nil

}

func (w *wishlistUseCase) AddToWishlist(user_id, inventory_id int) error {

	//find user wishlist id
	wishlistID, err := w.wishlistRepo.GetWishlistID(user_id)
	if err != nil {
		return errors.New("some error in geting user wishlist")
	}
	//if user has no existing cart create new cart
	if wishlistID == 0 {
		wishlistID, err = w.wishlistRepo.CreateNewWishlist(user_id)
		if err != nil {
			return errors.New("cannot create wishlist from user")
		}
	}

	//add product to line items
	if err := w.wishlistRepo.AddWishlistItems(wishlistID, inventory_id); err != nil {
		return errors.New("error in adding products")
	}

	return nil
}
