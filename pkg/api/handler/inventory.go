package handler

import (
	services "Teeverse/pkg/usecase/interface"
	"Teeverse/pkg/utils/models"
	"Teeverse/pkg/utils/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type InventoryHandler struct {
	InventoryUseCase services.InventoryUseCase
}

func NewInventoryHandler(usecase services.InventoryUseCase) *InventoryHandler {
	return &InventoryHandler{
		InventoryUseCase: usecase,
	}
}

// @Summary		Add Inventory
// @Description	Admin can add new  products
// @Tags			Admin
// @Accept			multipart/form-data
// @Produce		    json
// @Param			category_id		formData	string	true	"category_id"
// @Param			product_name	formData	string	true	"product_name"
// @Param			price	formData	string	true	"price"
// @Param			stock		formData	string	true	"stock"
// @Param           image      formData     string   true   "image"
// @Security		Bearer
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/admin/inventories/add [post]
func (i *InventoryHandler) AddInventory(c *gin.Context) {

	var inventory models.Inventory
	categoryID, err := strconv.Atoi(c.Request.FormValue("category_id"))
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "form file error", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	product_name := c.Request.FormValue("product_name")
	p, err := strconv.Atoi(c.Request.FormValue("price"))
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "form file error", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	price := float64(p)
	stock, err := strconv.Atoi(c.Request.FormValue("stock"))
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "form file error", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	image := c.Request.FormValue("image")
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "retrieving image from form error", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	inventory.CategoryID = categoryID
	inventory.ProductName = product_name
	inventory.Price = price
	inventory.Stock = stock
	inventory.Image = image

	InventoryResponse, err := i.InventoryUseCase.AddInventory(inventory, image)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "Could not add the Inventory", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Successfully added Inventory", InventoryResponse, nil)
	c.JSON(http.StatusOK, successRes)

}

// @Summary		Update Stock
// @Description	Admin can update stock of the inventories
// @Tags			Admin
// @Accept			json
// @Produce		    json
// @Param			add-stock	body	models.InventoryUpdate	true	"update stock"
// @Security		Bearer
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/admin/inventories/update [put]
func (i *InventoryHandler) UpdateInventory(c *gin.Context) {

	var p models.InventoryUpdate

	if err := c.BindJSON(&p); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	a, err := i.InventoryUseCase.UpdateInventory(p.Productid, p.Stock)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "could not update the inventory stock", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Successfully updated the inventory stock", a, nil)
	c.JSON(http.StatusOK, successRes)

}

// @Summary		Delete Inventory
// @Description	Admin can delete a product
// @Tags			Admin
// @Accept			json
// @Produce		    json
// @Param			id	query	string	true	"id"
// @Security		Bearer
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/admin/inventories/delete [delete]
func (i *InventoryHandler) DeleteInventory(c *gin.Context) {

	inventoryID := c.Query("id")
	err := i.InventoryUseCase.DeleteInventory(inventoryID)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Successfully deleted the inventory", nil, nil)
	c.JSON(http.StatusOK, successRes)

}

// @Summary		Show Product Details
// @Description	user can view the details of the product
// @Tags			User
// @Accept			json
// @Produce		    json
// @Param			inventoryID	query	string	true	"Inventory ID"
// @Security		Bearer
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/users/home/products/details [get]
func (i *InventoryHandler) ShowIndividualProducts(c *gin.Context) {

	id := c.Query("inventoryID")
	product, err := i.InventoryUseCase.ShowIndividualProducts(id)

	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "path variables in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Product details retrieved successfully", product, nil)
	c.JSON(http.StatusOK, successRes)

}

// @Summary		List Products
// @Description	user can view the list of available products
// @Tags			User
// @Accept			json
// @Produce		    json
// @Param			page	query  string 	true	"page"
// @Param			limit	query  string 	true	"limit"
// @Security		Bearer
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/users/home/products [get]
func (i *InventoryHandler) ListProducts(c *gin.Context) {
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
	products, err := i.InventoryUseCase.ListProducts(page, limit)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "could not retrieve records", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "Successfully got all records", products, nil)
	c.JSON(http.StatusOK, successRes)
}

// @Summary		Search Products
// @Description	user can search with a key and get the list of  products similar to that key
// @Tags			User
// @Accept			json
// @Produce		    json
// @Param			page	query  string 	true	"page"
// @Param			limit	query  string 	true	"limit"
// @Param			searchkey 	query  string 	true	"searchkey"
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/users/search [post]
func (i *InventoryHandler) SearchProducts(c *gin.Context) {
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
	searchkey := c.Query("searchkey")
	results, err := i.InventoryUseCase.SearchProducts(searchkey, page, limit)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "could not retrieve the records", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Successfully got all records", results, nil)
	c.JSON(http.StatusOK, successRes)
}
