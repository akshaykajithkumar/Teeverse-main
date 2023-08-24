package interfaces

import "Teeverse/pkg/utils/models"

type OfferRepository interface {
	AddNewOffer(offer models.CreateOffer) error
	MakeOfferExpire(catID int) error
	FindDiscountPercentage(catID int) (int, error)
}
