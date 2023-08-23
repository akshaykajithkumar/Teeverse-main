package repository

import (
	interfaces "Teeverse/pkg/repository/interface"
	"Teeverse/pkg/utils/models"

	"gorm.io/gorm"
)

type otpRepository struct {
	DB *gorm.DB
}

func NewOtpRepository(DB *gorm.DB) interfaces.OtpRepository {
	return &otpRepository{
		DB: DB,
	}
}

func (ot *otpRepository) FindUserByMobileNumber(phone string) bool {

	var count int
	if err := ot.DB.Raw("select count(*) from users where phone = ?", phone).Scan(&count).Error; err != nil {
		return false
	}

	return count > 0

}

func (ot *otpRepository) UserDetailsUsingPhone(phone string) (models.UserResponse, error) {

	var usersDetails models.UserResponse
	if err := ot.DB.Raw("select * from users where phone = ?", phone).Scan(&usersDetails).Error; err != nil {
		return models.UserResponse{}, err
	}

	return usersDetails, nil

}
