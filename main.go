package main

import (
	"fmt"
	"log"
	"secKill/app"
	"secKill/app/common"
	"secKill/app/initialization"
	"secKill/app/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	initialization.Init()

	engine := gin.Default()
	engine.Use(
		middleware.TraceMiddleware(),
		middleware.Recovery(),
	)
	app.RegisterRouter(engine)

	if err := engine.Run(fmt.Sprintf("%s:%d", common.Config.Sys.Host, common.Config.Sys.Port)); err != nil {
		log.Printf("gin Engine.Run failed. err:[%+v]\n", err)
		panic(err)
	}
}
