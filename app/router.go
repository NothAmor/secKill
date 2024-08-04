package app

import (
	"secKill/app/handler"

	"github.com/gin-gonic/gin"
)

func RegisterRouter(engine *gin.Engine) {
	v1 := engine.Group("/v1")

	// 用户
	user := v1.Group("/user")
	user.POST("/register", handler.Register)
	user.POST("/login")

	// 秒杀商品
	product := v1.Group("/product")
	product.POST("/create")
	product.POST("/update")
	product.POST("/delete")
	product.POST("/list")
	product.POST("/get")
}
