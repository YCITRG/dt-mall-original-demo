package main

import (
	"fmt"
	"github.com/dtm-labs/dtmcli/logger"
	"github.com/gin-gonic/gin"
	// 非直接引用包，使用 _ 避免语句被优化
	_ "github.com/go-sql-driver/mysql"
	"inventory/conf"
	"inventory/router"
	"inventory/svc"
	"log"
)

func main() {

	c := conf.Config{
		AppApiPrefix: "/api/inventory",
		AppPort:      8081,
		DSN:          "root:mYsql123456_@tcp(192.168.27.218:3306)/yecao_mall_inventory",
	}

	logger.InitLog("debug")
	gin.SetMode(gin.ReleaseMode)
	app := gin.Default()
	svcCtx := svc.NewServiceContext(c)
	router.RegisterRouters(app, svcCtx)
	log.Printf("inventory service listening at %d", c.AppPort)
	err := app.Run(fmt.Sprintf(":%d", c.AppPort))
	if err != nil {
		panic(err)
	}
}
