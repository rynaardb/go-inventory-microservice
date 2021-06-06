package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"rynaardb.com/inventory-service/common"
	"rynaardb.com/inventory-service/dao"
	"rynaardb.com/inventory-service/models"
)

// Item manages
type Item struct {
	itemDAO dao.Item
}

// ListItems godoc
// @Summary List all existing items
// @Description List all existing items
// @Tags items
// @Accept  json
// @Produce  json
// @Param Authorization header string true "JWT"
// @Failure 500 {object} models.Error
// @Success 200 {array} models.Item
// @Router /items [get]
func (i *Item) ListItems(ctx *gin.Context) {
	var items []models.Item
	var err error

	items, err = i.itemDAO.GetAll()

	if err == nil {
		ctx.JSON(http.StatusOK, items)
	} else {
		ctx.JSON(http.StatusInternalServerError, models.Error{Code: common.StatusCodeUnknown, Message: err.Error()})
		fmt.Println("[ERROR]: ", err)
	}
}

// GetItemByID godoc
// @Summary Get a item by ID
// @Description Get a item by ID
// @Tags items
// @Accept  json
// @Produce  json
// @Param Authorization header string true "JWT"
// @Param id path string true "Item ID"
// @Failure 500 {object} models.Error
// @Success 200 {object} models.Item
// @Router /items/{id} [get]
func (i *Item) GetItemByID(ctx *gin.Context) {
	var item models.Item
	var err error
	id := ctx.Params.ByName("id")

	item, err = i.itemDAO.GetByID(id)

	if err == nil {
		ctx.JSON(http.StatusOK, item)
	} else {
		ctx.JSON(http.StatusInternalServerError, models.Error{Code: common.StatusCodeUnknown, Message: err.Error()})
		fmt.Println("[ERROR]: ", err)
	}
}

// GetItemByParams godoc
// @Summary Get a item by ID parameter
// @Description Get a item by ID parameter
// @Tags items
// @Accept  json
// @Produce  json
// @Param Authorization header string true "JWT"
// @Param id query string true "Item ID"
// @Failure 500 {object} models.Error
// @Success 200 {object} models.Item
// @Router /items [get]
func (i *Item) GetItemByParams(ctx *gin.Context) {
	var item models.Item
	var err error
	id := ctx.Request.URL.Query()["id"][0]

	item, err = i.itemDAO.GetByID(id)

	if err == nil {
		ctx.JSON(http.StatusOK, item)
	} else {
		ctx.JSON(http.StatusInternalServerError, models.Error{Code: common.StatusCodeUnknown, Message: err.Error()})
		fmt.Println("[ERROR]: ", err)
	}
}

// AddItem godoc
// @Summary Add a new item
// @Description Add a new item
// @Tags items
// @Accept  json
// @Produce  json
// @Param Authorization header string true "JWT"
// @Param item body models.ItemRequest true "Item request object"
// @Failure 500 {object} models.Error
// @Failure 400 {object} models.Error
// @Success 200 {object} models.Message
// @Router /items [post]
func (i *Item) AddItem(ctx *gin.Context) {
	var addItem models.ItemRequest

	if err := ctx.BindJSON(&addItem); err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Error{Code: common.StatusCodeUnknown, Message: err.Error()})
		fmt.Println("[ERROR]: ", err)
		return
	}

	if err := addItem.Validate(); err != nil {
		ctx.JSON(http.StatusBadRequest, models.Error{Code: common.StatusCodeUnknown, Message: err.Error()})
		fmt.Println("[ERROR]: ", err)
		return
	}

	itemID, _ := uuid.NewV4()

	item := models.Item{
		ID:         itemID.String(),
		Name:       addItem.Name,
		CategoryID: addItem.CategoryID,
		Active:     addItem.Active,
		SoldByType: addItem.SoldByType,
		Price:      addItem.Price,
		Cost:       addItem.Cost,
		Sku:        addItem.Sku,
		Barcode:    addItem.Barcode,
		TrackStock: addItem.TrackStock,
		Color:      addItem.Color,
		Image:      addItem.Image,
	}

	err := i.itemDAO.Create(item)

	if err == nil {
		ctx.JSON(http.StatusOK, item)
		fmt.Println("[DEBUG]: New item created with ID: ", itemID.String())
	} else {
		ctx.JSON(http.StatusInternalServerError, models.Error{Code: common.StatusCodeUnknown, Message: err.Error()})
		fmt.Println("[ERROR]: ", err)
	}
}

// UpdateItem godoc
// @Summary Update an existing item
// @Description Update an existing item
// @Tags items
// @Accept  json
// @Produce  json
// @Param Authorization header string true "JWT"
// @Param item body models.Item true "Item object"
// @Failure 500 {object} models.Error
// @Success 200 {object} models.Message
// @Router /items [put]
func (i *Item) UpdateItem(ctx *gin.Context) {
	var item models.Item
	if err := ctx.ShouldBindJSON(&item); err != nil {
		ctx.JSON(http.StatusBadRequest, models.Error{Code: common.StatusCodeUnknown, Message: err.Error()})
		fmt.Println("[ERROR]: ", err)
		return
	}

	err := i.itemDAO.Update(item)

	if err == nil {
		ctx.JSON(http.StatusOK, item)
	} else {
		ctx.JSON(http.StatusInternalServerError, models.Error{Code: common.StatusCodeUnknown, Message: err.Error()})
		fmt.Println("[ERROR]: ", err)
	}
}

// DeleteItem godoc
// @Summary Delete a item by ID
// @Description Delete a item by ID
// @Tags items
// @Accept  json
// @Produce  json
// @Param Authorization header string true "JWT"
// @Param item body models.Item true "Item object"
// @Failure 500 {object} models.Error
// @Success 200 {object} models.Message
// @Router /items [delete]
func (i *Item) DeleteItem(ctx *gin.Context) {
	var item models.Item
	if err := ctx.ShouldBindJSON(&item); err != nil {
		ctx.JSON(http.StatusBadRequest, models.Error{Code: common.StatusCodeUnknown, Message: err.Error()})
		fmt.Println("[ERROR]: ", err)
		return
	}

	err := i.itemDAO.Delete(item)

	if err == nil {
		ctx.JSON(http.StatusOK, models.Message{Message: "success"})
	} else {
		ctx.JSON(http.StatusInternalServerError, models.Error{Code: common.StatusCodeUnknown, Message: err.Error()})
		fmt.Println("[ERROR]: ", err)
	}
}
