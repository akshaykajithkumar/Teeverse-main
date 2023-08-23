package handler

import (
	"Teeverse/pkg/helper"
	services "Teeverse/pkg/usecase/interface"
	"Teeverse/pkg/utils/response"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PaymentHandler struct {
	paymentUseCase services.PaymentUseCase
}

func NewPaymentHandler(useCase services.PaymentUseCase) *PaymentHandler {
	return &PaymentHandler{
		paymentUseCase: useCase,
	}
}

// @Summary		Add new payment method
// @Description	admin can add a new payment method
// @Tags			Admin
// @Produce		    json
// @Param			paymentMethod	query  string 	true	"Payment Method"
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/admin/paymentmethods/add [post]
func (p *PaymentHandler) AddNewPaymentMethod(c *gin.Context) {
	method := c.Query("paymentMethod")
	if err := p.paymentUseCase.AddNewPaymentMethod(method); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "Could not add new payment method", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Successfully added new payment method", nil, nil)
	c.JSON(http.StatusOK, successRes)
}

// @Summary		Remove payment method
// @Description	admin can remove a  payment method
// @Tags			Admin
// @Produce		    json
// @Param			paymentMethodID	query  int 	true	"Payment Method ID"
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/admin/paymentmethods/remove [delete]
func (p *PaymentHandler) RemovePaymentMethod(c *gin.Context) {
	methodstr := c.Query("paymentMethodID")
	method, err := strconv.Atoi(methodstr)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "Could not convert paymentid strign to int", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	if err := p.paymentUseCase.RemovePaymentMethod(method); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "Could not remove payment method", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Successfully removed payment method", nil, nil)
	c.JSON(http.StatusOK, successRes)
}

// @Summary		Get payment methods
// @Description	admin can get all  payment methods
// @Tags			Admin
// @Produce		    json
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/admin/paymentmethods [get]
func (p *PaymentHandler) GetPaymentMethods(c *gin.Context) {
	paymentMethods, err := p.paymentUseCase.GetPaymentMethods()
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "Could not get payment methods", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Successfully collected payment methods", paymentMethods, nil)
	c.JSON(http.StatusOK, successRes)
}

func (p *PaymentHandler) MakePaymentRazorPay(c *gin.Context) {

	orderID := c.Query("id")
	userID, err := helper.GetUserID(c)
	fmt.Println("====", userID, orderID)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "Could not get userID", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	orderDetail, err := p.paymentUseCase.MakePaymentRazorPay(orderID, userID)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusInternalServerError, "could not generate order details", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}
	fmt.Println(orderDetail)
	c.HTML(http.StatusOK, "razorpay.html", orderDetail)
}

func (p *PaymentHandler) VerifyPayment(c *gin.Context) {

	orderID := c.Query("order_id")
	paymentID := c.Query("payment_id")
	razorID := c.Query("razor_id")
	fmt.Println(orderID, paymentID, razorID)
	err := p.paymentUseCase.VerifyPayment(paymentID, razorID, orderID)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusInternalServerError, "could not update payment details", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}

	//clear cart
	successRes := response.ClientResponse(http.StatusOK, "Successfully updated payment details", nil, nil)
	c.JSON(http.StatusOK, successRes)

}
