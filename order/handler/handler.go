package handler

import (
	"github.com/gin-gonic/gin"
	"order/entity"
	"order/service"
	"order/svc"
	"order/utility"
)

func OrderCreateHandler(svc *svc.ServiceContext) gin.HandlerFunc {
	return utility.WrapHandler(func(c *gin.Context) interface{} {
		var req entity.OrderCreateReq
		err := c.BindJSON(&req)
		if err != nil {
			panic(err)
		}
		if req.OrderID == 0 {
			panic("user_id not specified")
		}
		if req.UserID == 0 {
			panic("user_id not specified")
		}
		if req.OrderItems == nil {
			panic("order_items not specified")
		}
		if req.CouponID == 0 {
			panic("coupon_id not specified")
		}
		if req.CouponID != 0 && req.DiscountAmount == 0 {
			panic("discount_amount not specified")
		}

		s := service.NewOrderService(c, svc)
		err = s.OrderCreate(&req)
		if err != nil {
			entity.ResUnKnowError.WithError(err.Error())
		}
		return entity.ResOk
	})
}
