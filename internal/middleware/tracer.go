package middleware

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/lc-1010/OneBlogService/global"
	opentracing "github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/uber/jaeger-client-go"
)

// Tracing returns a middleware handler function for the Gin framework.
//
// The returned function is the actual middleware handler.
func Tracing() func(c *gin.Context) {
	// 这个函数返回一个用于 Gin 框架的中间件处理函数。

	return func(ctx *gin.Context) {
		// 这是实际的中间件处理函数。

		var newCtx context.Context
		var span opentracing.Span

		// 从传入的 HTTP 头中提取跟踪上下文。
		spanCtx, err := opentracing.GlobalTracer().Extract(
			opentracing.HTTPHeaders,
			opentracing.HTTPHeadersCarrier(ctx.Request.Header),
		)
		if err != nil {
			// 如果在头中找不到跟踪上下文，则启动一个新的跟踪。
			span, newCtx = opentracing.StartSpanFromContextWithTracer(
				ctx.Request.Context(),
				global.Tracer,
				ctx.Request.URL.Path,
			)
		} else {
			// 如果在头中找到跟踪上下文，则启动一个新的子跟踪。
			span, newCtx = opentracing.StartSpanFromContextWithTracer(
				ctx.Request.Context(),
				global.Tracer,
				ctx.Request.URL.Path,
				opentracing.ChildOf(spanCtx),
				opentracing.Tag{Key: string(ext.Component), Value: "HTTP"},
			)
		}
		defer span.Finish()

		var traceID string
		var spanID string

		switch spanContext := span.Context().(type) {
		case jaeger.SpanContext:
			// 如果跟踪上下文的类型是 "jaeger.SpanContext"（Jaeger 跟踪器），
			// 则从中获取跟踪 ID 和子跟踪 ID。
			//jaegerContext := spanContext.(jaeger.SpanContext)
			traceID = spanContext.TraceID().String()
			spanID = spanContext.SpanID().String()
		}

		// 将跟踪 ID 和子跟踪 ID 设置为 Gin 上下文的自定义头部。
		ctx.Set("X-Trace-ID", traceID)
		ctx.Set("X-Span-ID", spanID)

		// 使用包含跟踪的新上下文更新 Gin 请求的上下文。
		ctx.Request = ctx.Request.WithContext(newCtx)

		// 调用链中的下一个处理程序。
		ctx.Next()
	}
}
