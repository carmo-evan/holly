package app

import "github.com/carmo-evan/holly/model"

// CreateProduct takes a name, displayName and Description and creates a product
func (a *App) CreateProduct(name, displayName, description string) (*model.Product, error) {
	p := &model.Product{Name: name, DisplayName: displayName, Description: description}
	return a.Store.Product().Insert(p)
}

// CreateSKU takes a price, name skuCode, Description and productID and creates a SkU
func (a *App) CreateSKU(price int64, name, description, skuCode, productID string) (*model.SKU, error) {
	s := &model.SKU{Name: name, Description: description, Price: price, SKUCode: skuCode, ProductID: productID}
	return a.Store.SKU().Insert(s)
}

// CreateProductWithDefaultSku takes a price, name, displayName and Description and creates a product with a default SKU
func (a *App) CreateProductWithDefaultSku(price int64, name, displayName, description string) (*model.Product, *model.SKU, error) {
	p, err := a.CreateProduct(name, displayName, description)
	if err != nil {
		return p, nil, err
	}
	s, err := a.CreateSKU(price, name, description, name, p.ProductID)
	p.DefaultSKUID = s.SKUID
	p, err = a.Store.Product().Update(p)
	return p, s, err
}
