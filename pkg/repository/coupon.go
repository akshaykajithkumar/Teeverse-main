package repository

import (
	"Teeverse/pkg/domain"
	interfaces "Teeverse/pkg/repository/interface"

	"gorm.io/gorm"
)

type couponRepository struct {
	db *gorm.DB
}

func NewCouponRepository(DB *gorm.DB) interfaces.CouponRepository {
	return &couponRepository{
		db: DB,
	}
}

func (c *couponRepository) AddCoupon(coup domain.Coupon) error {
	if err := c.db.Exec("INSERT INTO coupons(name,discount_rate,valid) values($1,$2,$3)", coup.Name, coup.DiscountRate, coup.Valid).Error; err != nil {
		return err
	}

	return nil
}

func (c *couponRepository) MakeCouponInvalid(id int) error {
	if err := c.db.Exec("UPDATE coupons SET valid=false where id=$1", id).Error; err != nil {
		return err
	}

	return nil
}

func (c *couponRepository) FindCouponDiscount(couponID int) int {
	var coupon domain.Coupon
	err := c.db.Raw("select name,discount_rate,valid from coupons where id=$1", couponID).Scan(&coupon).Error
	if err != nil {
		return 0
	}
	if !coupon.Valid {
		return 1
	}

	return coupon.DiscountRate
}
