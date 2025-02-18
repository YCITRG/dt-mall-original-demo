package handler

import (
	"github.com/gin-gonic/gin"
	"inventory/entity"
	"inventory/service"
	"inventory/svc"
	"inventory/utility"
)

func InventoryUseHandler(svc *svc.ServiceContext) gin.HandlerFunc {
	return utility.WrapHandler(func(c *gin.Context) interface{} {

		var req entity.InventoryUseReq
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
		if req.ProductID == 0 {
			panic("product_id not specified")
		}
		s := service.NewInventoryService(c, svc)
		err = s.InventoryUse(&req)
		if err != nil {
			return entity.ResUnKnowError.WithError(err.Error())
		}
		return entity.ResOk
	})
}
