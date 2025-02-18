package main

import (
	"fmt"
	"github.com/dtm-labs/dtmcli/logger"
	"github.com/gin-gonic/gin"
	// 非直接引用包，使用 _ 避免语句被优化
	_ "github.com/go-sql-driver/mysql"
	"log"
	"web-bff/conf"
	"web-bff/router"
	"web-bff/svc"
)

func main() {

	c := conf.Config{
		AppApiPrefix: "/api/web",
		AppPort:      8080,
		DSN:          "root:mYsql123456_@tcp(192.168.9.218:3306)/yecao_mall_web_bff",
	}

	logger.InitLog("debug")
	gin.SetMode(gin.ReleaseMode)
	app := gin.Default()
	svcCtx := svc.NewServiceContext(c)
	router.RegisterRouters(app, svcCtx)
	log.Printf("web-bff service listening at %d", c.AppPort)
	err := app.Run(fmt.Sprintf(":%d", c.AppPort))
	if err != nil {
		panic(err)
	}
}
