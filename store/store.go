package store

import (
	"github.com/carmo-evan/holly/model"
)

// Store outlines methods that return the specific store for each entity in the app
type Store interface {
	Product() ProductStore
	SKU() SKUStore
	Close() error
	Commit() error
}

//ProductStore outlines CRUD methods related to Products
type ProductStore interface {
	Insert(product *model.Product) (*model.Product, error)
	Update(product *model.Product) (*model.Product, error)
	Get(productID string) (*model.Product, error)
	CreateTable() error
}

//SKUStore outlines CRUD methods related to SKUs
type SKUStore interface {
	Insert(product *model.SKU) (*model.SKU, error)
	Update(product *model.SKU) (*model.SKU, error)
	Get(SKUID string) (*model.SKU, error)
	CreateTable() error
}
