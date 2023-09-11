package interfaces

import (
	"Teeverse/pkg/domain"
)

type WalletRepository interface {
	CreditToUserWallet(amount float64, walletID int) error
	FindUserIdFromOrderID(id int) (int, error)
	FindWalletIdFromUserID(userId int) (int, error)
	CreateNewWallet(userID int) (int, error)
	GetBalance(WalletID int) (int, error)
	GetHistory(walletID, page, limit int) ([]domain.WalletHistory, error)
	AddHistory(amount, WalletID int, purpose string) error
	PayFromWallet(userID, orderID int, price float64) (float64, error)
}
