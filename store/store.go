package store

import (
	"github.com/carmo-evan/holly/model"
)

type DbFlavor int

const (
	Postgres DbFlavor = 0
)

// Store outlines methods that return the specific store for each entity in the app
type Store interface {
	Product() ProductStore
	SKU() SKUStore
	Order() OrderStore
	OrderItem() OrderItemStore
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

//OrderStore outlines CRUD methods related to Orders
type OrderStore interface {
	Insert(product *model.Order) (*model.Order, error)
	Update(product *model.Order) (*model.Order, error)
	Get(OrderID string) (*model.Order, error)
	CreateTable() error
}

//OrderItemStore outlines CRUD methods related to Order Items
type OrderItemStore interface {
	Insert(product *model.OrderItem) (*model.OrderItem, error)
	Update(product *model.OrderItem) (*model.OrderItem, error)
	Get(OrderItemID string) (*model.OrderItem, error)
	CreateTable() error
}
