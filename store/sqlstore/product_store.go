package store

import "github.com/carmo-evan/holly/model"

// SQLProductStore is a wrapper with Product methods for the SqlStore
type SQLProductStore struct {
	SQLStore
}

// Save commits p to the database
func (s *SQLProductStore) Save(p *model.Product) (*model.Product, error) {
	return p, nil
}

// Update saves changes made to p to the database
func (s *SQLProductStore) Update(p *model.Product) (*model.Product, error) {
	return p, nil
}

// Get return the product with the provided id
func (s *SQLProductStore) Get(id string) (*model.Product, error) {
	p := &model.Product{}
	return p, nil
}
