package db

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	config "Teeverse/pkg/config"
	domain "Teeverse/pkg/domain"
)

func ConnectDatabase(cfg config.Config) (*gorm.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s user=%s dbname=%s port=%s password=%s", cfg.DBHost, cfg.DBUser, cfg.DBName, cfg.DBPort, cfg.DBPassword)
	db, dbErr := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{
		SkipDefaultTransaction: true,
	})

	db.AutoMigrate(&domain.Inventory{})
	db.AutoMigrate(&domain.Category{})
	db.AutoMigrate(&domain.User{})
	db.AutoMigrate(&domain.Admin{})
	db.AutoMigrate(&domain.Cart{})
	db.AutoMigrate(&domain.Wishlist{})
	db.AutoMigrate(&domain.WishlistItems{})
	db.AutoMigrate(&domain.Address{})
	db.AutoMigrate(&domain.Order{})
	db.AutoMigrate(&domain.OrderItem{})
	db.AutoMigrate(&domain.LineItems{})
	db.AutoMigrate(&domain.PaymentMethod{})
	db.AutoMigrate(&domain.Offer{})
	db.AutoMigrate(&domain.Coupon{})
	db.AutoMigrate(&domain.Wallet{})

	return db, dbErr
}
