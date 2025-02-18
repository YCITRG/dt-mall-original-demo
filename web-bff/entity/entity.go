package entity

import "time"

type OrderItem struct {
	ProductID int64   `json:"product_id"`
	Quantity  int     `json:"quantity"`
	UnitPrice float64 `json:"unit_price"`
}

type Coupon struct {
	CouponID       int64   `json:"coupon_id"`
	DiscountAmount float64 `json:"discount_amount"`
}

type OrderCreateReq struct {
	Coupon     Coupon      `json:"coupon"`
	OrderItems []OrderItem `json:"order_items"`
}

type Order struct {
	ID          int64     `json:"id"`
	OrderID     int64     `json:"order_id"`
	UserID      int64     `json:"user_id"`
	OrderStatus string    `json:"order_status"`
	TotalAmount float64   `json:"total_amount"`
	CreatedTime time.Time `json:"created_time"`
	UpdatedTime time.Time `json:"updated_time"`
}
