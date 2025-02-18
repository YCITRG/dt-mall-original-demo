package entity

type CouponUseReq struct {
	UserID         int64   `json:"user_id"`
	OrderID        int64   `json:"order_id"`
	CouponID       int64   `json:"coupon_id"`
	DiscountAmount float64 `json:"discount_amount"`
}
