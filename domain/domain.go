package domain

import (
	"time"
)

type Order struct {
	ID             string            `json:"id"`
	UserName       string            `json:"user_name"`
	TenantID       int               `bson:"tenant_id,omitempty" json:"tenantid"`
	HubID          int               `bson:"hub_id,omitempty" json:"hubid"`
	SellerID       int               `json:"seller_id"`
	Status         string            `json:"status"`
	AdditionalInfo map[string]string `json:"additional_info"`
	CreatedAt      time.Time         `json:"createdat"`
	UpdatedAt      time.Time         `json:"updatedat"`
	OrderItems     []OrderItem       `json:"order_items"`
}

type OrderItem struct {
	SKU_ID   int `json:"sku_id"`
	Quantity int `json:"quantity"`
}

type CreateOrderRequest struct {
	Path           string            `json:"path"`
	UserID         int               `json:"user_id"`
	TenantID       int               `json:"tenant_id"`
	HubID          int               `json:"hub_id"`
	SellerID       int               `json:"seller_id"`
	AdditionalInfo map[string]string `json:"additional_info"`
	OrderItems     []OrderItem       `json:"order_items"`
}

type BulkOrder struct {
	FilePath    string    `json:"filePath"`
	User        User      `json:"user"`
	RequestTime time.Time `json:"requestTime"`
}

type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"fist_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}
