package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"Teeverse/pkg/helper"
	services "Teeverse/pkg/usecase/interface"
	"Teeverse/pkg/utils/models"
	"Teeverse/pkg/utils/response"
)

type UserHandler struct {
	userUseCase services.UserUseCase
}

func NewUserHandler(usecase services.UserUseCase) *UserHandler {
	return &UserHandler{
		userUseCase: usecase,
	}
}

// @Summary		Add Address
// @Description	user can add their addresses
// @Tags			User
// @Accept			json
// @Produce		    json
// @Param			address  body  models.AddAddress  true	"address"
// @Security		Bearer
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/users/profile/address/add [post]
func (i *UserHandler) AddAddress(c *gin.Context) {

	userID, err := helper.GetUserID(c)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "Could not get userID", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	var address models.AddAddress
	if err := c.BindJSON(&address); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	if err := i.userUseCase.AddAddress(int(userID), address); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "Could not add the address", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Successfully added address", nil, nil)
	c.JSON(http.StatusOK, successRes)

}

// @Summary		Change Password
// @Description	user can change their password
// @Tags			User
// @Accept			json
// @Produce		    json
// @Param			changepassword  body  models.ChangePassword  true	"changepassword"
// @Security		Bearer
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/users/profile/security/change-password [patch]
func (i *UserHandler) ChangePassword(c *gin.Context) {

	id, err := helper.GetUserID(c)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "Could not get userID", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	var ChangePassword models.ChangePassword
	if err := c.BindJSON(&ChangePassword); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	if err := i.userUseCase.ChangePassword(id, ChangePassword.Oldpassword, ChangePassword.Password, ChangePassword.Repassword); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "Could not change the password", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "password changed Successfully ", nil, nil)
	c.JSON(http.StatusOK, successRes)

}

// @Summary		Edit User
// @Description	user can change their Details
// @Tags			User
// @Accept			json
// @Produce		    json
// @Param			userData  body  models.EditUser true	"edit-user"
// @Security		Bearer
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/users/profile/edit [patch]
func (i *UserHandler) EditUser(c *gin.Context) {

	id, err := helper.GetUserID(c)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "Could not get userID", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	var userData models.EditUser
	if err := c.BindJSON(&userData); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	if err := i.userUseCase.EditUser(id, userData); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "Could not change the User Details", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Successfully changed the User Data", nil, nil)
	c.JSON(http.StatusOK, successRes)

}

// @Summary		Get Addresses
// @Description	user can get all their addresses
// @Tags			User
// @Accept			json
// @Produce		    json
// @Security		Bearer
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/users/profile/address [get]
func (i *UserHandler) GetAddresses(c *gin.Context) {
	id, err := helper.GetUserID(c)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "Could not get userID", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	addresses, err := i.userUseCase.GetAddresses(id)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "could not retrieve records", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "Successfully got all records", addresses, nil)
	c.JSON(http.StatusOK, successRes)
}

// @Summary		Get Cart
// @Description	user can view their cart details
// @Tags			User
// @Produce		    json
// @Security		Bearer
// @Param			page	query  string 	true	"page"
// @Param			limit	query  string 	true	"limit"
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/users/cart [get]
func (i *UserHandler) GetCart(c *gin.Context) {
	pageStr := c.Query("page")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "page number not in right format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	limitStr := c.Query("limit")
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "limit number not in right format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	id, err := helper.GetUserID(c)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "Could not get userID", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	products, err := i.userUseCase.GetCart(id, page, limit)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "could not retrieve cart", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "Successfully got all products in cart", products, nil)
	c.JSON(http.StatusOK, successRes)
}

// @Summary		Get User Details
// @Description	user can get all their details
// @Tags			User
// @Accept			json
// @Produce		    json
// @Security		Bearer
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/users/profile/details [get]
func (i *UserHandler) GetUserDetails(c *gin.Context) {
	id, err := helper.GetUserID(c)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "Could not get userID", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	details, err := i.userUseCase.GetUserDetails(id)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "could not retrieve records", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "Successfully got all records", details, nil)
	c.JSON(http.StatusOK, successRes)
}

// Login is a handler for user login
// @Summary		User Login
// @Description	user can log in by giving their details
// @Tags			User
// @Accept			json
// @Produce		    json
// @Param			login  body  models.UserLogin  true	"login"
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/users/login [post]
func (i *UserHandler) Login(c *gin.Context) {
	var user models.UserLogin

	if err := c.BindJSON(&user); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	userDetails, err := i.userUseCase.Login(user)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "User could not be logged in", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "User successfully logged in", userDetails, nil)
	//c.SetCookie("Authorization", userDetails.Token, 3600, "/", "teeverse.online", true, false)
	c.SetCookie("Authorization", userDetails.Token, 3600, "", "", true, false)
	c.JSON(http.StatusOK, successRes)
}

// @Summary		Remove from Cart
// @Description	user can remove products from their cart
// @Tags			User
// @Produce		    json
// @Param			inventory	query	string	true	"inventory id"
// @Security		Bearer
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/users/cart/remove [delete]
func (i *UserHandler) RemoveFromCart(c *gin.Context) {
	id, err := helper.GetUserID(c)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "Could not get userID", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	cartID, err := i.userUseCase.GetCartID(id)
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

	if err := i.userUseCase.RemoveFromCart(cartID, inv); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "could not retrieve cart", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Successfully Removed product from cart", nil, nil)
	c.JSON(http.StatusOK, successRes)
}

// Signup is a handler for user Registration
// @Summary		User Signup
// @Description	user can signup by giving their details
// @Tags			User
// @Accept			json
// @Produce		    json
// @Param			signup  body  models.UserDetails  true	"signup"
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/users/signup [post]
func (i *UserHandler) SignUp(c *gin.Context) {

	var user models.UserDetails
	// bind the user details to the struct
	if err := c.BindJSON(&user); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	// business logic goes inside this function
	userCreated, err := i.userUseCase.SignUp(user)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "User could not signed up", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusCreated, "User successfully signed up", userCreated, nil)
	c.JSON(http.StatusCreated, successRes)

}

// @Summary		Add quantity in cart by one
// @Description	user can add 1 quantity of product to their cart
// @Tags			User
// @Accept			json
// @Produce		    json
// @Param			inventory	query	string	true	"inventory id"
// @Security		Bearer
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/users/cart/updateQuantity/plus [put]
func (i *UserHandler) UpdateQuantityAdd(c *gin.Context) {
	id, err := helper.GetUserID(c)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "Could not get userID", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	cartID, err := i.userUseCase.GetCartID(id)
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

	if err := i.userUseCase.UpdateQuantityAdd(cartID, inv); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "could not Add the quantity", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Successfully added quantity", nil, nil)
	c.JSON(http.StatusOK, successRes)
}

// @Summary		Subtract quantity in cart by one
// @Description	user can subtract 1 quantity of product from their cart
// @Tags			User
// @Accept			json
// @Produce		    json
// @Param			inventory	query	string	true	"inventory id"
// @Security		Bearer
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/users/cart/updateQuantity/minus [put]
func (i *UserHandler) UpdateQuantityLess(c *gin.Context) {
	id, err := helper.GetUserID(c)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "Could not get userID", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	cartID, err := i.userUseCase.GetCartID(id)
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

	if err := i.userUseCase.UpdateQuantityLess(cartID, inv); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "could not  subtract quantity", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Successfully subtracted quantity", nil, nil)
	c.JSON(http.StatusOK, successRes)
}

// @Summary		Get Wallet
// @Description	user can get wallet details and history
// @Tags			User
// @Accept			json
// @Produce		    json
// @Param			page	query  string 	true	"page"
// @Param			limit	query  string 	true	"limit"// @Security		Bearer
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/users/profile/wallet [get]
func (i *UserHandler) GetWallet(c *gin.Context) {
	pageStr := c.Query("page")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "page number not in right format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	limitStr := c.Query("limit")
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "limit number not in right format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	id, err := helper.GetUserID(c)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "Could not get userID", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	wallet, err := i.userUseCase.GetWallet(id, page, limit)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "could not retrieve wallet", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "Successfully got wallet", wallet, nil)
	c.JSON(http.StatusOK, successRes)
}

// Login is a handler for user login
// @Summary		User Logout
// @Description	user can log in by giving their details
// @Tags			User
// @Accept			json
// @Produce		    json
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/users/logout [post]
func (i *UserHandler) Logout(c *gin.Context) {

	successRes := response.ClientResponse(http.StatusOK, "User successfully logged out", nil, nil)
	//c.SetCookie("Authorization", userDetails.Token, 3600, "/", "teeverse.online", true, false)
	c.SetCookie("Authorization", "", -1, "", "", true, false)
	c.JSON(http.StatusOK, successRes)
}
