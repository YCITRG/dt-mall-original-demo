package handler

import (
	"github.com/gin-gonic/gin"
	"web-bff/entity"
	"web-bff/service"
	"web-bff/svc"
	"web-bff/utility"
)

func OrderCreateHandler(svc *svc.ServiceContext) gin.HandlerFunc {
	return utility.WrapHandler(func(c *gin.Context) interface{} {
		var req entity.OrderCreateReq
		err := c.BindJSON(&req)
		if err != nil {
			panic(err)
		}
		if req.Coupon.CouponID == 0 {
			panic("coupon_id not specified")
		}
		if req.Coupon.DiscountAmount == 0 {
			panic("discount_amount not specified")
		}
		if req.OrderItems == nil {
			panic("order_items not specified")
		}
		s := service.NewOrderService(c, svc)
		res, err := s.OrderCreate(&req)
		if err != nil {
			return entity.ResUnKnowError.WithError(err.Error())
		}
		return entity.ResOk.WithData(res)
	})
}
