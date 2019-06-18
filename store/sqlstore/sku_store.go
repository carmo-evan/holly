package sqlstore

import (
	"fmt"
	"time"

	"github.com/carmo-evan/holly/model"
	uuid "github.com/satori/go.uuid"
	"gopkg.in/gorp.v1"
)

// SQLSKUStore is a wrapper with SKU methods for the SqlStore
type SQLSKUStore struct {
	SQLStore *SQLStore
	tableMap *gorp.TableMap
}

// Insert commits new p to the database
func (sks *SQLSKUStore) Insert(s *model.SKU) (*model.SKU, error) {
	s.SKUID = uuid.NewV4().String()
	s.CreatedAt = time.Now().UnixNano()
	s.UpdatedAt = s.CreatedAt
	err := sks.SQLStore.Tx.Insert(s)
	if err != nil {
		err = fmt.Errorf("[SQLKUStore] error in calling Insert: %v", err)
	}

	return s, err
}

// Update saves changes made to s to the database
func (sks *SQLSKUStore) Update(s *model.SKU) (*model.SKU, error) {
	s.UpdatedAt = time.Now().UnixNano()
	_, err := sks.SQLStore.Tx.Update(s)
	return s, err
}

// Get returns the product with the provided id
func (sks *SQLSKUStore) Get(id string) (*model.SKU, error) {
	s, err := sks.SQLStore.Tx.Get(model.SKU{}, id)
	return s.(*model.SKU), err
}

// CreateTable executes a gorp mapping for model.Product
func (sks *SQLSKUStore) CreateTable() error {
	sks.SQLStore.dbMap.AddTableWithName(model.SKU{}, "skus").SetKeys(false, "sku_id")
	err := sks.SQLStore.dbMap.CreateTablesIfNotExists()
	return err
}
