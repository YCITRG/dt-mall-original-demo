package router

import (
	"github.com/gin-gonic/gin"
	"web-bff/handler"
	"web-bff/svc"
)

func RegisterRouters(app *gin.Engine, svcCtx *svc.ServiceContext) {
	app.POST("/api/web-bff/order", handler.OrderCreateHandler(svcCtx))
}
