package domain

import "time"

// Order represents the order of user
//
//	type Order struct {
//		//gorm.Model      `json:"-"`
//		ID              int     `json:"id" gorm:"primaryKey;autoIncrement"`
//		UserID          int     `json:"user_id" gorm:"not null"`
//		User            User    `json:"-" gorm:"foreignkey:UserID"`
//		AddressID       int     `json:"address_id" gorm:"not null"`
//		Address         Address `json:"-" gorm:"foreignkey:AddressID"`
//		PaymentMethodID int     `json:"paymentmethodID" gorm:"default:1"`
//		// PaymentMethod   PaymentMethod `json:"-" gorm:"foreignkey:PaymentMethodID"`
//		PaymentMethod PaymentMethod `json:"paymentMethod" gorm:"foreignKey:PaymentMethodID"`
//		PaymentID     string        `json:"paymentID"`
//		Price         float64       `json:"price"`
//		OrderedAt     time.Time     `json:"orderedAt"`
//		OrderStatus   string        `json:"order_status" gorm:"order_status:4;default:'PENDING';check:order_status IN ('PENDING', 'SHIPPED','DELIVERED','CANCELED','RETURNED')"`
//		PaymentStatus string        `json:"paymentStatus" gorm:"default:'Pending'"`
//	}
type Order struct {
	ID        int     `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID    int     `json:"user_id" gorm:"not null"`
	User      User    `json:"-" gorm:"foreignkey:UserID"`
	AddressID int     `json:"address_id" gorm:"not null"`
	Address   Address `json:"-" gorm:"foreignkey:AddressID"`
	//PaymentMethodID int     `json:"-" gorm:"column:payment_method_id"`
	// PaymentMethod   PaymentMethod `json:"paymentMethod"`
	PaymentMethodID int       `json:"paymentmethodID" gorm:"default:1"`
	PaymentMethod   string    `json:"paymentMethod" gorm:"column:payment_method"`
	PaymentID       string    `json:"paymentID"`
	Price           float64   `json:"price"`
	OrderedAt       time.Time `json:"orderedAt"`
	OrderStatus     string    `json:"order_status" gorm:"order_status:4;default:'PENDING';check:order_status IN ('PENDING', 'SHIPPED','DELIVERED','CANCELED','RETURNED')"`
	PaymentStatus   string    `json:"paymentStatus" gorm:"default:'Pending'"`
}

// OrderItem represents the product details of the order
type OrderItem struct {
	ID          int       `json:"id" gorm:"primaryKey;autoIncrement"`
	OrderID     int       `json:"order_id"`
	Order       Order     `json:"-" gorm:"foreignkey:OrderID"`
	InventoryID int       `json:"inventory_id"`
	Inventory   Inventory `json:"-" gorm:"foreignkey:InventoryID"`
	Quantity    int       `json:"quantity"`
	TotalPrice  float64   `json:"total_price"`
}

// AdminOrdersResponse represents the order details with order status
type AdminOrdersResponse struct {
	Pending   []OrderDetails
	Shipped   []OrderDetails
	Delivered []OrderDetails
	Canceled  []OrderDetails
	Returned  []OrderDetails
}

// OrderDetails represents the details of order
type OrderDetails struct {
	ID            int    `json:"order_id" gorm:"column:order_id"`
	Username      string `json:"name"`
	Address       string `json:"address"`
	PaymentMethod string `json:"paymentmethod"`
	//PaymentMethod   PaymentMethod `json:"-" gorm:"foreignkey:PaymentMethodID"`
	Total float64 `json:"total"`
}

type PaymentMethod struct {
	ID            int    `gorm:"primaryKey"`
	PaymentMethod string `json:"PaymentMethod" validate:"required" gorm:"unique"`
}

type SalesReport struct {
	Orders       []Order
	TotalRevenue float64
	TotalOrders  int
	BestSellers  []string
}

type ProductReport struct {
	InventoryID int
	Quantity    int
}
