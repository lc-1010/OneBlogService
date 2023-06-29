package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lc-1010/OneBlogService/global"
	"github.com/lc-1010/OneBlogService/pkg/app"
	"github.com/lc-1010/OneBlogService/pkg/email"
	"github.com/lc-1010/OneBlogService/pkg/errcode"
)

// Recovery recovery middleware
// Recovery is a middleware function that recovers from panics in the Gin framework.
func Recovery() gin.HandlerFunc {
	// Create a new email object with the SMTP settings from the global configuration.
	defailMailer := email.NewEmail(&email.SMTPInfo{
		Host:     global.EmailSettings.Host,
		Port:     global.EmailSettings.Port,
		IsSSL:    global.EmailSettings.IsSSl,
		UserName: global.EmailSettings.UserName,
		Password: global.EmailSettings.Password,
		From:     global.EmailSettings.From,
	})

	// Return a closure function that will be executed as a middleware in the Gin framework.
	return func(ctx *gin.Context) {
		// Defer a function to recover from any panics that occur within the middleware.
		defer func() {
			// Recover from a panic and store the error value in the 'err' variable.
			err := recover()
			if err != nil {
				// Log the recovered error with the caller's frames using the global logger.
				global.Logger.WithCallersFrames().Errorf(ctx, "panic recover err: %v", err)
			}

			// Send an email notification with the error details using the 'defailMailer' object.
			err = defailMailer.SendMail(
				global.EmailSettings.To,
				fmt.Sprintf("异常抛出，发生时间: %d", time.Now().Unix()),
				fmt.Sprintf("错误信息: %v", err),
			)
			if err != nil {
				// If sending the email fails, log the error and panic.
				global.Logger.Panicf(ctx, "defailMailer.SendMail err: %v", err)
			}

			// Return an error response to the client using the 'app' package.
			app.NewResponse(ctx).ToErrorResponse(errcode.ServerError)
		}()

		// Call the next middleware or handler in the Gin framework.
		ctx.Next()
	}
}
