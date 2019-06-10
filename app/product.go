package app

import "github.com/carmo-evan/holly/model"

// CreateProduct takes a name, displayName and Description and creates a product witha default SKU
func CreateProduct(price int64, name, displayName, description string) (*model.Product, error) {
	p := &model.Product{Name: name, DisplayName: displayName, Description: description}

	return p, nil
}
