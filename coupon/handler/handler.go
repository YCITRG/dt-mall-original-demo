package handler

import (
	"coupon/entity"
	"coupon/service"
	"coupon/svc"
	"coupon/utility"
	"github.com/gin-gonic/gin"
)

func CouponUseHandler(svc *svc.ServiceContext) gin.HandlerFunc {
	return utility.WrapHandler(func(c *gin.Context) interface{} {
		var req entity.CouponUseReq
		err := c.BindJSON(&req)
		if err != nil {
			panic(err)
		}
		if req.UserID == 0 {
			panic("user_id not specified")
		}
		if req.OrderID == 0 {
			panic("order_Id not specified")
		}
		if req.CouponID == 0 {
			panic("coupon_id not specified")
		}
		if req.DiscountAmount == 0 {
			panic("discount_amount not specified")
		}
		s := service.NewCouponService(c, svc)
		err = s.CouponUse(&req)
		if err != nil {
			return entity.ResUnKnowError.WithError(err.Error())
		}
		return entity.ResOk
	})
}
