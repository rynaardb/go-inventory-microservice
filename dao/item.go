/*
 * @File: dao.item.go
 * @Description: Implements Item CRUD functions for Postgres
 * @Author: Rynaard Burger (rynaardb.com)
 */

package dao

import (
	"gorm.io/gorm/clause"
	"rynaardb.com/inventory-service/databases"
	"rynaardb.com/inventory-service/models"
)

// Item manages Item CRUD
type Item struct {
}

// GetAll gets the list of Items
func (i *Item) GetAll() ([]models.Item, error) {
	var items []models.Item
	db := databases.Database.DBConn
	err := db.Debug().Preload(clause.Associations).Find(&items).Error
	return items, err
}

// GetByID finds a Item by ID
func (i *Item) GetByID(id string) (models.Item, error) {
	var item models.Item
	db := databases.Database.DBConn
	err := db.Debug().Preload(clause.Associations).Where("id = ?", id).First(&item).Error
	return item, err
}

// DeleteByID finds and deletes a Item by ID
func (i *Item) DeleteByID(id string) error {
	var item models.Item
	db := databases.Database.DBConn
	err := db.Debug().Where("id = ?", id).Delete(&item).Error
	return err
}

// Create a new Item
func (i *Item) Create(item models.Item) error {
	db := databases.Database.DBConn
	err := db.Debug().Create(&item).Error
	return err
}

// Update an existing Item
func (i *Item) Update(item models.Item) error {
	db := databases.Database.DBConn
	err := db.Debug().Save(&item).Error
	return err
}

// Delete an existing Item
func (i *Item) Delete(item models.Item) error {
	db := databases.Database.DBConn
	var err error
	db.Select(clause.Associations).Delete(&item)
	err = db.Debug().Delete(&item).Error
	return err
}
