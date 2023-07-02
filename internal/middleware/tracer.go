package middleware

import (
	"github.com/gin-gonic/gin"
)

// Tracing returns a middleware handler function for the Gin framework.
//
// The returned function is the actual middleware handler.
func Tracing() func(c *gin.Context) {
	// 这个函数返回一个用于 Gin 框架的中间件处理函数。

	return func(ctx *gin.Context) {
		// 这是实际的中间件处理函数。

		// var newCtx context.Context

		// var traceID string
		// var spanID string
		// //tr := global.Tracer.Tracer("middleware")

		// // 将跟踪 ID 和子跟踪 ID 设置为 Gin 上下文的自定义头部。
		// ctx.Set("X-Trace-ID", traceID)
		// ctx.Set("X-Span-ID", spanID)

		// // 使用包含跟踪的新上下文更新 Gin 请求的上下文。
		// ctx.Request = ctx.Request.WithContext(newCtx)

		// 调用链中的下一个处理程序。
		ctx.Next()
	}
}
