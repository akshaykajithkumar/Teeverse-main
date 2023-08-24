package usecase

import (
	"Teeverse/pkg/domain"
	interfaces "Teeverse/pkg/repository/interface"
	services "Teeverse/pkg/usecase/interface"
)

type couponUseCase struct {
	couponRepo interfaces.CouponRepository
}

func NewCouponUseCase(couponRepo interfaces.CouponRepository) services.CouponUseCase {
	return &couponUseCase{
		couponRepo: couponRepo,
	}
}

func (coup *couponUseCase) AddCoupon(coupon domain.Coupon) error {
	if err := coup.couponRepo.AddCoupon(coupon); err != nil {
		return err
	}

	return nil
}

func (coup *couponUseCase) MakeCouponInvalid(id int) error {
	if err := coup.couponRepo.MakeCouponInvalid(id); err != nil {
		return err
	}

	return nil
}
