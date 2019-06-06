package store

import (
	"github.com/carmo-evan/holly/model"
)

// Store outlines methods that return the specific store for each entity in the app
type Store interface {
	Product() ProductStore
}

//ProductStore outlines CRUD methods related to Products
type ProductStore interface {
	Save(product *model.Product) (*model.Product, error)
	Update(product *model.Product) (*model.Product, error)
	Get(productID string) (*model.Product, error)
}
