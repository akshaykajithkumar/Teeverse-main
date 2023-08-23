package interfaces

import "Teeverse/pkg/utils/models"

type OtpUseCase interface {
	VerifyOTP(code models.VerifyData) (models.TokenUser, error)
	SendOTP(phone string) error
}
