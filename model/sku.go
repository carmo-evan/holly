package model

import (
	"time"

	"gopkg.in/gorp.v1"
)

// SKU represents a unique combination of product options
type SKU struct {
	SKUID       string `json:"skuID" db:"sku_id"`
	CreatedAt   int64  `json:"createdAt" db:"created_at"`
	UpdatedAt   int64  `json:"updatedAt" db:"updated_at"`
	DeletedAt   int64  `json:"deletedAt" db:"deleted_at"`
	Price       int64  `json:"price" db:"price"`
	SKUCode     string `json:"skuCode" db:"sku_code"`
	Name        string `json:"name" db:"name"`
	Description string `json:"description" db:"description"`
	ProductID   string `json:"productID" db:"productID"`
}

// implement the PreInsert and PreUpdate hooks
func (sku *SKU) PreInsert(s gorp.SqlExecutor) error {
	sku.CreatedAt = time.Now().UnixNano()
	sku.UpdatedAt = sku.CreatedAt
	return nil
}

func (sku *SKU) PreUpdate(s gorp.SqlExecutor) error {
	sku.UpdatedAt = time.Now().UnixNano()
	return nil
}
