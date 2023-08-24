package repository

import (
	interfaces "Teeverse/pkg/repository/interface"
	"Teeverse/pkg/utils/models"

	"gorm.io/gorm"
)

type offerRepository struct {
	DB *gorm.DB
}

func NewOfferRepository(DB *gorm.DB) interfaces.OfferRepository {
	return &offerRepository{
		DB: DB,
	}
}

func (o *offerRepository) AddNewOffer(offer models.CreateOffer) error {
	if err := o.DB.Exec("INSERT INTO offers(category_id,discount_rate) values($1,$2)", offer.CategoryID, offer.Discount).Error; err != nil {
		return err
	}

	return nil
}

func (o *offerRepository) MakeOfferExpire(catID int) error {
	if err := o.DB.Exec("UPDATE offers SET valid=false where id=$1", catID).Error; err != nil {
		return err
	}

	return nil
}

func (o *offerRepository) FindDiscountPercentage(catID int) (int, error) {
	var percentage int
	err := o.DB.Raw("select discount_rate from offers where category_id=$1 and valid=true", catID).Scan(&percentage).Error
	if err != nil {
		return 0, err
	}

	return percentage, nil
}
