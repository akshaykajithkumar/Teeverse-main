package interfaces

import (
	"Teeverse/pkg/domain"
)

type CouponRepository interface {
	AddCoupon(domain.Coupon) error
	MakeCouponInvalid(id int) error
	FindCouponDiscount(couponID int) int
}
