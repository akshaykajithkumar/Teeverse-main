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
// @Param			description		formData	string	true	"description"
// @Param			price	formData	string	true	"price"
// @Param			stock		formData	string	true	"stock"
// @Param           image      formData     file   true   "image"
// @Security		Bearer
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/admin/inventories/add [post]
func (i *InventoryHandler) AddInventory(c *gin.Context) {
	//change
	var inventory models.Inventory
	categoryID, err := strconv.Atoi(c.Request.FormValue("category_id"))
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "form file error", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	product_name := c.Request.FormValue("product_name")
	description := c.Request.FormValue("description")
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

	image, err := c.FormFile("image")
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "retrieving image from form error", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	uploadPath := "./assets/" + image.Filename

	// Save the uploaded image to the specified path
	if err := c.SaveUploadedFile(image, uploadPath); err != nil {
		errorRes := response.ClientResponse(http.StatusInternalServerError, "Could not save uploaded image", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}
	imageURL := "https://teeverse.online/assets/" + image.Filename
	inventory.CategoryID = categoryID
	inventory.ProductName = product_name
	inventory.Description = description
	inventory.Price = price
	inventory.Stock = stock
	inventory.Image = imageURL

	InventoryResponse, err := i.InventoryUseCase.AddInventory(inventory, image.Filename)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "Could not add the Inventory", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Successfully added Inventory", InventoryResponse, nil)
	c.JSON(http.StatusOK, successRes)
}

// @Summary		Add image to an Inventory
// @Description	Admin can add new image to product
// @Tags			Admin
// @Accept			multipart/form-data
// @Produce		    json
// @Param			product_id	formData	string	true	"product_id"
// @Param           image      formData     file   true   "image"
// @Security		Bearer
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/admin/inventories/add-image [post]
func (i *InventoryHandler) AddImage(c *gin.Context) {

	product_id, err := strconv.Atoi(c.Request.FormValue("product_id"))
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "form file error", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	image, err := c.FormFile("image")
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "retrieving image from form error", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	uploadPath := "./assets/" + image.Filename

	// Save the uploaded image to the specified path
	if err := c.SaveUploadedFile(image, uploadPath); err != nil {
		errorRes := response.ClientResponse(http.StatusInternalServerError, "Could not save uploaded image", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}
	imageURL := "https://teeverse.online/assets/" + image.Filename
	InventoryResponse, err := i.InventoryUseCase.AddImage(product_id, imageURL)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "Could not add the Inventory image", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Successfully added Inventory image", InventoryResponse, nil)
	c.JSON(http.StatusOK, successRes)
}

// @Summary		Delete Inventory image
// @Description	Admin can delete a product image
// @Tags			Admin
// @Accept			json
// @Produce		    json
// @Param			product_id	query	string	true	"product_id"
// @Param			image_id	query	string	true	"image_id"
// @Security		Bearer
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/admin/inventories/delete-image [delete]
func (i *InventoryHandler) DeleteImage(c *gin.Context) {

	inventoryID, err := strconv.Atoi(c.Query("product_id"))
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "image ID   not in right format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	imageID, err := strconv.Atoi(c.Query("image_id"))
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "image ID   not in right format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	err = i.InventoryUseCase.DeleteImage(inventoryID, imageID)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "Could not remove the Inventory", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Successfully deleted the inventory image", nil, nil)
	c.JSON(http.StatusOK, successRes)
}

// @Summary		List Productss
// @Description	client can view the list of available products
// @Tags			Admin
// @Accept			json
// @Produce		    json
// @Param			page	query  string 	true	"page"
// @Param			limit	query  string 	true	"limit"
// @Security		Bearer
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/admin/products [get]
func (i *InventoryHandler) AdminListProducts(c *gin.Context) {
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
// @Description	client can search with a key and get the list of products similar to that key
// @Tags			Products
// @Accept		json
// @Produce		json
// @Param		page	query  string 	true	"page"
// @Param		limit	query  string 	true	"limit"
// @Param		searchkey	query  string 	true	"searchkey"
// @Param		sortBY	query  string 	false	"sortBY (asc/desc) - Sort by price in ascending (asc) or descending (desc) order"
// @Success		200	{object}	response.Response{}
// @Failure		400	{object}	response.Response{}
// @Router		/products/search [get]
func (i *InventoryHandler) SearchProducts(c *gin.Context) {
	pageStr := c.Query("page")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "page number not in the right format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	limitStr := c.Query("limit")
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "limit number not in the right format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	searchKey := c.Query("searchkey")
	sortBY := c.Query("sortBY") // Add this line to get the sorting parameter

	results, err := i.InventoryUseCase.SearchProducts(searchKey, page, limit, sortBY)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "could not retrieve the records", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Successfully got all records", results, nil)
	c.JSON(http.StatusOK, successRes)
}

// @Summary		filter Products by category
// @Description	client can filter with a category and get the list of  products in the category
// @Tags			Products
// @Accept			json
// @Produce		    json
// @Param			page	query  string 	true	"page"
// @Param			limit	query  string 	true	"limit"
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/products/category [get]
func (i *InventoryHandler) GetCategories(c *gin.Context) {
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
	// catIDstr := c.Query("catID")
	// catID, err := strconv.Atoi(catIDstr)
	// if err != nil {
	// 	errorRes := response.ClientResponse(http.StatusBadRequest, "category ID not in right format", nil, err.Error())
	// 	c.JSON(http.StatusBadRequest, errorRes)
	// 	return
	// }

	categories, err := i.InventoryUseCase.GetCategories(page, limit)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "could not retrieve records", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "Successfully retrieved the categories", categories, nil)
	c.JSON(http.StatusOK, successRes)
}

// @Summary		Update Stock
// @Description	Admin can update stock of the inventories
// @Tags			Admin
// @Accept			json
// @Produce		    json
// @Param			id	query	string	true	"id"
// @Param			updateinventory	body	models.UpdateInventory	true	"Update Inventory"
// @Security		Bearer
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/admin/inventories/update [patch]
func (i *InventoryHandler) UpdateInventory(c *gin.Context) {
	//change
	inventoryIDstr := c.Query("id")
	invID, err := strconv.Atoi(inventoryIDstr)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "id is not valid", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	var invData models.UpdateInventory

	if err := c.BindJSON(&invData); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	invRes, err := i.InventoryUseCase.UpdateInventory(invID, invData)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "could not update the inventory stock", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Successfully updated the inventory stock", invRes, nil)
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
// @Description	client can view the details of the product
// @Tags			Products
// @Accept			json
// @Produce		    json
// @Param			inventoryID	query	string	true	"Inventory ID"
// @Security		Bearer
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/products/details [get]
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
// @Description	client can view the list of available products
// @Tags			Products
// @Accept			json
// @Produce		    json
// @Param			page	query  string 	true	"page"
// @Param			limit	query  string 	true	"limit"
// @Security		Bearer
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/products [get]
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

// @Summary		List Products based on category
// @Description	client can view the list of available products
// @Tags			Products
// @Accept			json
// @Produce		    json
// @Param			page	query  string 	true	"page"
// @Param			limit	query  string 	true	"limit"
// @Param			catid	query  string 	true	"id"
// @Security		Bearer
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/products/category_details [get]
func (i *InventoryHandler) GetCategoryProducts(c *gin.Context) {
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
	catidStr := c.Query("catid")
	catid, err := strconv.Atoi(catidStr)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "catid number not in right format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	products, err := i.InventoryUseCase.GetCategoryProducts(catid, page, limit)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "could not retrieve records", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "Successfully got all records", products, nil)
	c.JSON(http.StatusOK, successRes)
}
