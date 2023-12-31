package interfaces

import (
	"Teeverse/pkg/domain"
)

type CouponRepository interface {
	AddCoupon(domain.Coupon) error
	MakeCouponInvalid(id int) error
	FindCouponDiscount(coupon string) int
	GetCoupons(page, limit int) ([]domain.Coupon, error)
	ValidateCoupon(coupon string) (bool, error)
}
