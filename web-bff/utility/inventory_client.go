package utility

import (
	"errors"
	"web-bff/entity"
)

type InventoryClient struct {
	BizUnifiedClient
}

func (receiver *InventoryClient) UseInventory(userID, orderID, productID int64, productCount int) error {

	res := &entity.UseInventoryResp{}

	err := receiver.PostJSONAndBindJSON("http://127.0.0.1:8081/api/inventory/use", map[string]any{
		"user_id":       userID,
		"order_id":      orderID,
		"product_id":    productID,
		"product_count": productCount,
	}, res)
	if err != nil {
		return err
	}

	if res.Code != entity.ResOk.Code {
		return errors.New(res.Msg)
	}
	return nil
}
