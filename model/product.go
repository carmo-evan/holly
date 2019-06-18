package model

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
