package handler

import (
	"net/http"
	"secKill/app/logic"
	"secKill/app/proto"
	"secKill/utils/ctx"
	"secKill/utils/logger"
	"secKill/utils/secgin"

	"github.com/gin-gonic/gin"
)

// Register 用户注册
func Register(c *gin.Context) {
	tag := "handler.User.Register"
	ctx := ctx.TransferToContext(c)
	var args proto.CreateUserReq
	if err := c.ShouldBindJSON(&args); err != nil {
		logger.Ex(ctx, tag, "ShouldBindJson failed. err:[%+v]", err)
		c.JSON(http.StatusOK, secgin.Error(secgin.PARAM_ERROR, nil))
		return
	}

	reply := proto.CreateUserResp{}
	err := logic.Register(ctx, args, &reply)
	if err != nil {
		logger.Ex(ctx, tag, "logic.Register failed. err:[%+v]", err)
		c.JSON(http.StatusOK, secgin.Error(secgin.SYS_DEFAULT, nil))
		return
	}

	c.JSON(http.StatusOK, secgin.Success(reply))
}
