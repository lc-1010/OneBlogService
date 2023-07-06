package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/lc-1010/OneBlogService/global"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

// Tracing returns a middleware handler function for the Gin framework.
//
// The returned function is the actual middleware handler.
// todo: 增加全链路的追踪
func Tracing() func(c *gin.Context) {
	// 这个函数返回一个用于 Gin 框架的中间件处理函数。

	return func(ctx *gin.Context) {
		// 这是实际的中间件处理函数。

		// var traceID string 0ad1348f1403169275002100356696 产生规则是： 服务器 IP + ID 产生的时间 + 自增序列 + 当前进程号
		//前 8 位 0ad1348f 即产生 TraceId 的机器的 IP，这是一个十六进制的数字，每两位代表 IP 中的一段，我们把这个数字，按每两位转成 10 进制即可得到常见的 IP 地址表示方式 10.209.52.143，您也可以根据这个规律来查找到请求经过的第一个服务器。
		//后面的 13 位 1403169275002 是产生 TraceId 的时间。
		//之后的 4 位 1003 是一个自增的序列，从 1000 涨到 9000，到达 9000 后回到 1000 再开始往上涨。
		//最后的 5 位 56696 是当前的进程 ID，为了防止单机多进程出现 TraceId 冲突的情况，所以在 TraceId 末尾添加了当前的进程 ID
		// var spanID string
		// SpanId 代表本次调用在整个调用链路树中的位置。假设一个 Web 系统 A 接收了一次用户请求，那么在这个系统的 SOFATracer MVC 日志中，记录下的 SpanId 是 0，
		//代表是整个调用的根节点，如果 A 系统处理这次请求，需要通过 RPC 依次调用 B、C、D 三个系统，
		//那么在 A 系统的 SOFATracer RPC 客户端日志中，SpanId 分别是 0.1，0.2 和 0.3，在 B、C、D 三个系统的 SOFATracer RPC
		//服务端日志中，SpanId 也分别是 0.1，0.2 和 0.3；如果 C 系统在处理请求的时候又调用了 E，F 两个系统，那么 C 系统中对应的 SOFATracer RPC 客户端日志是
		//0.2.1 和 0.2.2，E、F 两个系统对应的 SOFATracer RPC 服务端日志也是 0.2.1 和 0.2.2。

		tr := global.Tracer.Tracer("api")
		// 将跟踪 ID 和子跟踪 ID 设置为 Gin 上下文的自定义头部。

		nctx, spn := tr.Start(ctx, ctx.Request.URL.Path, trace.WithAttributes(
			attribute.String("method", ctx.Request.Method),
			attribute.String("ip", ctx.ClientIP()),
			attribute.String("uri", ctx.Request.RequestURI),
		), trace.WithNewRoot()) //totdo

		defer spn.End()

		// 使用包含跟踪的新上下文更新 Gin 请求的上下文。
		ctx.Request = ctx.Request.WithContext(nctx)

		// 调用链中的下一个处理程序。
		ctx.Next()
	}
}
