package interfaces

import (
	"Teeverse/pkg/domain"
	"Teeverse/pkg/utils/models"
)

type OfferRepository interface {
	AddNewOffer(offer models.CreateOffer) error
	MakeOfferExpire(catID int) error
	FindDiscountPercentage(cat string) (int, error)
	GetOffers(page, limit int) ([]domain.Offer, error)
}
