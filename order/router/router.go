package router

import (
	"github.com/gin-gonic/gin"
	"order/handler"
	"order/svc"
)

func RegisterRouters(app *gin.Engine, svcCtx *svc.ServiceContext) {
	app.POST("/api/order/create", handler.OrderCreateHandler(svcCtx))
	app.POST("/api/order/ensure", handler.OrderEnsureHandler(svcCtx))
}
