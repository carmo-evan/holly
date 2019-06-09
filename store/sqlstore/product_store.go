package sqlstore

import (
	"fmt"

	"github.com/carmo-evan/holly/model"
	uuid "github.com/satori/go.uuid"
)

// SQLProductStore is a wrapper with Product methods for the SqlStore
type SQLProductStore struct {
	SQLStore *SQLStore
}

// Insert commits new p to the database
func (sp *SQLProductStore) Insert(p *model.Product) (*model.Product, error) {
	p.ProductID = uuid.NewV4().String()

	err := sp.SQLStore.Tx.Insert(p)
	if err != nil {
		err = fmt.Errorf("[SQLProductStore] error in calling Insert: %v", err)
	}

	return p, err
}

// Update saves changes made to p to the database
func (sp *SQLProductStore) Update(p *model.Product) (*model.Product, error) {
	_, err := sp.SQLStore.Tx.Update(p)
	return p, err
}

// Get returns the product with the provided id
func (sp *SQLProductStore) Get(id string) (*model.Product, error) {
	p, err := sp.SQLStore.Tx.Get(model.Product{}, id)
	return p.(*model.Product), err
}

// CreateTable executes a gorp mapping for model.Product
func (sp *SQLProductStore) CreateTable() error {
	// add a table, setting the table name to 'products' and
	// specifying that the Id property is an auto incrementing PK
	sp.SQLStore.dbMap.AddTableWithName(model.Product{}, "products").SetKeys(false, "product_id")

	// create the table. in a production system you'd generally
	// use a migration tool, or create the tables via scripts
	err := sp.SQLStore.dbMap.CreateTablesIfNotExists()
	return err
}
