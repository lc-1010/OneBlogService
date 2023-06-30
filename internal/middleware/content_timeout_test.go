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
	ts := httptest.NewServer(r)
	// 增加中间件
	r.Use(ContextTimeout(3 * time.Second))
	// 创建一个测试路由
	r.GET("/timeout", func(c *gin.Context) {
		time.Sleep(10 * time.Second)
		c.String(http.StatusOK, "OK~")
	})
	r.GET("/test", func(c *gin.Context) {
		//访问不能访问的地址
		_, err := ctxhttp.Get(c.Request.Context(), http.DefaultClient, ts.URL+"/timeout")
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
		}
	})

	t.Run("Test timeout with valid duration", func(t *testing.T) {
		// Set up test case
		req, _ := http.NewRequest("GET", ts.URL+"/test", nil)
		w := httptest.NewRecorder()

		// Execute the test case
		r.ServeHTTP(w, req)

		// Verify the result
		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.Equal(t, context.DeadlineExceeded.Error(), w.Body.String())
	})
	ts.Close()
}

// 创建一个新的 HTTP 客户端
// 	client := &http.Client{
// 		Transport: &http.Transport{
// 			DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
// 				return net.DialTimeout(network, addr, 3*time.Second)
// 			},
// 		},
// 	}
// 	// 使用新的客户端发送请求
// 	_, err := ctxhttp.Get(c.Request.Context(), client, "/timeout")
// 	if err != nil {
// 		c.String(http.StatusInternalServerError, err.Error())
// 	}
// })

// select time.AfterFunc(1*time.Second, func() {

//})
// req, _ := http.NewRequest("GET", "/test", nil)
// w := httptest.NewRecorder()

// // Execute the test case with shorter timeout
// timer := time.AfterFunc(1*time.Second, func() {
// 	r.ServeHTTP(w, req)
// })

// // Verify the result
// select {
// case <-timer.C:
// 	assert.Equal(t, http.StatusInternalServerError, w.Code)
// 	assert.Equal(t, context.DeadlineExceeded.Error(), w.Body.String())
// }
