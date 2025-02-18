package router

import (
	"github.com/gin-gonic/gin"
	"payment/handler"
	"payment/svc"
)

func RegisterRouters(app *gin.Engine, svcCtx *svc.ServiceContext) {
	app.POST("/api/payment/create", handler.PaymentCreateHandler(svcCtx))
}
