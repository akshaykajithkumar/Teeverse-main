package repository

import (
	interfaces "Teeverse/pkg/repository/interface"

	"gorm.io/gorm"
)

type walletRepository struct {
	db gorm.DB
}

func NewWalletRepositoy(DB *gorm.DB) interfaces.WalletRepository {
	return &walletRepository{
		db: *DB,
	}
}

func (w *walletRepository) CreditToUserWallet(amount float64, walletId int) error {

	if err := w.db.Exec("update wallets set amount=$1 where id=$2", amount, walletId).Error; err != nil {
		return err
	}

	return nil

}

func (w *walletRepository) FindUserIdFromOrderID(id int) (int, error) {

	var user_id int
	err := w.db.Raw("select user_id from orders where id = ?", id).Scan(&user_id).Error
	if err != nil {
		return 0, err
	}

	return user_id, nil
}

func (w *walletRepository) FindWalletIdFromUserID(userId int) (int, error) {

	var count int
	err := w.db.Raw("select count(*) from wallets where user_id = ?", userId).Scan(&count).Error
	if err != nil {
		return 0, err
	}

	var walletID int
	if count > 0 {
		err := w.db.Raw("select id from wallets where user_id = ?", userId).Scan(&walletID).Error
		if err != nil {
			return 0, err
		}
	}

	return walletID, nil

}

func (w *walletRepository) CreateNewWallet(userID int) (int, error) {

	var wallet_id int
	err := w.db.Exec("Insert into wallets(user_id,amount) values($1,$2)", userID, 0).Error
	if err != nil {
		return 0, err
	}

	if err := w.db.Raw("select id from wallets where user_id=$1", userID).Scan(&wallet_id).Error; err != nil {
		return 0, err
	}

	return wallet_id, nil
}
