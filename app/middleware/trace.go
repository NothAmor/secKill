package middleware

import (
	"secKill/utils/logger"

	"github.com/gin-gonic/gin"
)

// TraceMiddleware 链路追踪中间件
func TraceMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		traceNode := logger.GenLogTraceMetadata()
		c.Set(logger.GetMetadataKey(), traceNode)
		c.Next()
	}
}
