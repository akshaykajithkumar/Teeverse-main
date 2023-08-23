package handler

import (
	"Teeverse/pkg/helper"
	services "Teeverse/pkg/usecase/interface"
	"Teeverse/pkg/utils/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type WishlistHandler struct {
	usecase services.WishlistUseCase
}

func NewWishlistHandler(usecase services.WishlistUseCase) *WishlistHandler {
	return &WishlistHandler{
		usecase: usecase,
	}
}

// @Summary		Add To Wishlist
// @Description	Add products to Wishlsit  for the purchase
// @Tags			User
// @Accept			json
// @Produce		json
// @Param			inventory	query	string	true	"inventory ID"
// @Security		Bearer
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/users/home/add-to-wishlist [post]
func (i *WishlistHandler) AddToWishlist(c *gin.Context) {

	userID, err := helper.GetUserID(c)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "Could not get userID", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	inventoryID, err := strconv.Atoi(c.Query("inventory"))
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "check parameters properly", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	if err := i.usecase.AddToWishlist(userID, inventoryID); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "Could not add the wishlsit", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Successfully added To Wishlist", nil, nil)
	c.JSON(http.StatusOK, successRes)

}

// @Summary		Remove from Wishlist
// @Description	user can remove products from their wishlist
// @Tags			User
// @Produce		    json
// @Param			inventory	query	string	true	"inventory id"
// @Security		Bearer
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/users/wishlist/remove [delete]
func (i *WishlistHandler) RemoveFromWishlist(c *gin.Context) {
	id, err := helper.GetUserID(c)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "Could not get userID", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	wishlistID, err := i.usecase.GetWishlistID(id)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "check parameters properly", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	inv, err := strconv.Atoi(c.Query("inventory"))
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "check parameters properly", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	if err := i.usecase.RemoveFromWishlist(wishlistID, inv); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "could not retrieve cart", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Successfully Removed product from cart", nil, nil)
	c.JSON(http.StatusOK, successRes)
}

// @Summary		Get Wishlist
// @Description	user can view their wishlist details
// @Tags			User
// @Produce		    json
// @Security		Bearer
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/users/wishlist [get]
func (i *WishlistHandler) GetWishlist(c *gin.Context) {
	id, err := helper.GetUserID(c)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "Could not get userID", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	products, err := i.usecase.GetWishlist(id)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "could not retrieve wishlist", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "Successfully got all products in wishlist", products, nil)
	c.JSON(http.StatusOK, successRes)
}
