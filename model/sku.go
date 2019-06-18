package model

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
