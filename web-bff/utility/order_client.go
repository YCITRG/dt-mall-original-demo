package utility

import (
	"errors"
	"web-bff/entity"
)

type OrderClient struct {
	BizUnifiedClient
}

func (receiver *OrderClient) CreateOrder(orderID, userID int64, orderItems []entity.OrderItem, couponID int64, discountAmount float64) error {
	res := &entity.CreateOrderResp{}
	err := receiver.PostJSONAndBindJSON("http://127.0.0.1:8083/api/order/create", map[string]any{
		"order_id":        orderID,
		"user_id":         userID,
		"order_items":     orderItems,
		"coupon_id":       couponID,
		"discount_amount": discountAmount,
	}, res)
	if err != nil {
		return err
	}

	if res.Code != entity.ResOk.Code {
		return errors.New(res.Msg + ": " + res.Error)
	}
	return nil
}
