package utility

import (
	"errors"
	"web-bff/entity"
)

type OrderClient struct {
	BizUnifiedClient
}

func (receiver *OrderClient) CreateOrder(userID int64, orderItems []entity.OrderItem) (int64, error) {
	res := &entity.CreateOrderResp{}
	err := receiver.PostJSONAndBindJSON("http://127.0.0.1:8083/api/order/create", map[string]any{
		"user_id":     userID,
		"order_items": orderItems,
	}, res)
	if err != nil {
		return 0, err
	}

	if res.Code != entity.ResOk.Code {
		return 0, errors.New(res.Msg)
	}

	return res.Data.OrderID, nil
}

func (receiver *OrderClient) EnsureOrder(userID int64, orderID int64, discountAmount float64) (*entity.Order, error) {
	res := &entity.EnsureOrderResp{}
	err := receiver.PostJSONAndBindJSON("http://127.0.0.1:8083/api/order/ensure", map[string]any{
		"user_id":         userID,
		"order_id":        orderID,
		"discount_amount": discountAmount,
	}, res)
	if err != nil {
		return nil, err
	}

	if res.Code != entity.ResOk.Code {
		return nil, errors.New(res.Msg)
	}
	return &res.Data, nil
}
