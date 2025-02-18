package handler

import (
	"github.com/gin-gonic/gin"
	"payment/entity"
	"payment/service"
	"payment/svc"
	"payment/utility"
)

func PaymentCreateHandler(svc *svc.ServiceContext) gin.HandlerFunc {
	return utility.WrapHandler(func(c *gin.Context) interface{} {
		var req entity.PaymentCreateReq
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
		if req.PaymentAmount == 0 {
			panic("payment_amount not specified")
		}
		s := service.NewPaymentService(c, svc)
		data, err := s.PaymentCreate(&req)
		if err != nil {
			return entity.ResUnKnowError.WithError(err.Error())
		}
		return entity.ResOk.WithData(data)
	})
}
