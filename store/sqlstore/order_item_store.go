package sqlstore

import (
	"github.com/carmo-evan/holly/model"
	uuid "github.com/satori/go.uuid"
)

// SQLOrderItemStore is a wrapper with OrderItem methods for the SqlStore
type SQLOrderItemStore struct {
	SQLStore *SQLStore
}

// Insert commits new p to the database
func (so *SQLOrderItemStore) Insert(oi *model.OrderItem) (*model.OrderItem, error) {
	oi.OrderItemID = uuid.NewV4().String()
	err := so.SQLStore.Tx.Insert(oi)
	return oi, err
}

// Update saves changes made to p to the database
func (so *SQLOrderItemStore) Update(oi *model.OrderItem) (*model.OrderItem, error) {
	_, err := so.SQLStore.Tx.Update(oi)
	return oi, err
}

// Get returns the product with the provided id
func (so *SQLOrderItemStore) Get(id string) (*model.OrderItem, error) {
	oi, err := so.SQLStore.Tx.Get(model.OrderItem{}, id)
	return oi.(*model.OrderItem), err
}

// CreateTable executes a gorp mapping for model.OrderItem
func (so *SQLOrderItemStore) CreateTable() error {
	// add a table, setting the table name to 'products' and
	// specifying that the Id property is an auto incrementing PK
	so.SQLStore.dbMap.AddTableWithName(model.OrderItem{}, "order_items").SetKeys(false, "order_item_id")

	// create the table. in a production system you'd generally
	// use a migration tool, or create the tables via scripts
	err := so.SQLStore.dbMap.CreateTablesIfNotExists()
	return err
}
