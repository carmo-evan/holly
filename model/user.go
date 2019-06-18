package model

// User represents a customer/admin account
type User struct {
	UserID    string `json:"userID" db:"user_id"`
	CreatedAt int64  `json:"createdAt" db:"created_at"`
	UpdatedAt int64  `json:"updatedAt" db:"updated_at"`
	DeletedAt int64  `json:"deletedAt" db:"deleted_at"`
	Email     string `json:"email" db:"email"`
}
