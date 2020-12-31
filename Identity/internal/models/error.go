/*
 * @File: models.error.go
 * @Description: Defines Error information will be returned to the clients
 * @Author: Yuda Fatah (yudafatah@gmail.com)
 */
package models

// Error defines the response error
type Error struct {
	Code    int    `json:"code" example:"-1"`
	Message string `json:"message" example:"Error message"`
}