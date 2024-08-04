package app

import (
	user "secKill/app/handler/user"

	"github.com/gin-gonic/gin"
)

func RegisterRouter(engine *gin.Engine) {
	v1 := engine.Group("/v1")

	// 用户
	userGroup := v1.Group("/user")
	userGroup.POST("/register", user.Register)
	userGroup.POST("/login", user.Login)

	// 秒杀商品
	product := v1.Group("/product")
	product.POST("/create")
	product.POST("/update")
	product.POST("/delete")
	product.POST("/list")
	product.POST("/get")
}
