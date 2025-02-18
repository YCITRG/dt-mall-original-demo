package router

import (
	"coupon/handler"
	"coupon/svc"
	"github.com/gin-gonic/gin"
)

func RegisterRouters(app *gin.Engine, svcCtx *svc.ServiceContext) {
	app.POST("/api/coupon/use", handler.CouponUseHandler(svcCtx))
}
