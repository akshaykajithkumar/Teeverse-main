package interfaces

import "Teeverse/pkg/utils/models"

type OfferUseCase interface {
	AddNewOffer(model models.CreateOffer) error
	MakeOfferExpire(catID int) error
}
