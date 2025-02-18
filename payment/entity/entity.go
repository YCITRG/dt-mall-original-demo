package entity

type PaymentCreateReq struct {
	UserID        int64   `json:"user_id"`
	OrderID       int64   `json:"order_id"`
	PaymentAmount float64 `json:"payment_amount"`
}
