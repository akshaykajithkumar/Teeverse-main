package interfaces

import "Teeverse/pkg/utils/models"

type WishlistRepository interface {
	GetWishlist(id int) ([]models.GetWishlist, error)
	GetWishlistID(user_id int) (int, error)
	CreateNewWishlist(user_id int) (int, error)
	AddWishlistItems(wishlistID, inventory_id int) error
	GetProductsInWishlist(wishlistID int) ([]int, error)
	FindProductNames(inventory_id int) (string, error)
	FindPrice(inventory_id int) (float64, error)
	FindCategory(inventory_id int) (string, error)
	RemoveFromWishlist(wishlistID int, inventoryID int) error
}
