package interfaces

import (
	"Teeverse/pkg/utils/models"
)

type OtpRepository interface {
	FindUserByMobileNumber(phone string) bool
	UserDetailsUsingPhone(phone string) (models.UserResponse, error)
}
