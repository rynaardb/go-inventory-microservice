/*
 * @File: models.error.go
 * @Description: Defines Error information will be returned to the clients
 * @Author: Rynaard Burger (rynaardb.com)
 */

package models

// Error defines the response error
type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
