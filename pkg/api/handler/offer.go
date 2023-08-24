package handler

import (
	services "Teeverse/pkg/usecase/interface"
	models "Teeverse/pkg/utils/models"
	"Teeverse/pkg/utils/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type OfferHandler struct {
	offerUseCase services.OfferUseCase
}

func NewOfferHandler(offerUsecase services.OfferUseCase) *OfferHandler {
	return &OfferHandler{
		offerUseCase: offerUsecase,
	}
}

// @Summary		Add Offer
// @Description	Admin can add new  offers
// @Tags			Admin
// @Accept			json
// @Produce		    json
// @Param           offer      body     models.CreateOffer   true   "Offer"
// @Security		Bearer
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/admin/offers/create [post]
func (o *OfferHandler) AddOffer(c *gin.Context) {

	var offer models.CreateOffer

	if err := c.BindJSON(&offer); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	err := o.offerUseCase.AddNewOffer(offer)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "Could not add the Offer", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Successfully added Offer", nil, nil)
	c.JSON(http.StatusOK, successRes)

}

// @Summary		Expire Offer
// @Description	Admin can add Expire  offers
// @Tags			Admin
// @Accept			json
// @Produce		    json
// @Param           catID      query     string   true   "Category ID"
// @Security		Bearer
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/admin/offers/expire [post]
func (o *OfferHandler) ExpireValidity(c *gin.Context) {
	catIDStr := c.Query("catID")
	catID, err := strconv.Atoi(catIDStr)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "type conversion failed", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	err = o.offerUseCase.MakeOfferExpire(catID)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Successfully turned the offer invalid", nil, nil)
	c.JSON(http.StatusOK, successRes)
}
