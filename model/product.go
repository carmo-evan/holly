package model

import "github.com/jinzhu/gorm"

// Product represents the common properties of a collection of skus
type Product struct {
	gorm.Model
	DisplayName string `json:"display_name"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
