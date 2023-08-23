package usecase

import (
	"errors"

	config "Teeverse/pkg/config"
	interfaces "Teeverse/pkg/repository/interface"
	services "Teeverse/pkg/usecase/interface"
	"Teeverse/pkg/utils/models"

	"Teeverse/pkg/helper"
)

type otpUseCase struct {
	cfg           config.Config
	otpRepository interfaces.OtpRepository
}

func NewOtpUseCase(cfg config.Config, repo interfaces.OtpRepository) services.OtpUseCase {
	return &otpUseCase{
		cfg:           cfg,
		otpRepository: repo,
	}
}

func (ot *otpUseCase) SendOTP(phone string) error {

	ok := ot.otpRepository.FindUserByMobileNumber(phone)
	if !ok {
		return errors.New("the user does not exist")
	}

	helper.TwilioSetup(ot.cfg.ACCOUNTSID, ot.cfg.AUTHTOKEN)
	_, err := helper.TwilioSendOTP(phone, ot.cfg.SERVICESID)
	if err != nil {
		return errors.New("error ocurred while generating OTP")
	}

	return nil

}

func (ot *otpUseCase) VerifyOTP(code models.VerifyData) (models.TokenUser, error) {

	helper.TwilioSetup(ot.cfg.ACCOUNTSID, ot.cfg.AUTHTOKEN)
	err := helper.TwilioVerifyOTP(ot.cfg.SERVICESID, code.Code, code.PhoneNumber)
	if err != nil {
		//this guard clause catches the error code runs only until here
		return models.TokenUser{}, errors.New("error while verifying")
	}

	// if user is authenticated using OTP send back user details
	userDetails, err := ot.otpRepository.UserDetailsUsingPhone(code.PhoneNumber)
	if err != nil {
		return models.TokenUser{}, err
	}

	tokenString, err := helper.GenerateTokenUser(userDetails)
	if err != nil {
		return models.TokenUser{}, err

	}
	return models.TokenUser{
		Username: userDetails.Username,
		Token:    tokenString,
	}, nil
}
