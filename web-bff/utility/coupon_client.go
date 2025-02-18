package utility

import (
	"errors"
	"web-bff/entity"
)

type CouponClient struct {
	BizUnifiedClient
}

func (receiver *CouponClient) UseCoupon(userID, orderID, couponID int64, discountAmount float64) error {
	res := &entity.UseCouponResp{}
	err := receiver.PostJSONAndBindJSON("http://127.0.0.1:8082/api/coupon/use", map[string]any{
		"user_id":         userID,
		"order_id":        orderID,
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
