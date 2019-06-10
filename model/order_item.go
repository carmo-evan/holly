package model

import (
	"time"

	"gopkg.in/gorp.v1"
)

// Order Item represents a quantity of a certain sku placed in an order
type OrderItem struct {
	OrderItemID string `json:"orderItemID" db:"order_item_id"`
	OrderID     string `json:"orderID" db:"order_id"`
	SkuID       string `json:"skuID" db:"sku_id"`
	CreatedAt   int64  `json:"createdAt" db:"created_at"`
	UpdatedAt   int64  `json:"updatedAt" db:"updated_at"`
	DeletedAt   int64  `json:"deletedAt" db:"deleted_at"`
	Price       int64  `json:"total" db:"total"`
	Quantity    int64  `json:"quantity" db:"quantity"`
}

// implement the PreInsert and PreUpdate hooks
func (oi *OrderItem) PreInsert(s gorp.SqlExecutor) error {
	oi.CreatedAt = time.Now().UnixNano()
	oi.UpdatedAt = oi.CreatedAt
	return nil
}

func (oi *OrderItem) PreUpdate(s gorp.SqlExecutor) error {
	oi.UpdatedAt = time.Now().UnixNano()
	return nil
}
