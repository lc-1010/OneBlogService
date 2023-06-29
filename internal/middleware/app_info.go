package middleware

import "github.com/gin-gonic/gin"

// AppInfo set app info by middleware
func AppInfo() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Set("app_name", "blog-service")
		ctx.Set("version", "v1.0.0")
		ctx.Next()
	}
}
