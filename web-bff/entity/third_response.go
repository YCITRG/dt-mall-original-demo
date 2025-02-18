package entity

type UseCouponResp struct {
	R
}

type UseInventoryResp struct {
	R
	Data map[string]any `json:"data"`
}

type CreateOrderResp struct {
	R
	Data OnlyOrderID `json:"data"`
}
type OnlyOrderID struct {
	OrderID int64 `json:"order_id"`
}

type EnsureOrderResp struct {
	R
	Data Order `json:"data"`
}

type CreatePaymentResp struct {
	R
	Data Payment `json:"data"`
}

type Payment struct {
	PaymentID int64 `json:"payment_id"`
}
