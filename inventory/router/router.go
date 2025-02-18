package router

import (
	"github.com/gin-gonic/gin"
	"inventory/handler"
	"inventory/svc"
)

func RegisterRouters(app *gin.Engine, svcCtx *svc.ServiceContext) {
	app.POST("/api/inventory/use", handler.InventoryUseHandler(svcCtx))
}
