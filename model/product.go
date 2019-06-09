package model

import (
	"time"

	"gopkg.in/gorp.v1"
)

// Product represents the common properties of a collection of skus
type Product struct {
	ProductID    string `json:"productID" db:"product_id"`
	CreatedAt    int64  `json:"createdAt" db:"created_at"`
	UpdatedAt    int64  `json:"updatedAt" db:"updated_at"`
	DeletedAt    int64  `json:"deletedAt" db:"deleted_at"`
	DisplayName  string `json:"displayName" db:"display_name"`
	Name         string `json:"name" db:"name"`
	Description  string `json:"description" db:"description"`
	DefaultSKUID string `json:"defaultSkuID" db:"default_sku_id"`
	SKUIDs       string `json:"skuIDs" db:"sku_ids"`
}

// PreInsert guarantees timestamps are set and that a defaultSKU exists
func (p *Product) PreInsert(s gorp.SqlExecutor) error {
	p.CreatedAt = time.Now().UnixNano()
	p.UpdatedAt = p.CreatedAt
	return nil
}

func (p *Product) PreUpdate(s gorp.SqlExecutor) error {
	p.UpdatedAt = time.Now().UnixNano()
	return nil
}
