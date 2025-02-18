package utility

import (
	"errors"
	"web-bff/entity"
)

type PaymentClient struct {
	BizUnifiedClient
}

func (receiver *PaymentClient) CreatePayment(userId, orderId int64, paymentAmount float64) (int64, error) {
	res := &entity.CreatePaymentResp{}

	err := receiver.PostJSONAndBindJSON("http://127.0.0.1:8084/api/payment/create", map[string]interface{}{
		"user_id":        userId,
		"order_id":       orderId,
		"payment_amount": paymentAmount,
	}, res)

	if err != nil {
		return 0, err
	}

	if res.Code != entity.ResOk.Code {
		return 0, errors.New(res.Msg + ": " + res.Error)
	}

	return res.Data.PaymentID, nil
}
