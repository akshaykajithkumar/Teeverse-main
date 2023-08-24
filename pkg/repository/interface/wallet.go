package interfaces

type WalletRepository interface {
	CreditToUserWallet(amount float64, walletID int) error
	FindUserIdFromOrderID(id int) (int, error)
	FindWalletIdFromUserID(userId int) (int, error)
	CreateNewWallet(userID int) (int, error)
}
