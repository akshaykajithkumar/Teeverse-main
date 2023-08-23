package domain

// Cart represents the wishlist of user
type Wishlist struct {
	ID     uint `json:"id" gorm:"primarykey"`
	UserID uint `json:"user_id" gorm:"not null"`
	User   User `json:"-" gorm:"foreignkey:UserID"`
}

// LineItems represents products in the wishlist of user
type WishlistItems struct {
	ID          uint      `json:"id" gorm:"primarykey"`
	WishlistID  uint      `json:"cart_id" gorm:"not null"`
	Wishlist    Wishlist  `json:"-" gorm:"foreignkey:WishlistID"`
	InventoryID uint      `json:"inventory_id" gorm:"not null"`
	Inventory   Inventory `json:"-" gorm:"foreignkey:InventoryID"`
}
