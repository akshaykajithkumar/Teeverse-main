package interfaces

import "Teeverse/pkg/utils/models"

type WishlistUseCase interface {
	AddToWishlist(user_id, inventory_id int) error
	GetWishlistID(userID int) (int, error)
	GetWishlist(id int) ([]models.GetWishlist, error)
	RemoveFromWishlist(id int, inventoryID int) error
}
