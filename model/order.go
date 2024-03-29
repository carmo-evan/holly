package model

// Order represents a collection of order items
type Order struct {
	OrderID   string `json:"orderID" db:"order_id"`
	UserID    string `json:"userID" db:"user_id"`
	CreatedAt int64  `json:"createdAt" db:"created_at"`
	UpdatedAt int64  `json:"updatedAt" db:"updated_at"`
	DeletedAt int64  `json:"deletedAt" db:"deleted_at"`
	Total     int64  `json:"total" db:"total"`
}
