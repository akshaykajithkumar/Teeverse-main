package repository

import (
	"Teeverse/pkg/domain"
	interfaces "Teeverse/pkg/repository/interface"
	"errors"
	"fmt"

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

func (c *couponRepository) FindCouponDiscount(coupon string) int {
	var discountRate int
	err := c.db.Raw("select discount_rate from coupons where name=?", coupon).Scan(&discountRate).Error
	if err != nil {
		return 0
	}
	// if !coupon.Valid {
	// 	return 1
	// }
	fmt.Println("Discount:", discountRate)
	return discountRate
}

func (c *couponRepository) GetCoupons(page, limit int) ([]domain.Coupon, error) {
	if page == 0 {
		page = 1
	}
	if limit == 0 {
		limit = 10
	}
	offset := (page - 1) * limit
	var coupons []domain.Coupon

	if err := c.db.Raw("select id,name,discount_rate,valid from coupons limit ? offset ?", limit, offset).Scan(&coupons).Error; err != nil {
		return []domain.Coupon{}, err
	}

	return coupons, nil
}

func (c *couponRepository) ValidateCoupon(coupon string) (bool, error) {
	count := 0
	if err := c.db.Raw("select count(id) from coupons where name=?", coupon).Scan(&count).Error; err != nil {
		return false, err
	}
	if count < 1 {
		return false, errors.New("not a valid coupon")
	}
	valid := true
	if err := c.db.Raw("select valid from coupons where name=?", coupon).Scan(&valid).Error; err != nil {
		return false, err
	}
	if !valid {
		return false, errors.New("not a valid coupon")
	}
	return true, nil
}
