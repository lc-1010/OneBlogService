package middleware

import (
	"bytes"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lc-1010/OneBlogService/global"
	"github.com/lc-1010/OneBlogService/pkg/logger"
)

// AccessLogWriter access log writer
type AccessLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

// Write writes the given bytes to the AccessLogWriter.
//
// It takes a byte slice as the parameter and returns the number of bytes written and an error if any.

func (w AccessLogWriter) Write(b []byte) (int, error) {
	if n, err := w.body.Write(b); err != nil {
		return n, err
	}
	return w.ResponseWriter.Write(b)
}

// AccessLog returns a Gin middleware function that logs access information.
//
// This function takes a Gin context as a parameter and logs the request and response information.
// It returns no values.
func AccessLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		bodyWriter := &AccessLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = bodyWriter

		begin := time.Now().Unix()
		c.Next()
		end := time.Now().Unix()
		fields := logger.Fields{
			"request":  c.Request.PostForm.Encode(),
			"response": bodyWriter.body.String(),
		}
		global.Logger.WithFields(fields).Infof(c, "access log:method %s, uri %s, status %d, beginTime %d, endTime %d,cost %d",
			c.Request.Method,
			c.Request.RequestURI,
			c.Writer.Status(),
			begin,
			end,
			end-begin)
	}
}
