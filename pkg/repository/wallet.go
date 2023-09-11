package repository

import (
	"Teeverse/pkg/domain"
	interfaces "Teeverse/pkg/repository/interface"
	"fmt"
	"time"

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

	if err := w.db.Exec("update wallets set amount=amount+$1 where id=$2", amount, walletId).Error; err != nil {
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

func (w *walletRepository) GetBalance(walletID int) (int, error) {

	var balance float64
	if err := w.db.Raw("select amount from wallets where id=$1", walletID).Scan(&balance).Error; err != nil {
		return 0, err
	}
	fmt.Println(walletID, balance)
	return int(balance), nil
}

func (w *walletRepository) GetHistory(walletID, page, limit int) ([]domain.WalletHistory, error) {
	if page == 0 {
		page = 1
	}
	if limit == 0 {
		limit = 10
	}
	offset := (page - 1) * limit
	var history []domain.WalletHistory
	if err := w.db.Raw("select * from wallet_histories where wallet_id=? limit ? offset ?", walletID, limit, offset).Scan(&history).Error; err != nil {
		return []domain.WalletHistory{}, err
	}

	return history, nil
}

func (w *walletRepository) AddHistory(amount, WalletID int, purpose string) error {

	err := w.db.Exec("Insert into wallet_histories(wallet_id,amount,purpose,time) values(?,?,?,?)", WalletID, amount, purpose, time.Now()).Error
	if err != nil {
		return err
	}
	return nil
}

func (w *walletRepository) PayFromWallet(userID, orderID int, price float64) (float64, error) {
	if err := w.db.Exec("UPDATE wallets SET amount=amount-? WHERE user_id=?", price, userID).Error; err != nil {
		return 0.0, err
	}
	if err := w.db.Exec("UPDATE orders SET payment_status='PAID' WHERE id=?", orderID).Error; err != nil {
		return 0.0, err
	}
	var balance float64
	if err := w.db.Raw("select amount from wallets where user_id=$1", userID).Scan(&balance).Error; err != nil {
		return 0, err
	}
	return balance, nil
}
