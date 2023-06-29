package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/lc-1010/OneBlogService/global"
	"github.com/lc-1010/OneBlogService/pkg/app"
	"github.com/lc-1010/OneBlogService/pkg/errcode"
	"github.com/lc-1010/OneBlogService/pkg/limiter"
)

// Limiter for router
func Limiter(l limiter.LimiterInterface) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		key := l.Key(ctx)
		if bucket, ok := l.GetBucket(key); ok {
			count := bucket.TakeAvailable(1)
			if count == 0 {
				response := app.NewResponse(ctx)
				response.ToErrorResponse(errcode.TooManyRequests)
				global.Logger.Debugf(ctx, "too many requests")
				ctx.Abort()
				return
			}
		}
		ctx.Next()
	}

}
