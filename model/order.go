package model

import (
	"time"

	"gopkg.in/gorp.v1"
)

// Order represents a collection of order items
type Order struct {
	OrderID      string `json:"orderID" db:"order_id"`
	CreatedAt    int64  `json:"createdAt" db:"created_at"`
	UpdatedAt    int64  `json:"updatedAt" db:"updated_at"`
	DeletedAt    int64  `json:"deletedAt" db:"deleted_at"`
	Total        int64  `json:"total" db:"total"`
	OrderItemIDs string `json:"orderItemIDs" db:"order_item_ids"`
}

// implement the PreInsert and PreUpdate hooks
func (o *Order) PreInsert(s gorp.SqlExecutor) error {
	o.CreatedAt = time.Now().UnixNano()
	o.UpdatedAt = o.CreatedAt
	return nil
}

func (o *Order) PreUpdate(s gorp.SqlExecutor) error {
	o.UpdatedAt = time.Now().UnixNano()
	return nil
}
