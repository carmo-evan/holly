package sqlstore

import (
	"github.com/carmo-evan/holly/model"
	uuid "github.com/satori/go.uuid"
)

// SQLOrderStore is a wrapper with Order methods for the SqlStore
type SQLOrderStore struct {
	SQLStore *SQLStore
}

// Insert commits new p to the database
func (so *SQLOrderStore) Insert(o *model.Order) (*model.Order, error) {
	o.OrderID = uuid.NewV4().String()
	err := so.SQLStore.Tx.Insert(o)
	return o, err
}

// Update saves changes made to p to the database
func (so *SQLOrderStore) Update(o *model.Order) (*model.Order, error) {
	_, err := so.SQLStore.Tx.Update(o)
	return o, err
}

// Get returns the product with the provided id
func (so *SQLOrderStore) Get(id string) (*model.Order, error) {
	o, err := so.SQLStore.Tx.Get(model.Order{}, id)
	return o.(*model.Order), err
}

// CreateTable executes a gorp mapping for model.Order
func (so *SQLOrderStore) CreateTable() error {
	// add a table, setting the table name to 'products' and
	// specifying that the Id property is an auto incrementing PK
	so.SQLStore.dbMap.AddTableWithName(model.Order{}, "orders").SetKeys(false, "order_id")

	// create the table. in a production system you'd generally
	// use a migration tool, or create the tables via scripts
	err := so.SQLStore.dbMap.CreateTablesIfNotExists()
	return err
}
