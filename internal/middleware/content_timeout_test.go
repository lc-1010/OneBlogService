package middleware

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"golang.org/x/net/context/ctxhttp"
)

// TestContextTimeout is a Go function that tests the behavior of a route handler with a context timeout middleware.
//
// It sets up a Gin router and adds a context timeout middleware. Then it creates a test route that makes a request to an inaccessible URL.
// If the request fails with an error, it responds with an HTTP 500 status code and the error message.
//
// The function also includes a subtest "Test timeout with valid duration" that sets up a test case, executes it, and verifies the result.
//
// It does not take any parameters and does not return any values.
func TestContextTimeout(t *testing.T) {
	r := gin.Default()

	// 增加中间件
	r.Use(ContextTimeout(5 * time.Second))
	// 创建一个测试路由
	r.GET("/test", func(c *gin.Context) {
		//访问不能访问的地址
		_, err := ctxhttp.Get(c.Request.Context(), http.DefaultClient, "https://www.google.com/")
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
		}
	})

	t.Run("Test timeout with valid duration", func(t *testing.T) {
		// Set up test case
		req, _ := http.NewRequest("GET", "/test", nil)
		w := httptest.NewRecorder()

		// Execute the test case
		r.ServeHTTP(w, req)

		// Verify the result
		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.Equal(t, context.DeadlineExceeded.Error(), w.Body.String())
	})

}
