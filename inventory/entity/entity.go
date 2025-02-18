package entity

type InventoryUseReq struct {
	UserID       int64 `json:"user_id"`
	OrderID      int64 `json:"order_id"`
	ProductID    int   `json:"product_id"`
	ProductCount int   `json:"product_count"`
}
