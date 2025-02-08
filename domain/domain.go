package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Order struct {
	ID             string             `json:"id"`
	UserName       string             `json:"user_name"`
	TenantID       string             `bson:"tenant_id,omitempty" json:"tenantid"`
	HubID          string             `bson:"hub_id,omitempty" json:"hubid"`
	SellerID       string             `json:"seller_id"`
	Status         string             `json:"status"`
	AdditionalInfo map[string]string  `json:"additional_info"`
	CreatedAt      primitive.DateTime `json:"createdat"`
	UpdatedAt      primitive.DateTime `json:"updatedat"`
	OrderItems     []OrderItem        `json:"order_items"`
}

type OrderItem struct {
	SKU_ID   string `json:"sku_id"`
	Quantity int    `json:"quantity"`
}

type CreateOrderRequest struct {
	Path           string            `json:"path"`
	UserID         string            `json:"user_id"`
	TenantID       string            `json:"tenant_id"`
	HubID          string            `json:"hub_id"`
	SellerID       string            `json:"seller_id"`
	AdditionalInfo map[string]string `json:"additional_info"`
	OrderItems     []OrderItem       `json:"order_items"`
}
