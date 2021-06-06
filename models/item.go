/*
 * @File: models.item.go
 * @Description: Defines the Item model
 * @Author: Rynaard Burger (rynaardb.com)
 */

package models

import (
	"errors"

	"rynaardb.com/inventory-service/common"
)

// Item model
type Item struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	CategoryID string `json:"categoryID"`
	Active     bool   `json:"active"`
	SoldByType string `json:"soldByType"`
	Price      int64  `json:"price"`
	Cost       int64  `json:"cost"`
	Sku        string `json:"sku"`
	Barcode    string `json:"barcode"`
	TrackStock bool   `json:"trackStock"`
	Color      string `json:"color"`
	Image      string `json:"image"`
}

// AddItem request model
type ItemRequest struct {
	Name       string `json:"name"`
	CategoryID string `json:"categoryID"`
	Active     bool   `json:"active"`
	SoldByType string `json:"soldByType"`
	Price      int64  `json:"price"`
	Cost       int64  `json:"cost"`
	Sku        string `json:"sku"`
	Barcode    string `json:"barcode"`
	TrackStock bool   `json:"trackStock"`
	Color      string `json:"color"`
	Image      string `json:"image"`
}

// Validate item
func (i ItemRequest) Validate() error {
	switch {
	case len(i.Name) == 0:
		return errors.New(common.ErrItmNameEmpty)
	case len(i.CategoryID) == 0:
		return errors.New(common.ErrCatIDEmpty)
	default:
		return nil
	}
}
