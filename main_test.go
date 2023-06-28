package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/lc-1010/OneBlogService/pkg/app"
)

// func TestPing(t *testing.T) {
// 	r := gin.Default()
// 	r.GET("/ping", func(ctx *gin.Context) {
// 		global.Logger.SetTraceInfo(ctx).Infof(ctx, "%s for test ping,path:%s", ctx.HandlerName(), ctx.Request.URL.Path)
// 		ctx.JSON(http.StatusOK, gin.H{"message": "pong"})
// 	})
// 	//r.Run()
// }

func TestResponse(t *testing.T) {
	r := gin.New()
	r.GET("/test/ping", func(c *gin.Context) {
		app.NewResponse(c).ToResponse(map[string]string{"msg": "ping pong is ok"})
	})

	req := httptest.NewRequest("GET", "/test/ping", nil)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req
	r.HandleContext(c)
	assert.Equal(t, w.Code, http.StatusOK)

	expcted := `{"msg":"ping pong is ok"}`
	assert.Equal(t, w.Body.String(), expcted)

}
func TestMultiHandler(t *testing.T) {
	r := gin.Default()

	// 在路由中间件中设置消息
	r.Use(func(c *gin.Context) {
		c.Set("message", "Hello, world!")
		c.Next()
	})

	// 注册一个路由，将多个处理函数关联到这个路由上
	r.GET("/hello", func(c *gin.Context) {
		// message := c.MustGet("message").(string)
		// c.String(http.StatusOK, message)
		// c.Set("message", message)
		c.Next()
	}, func(c *gin.Context) {
		message := c.MustGet("message").(string)
		c.String(http.StatusOK, message+" from world")
	})

	// 创建一个新的 HTTP 请求
	req, err := http.NewRequest("GET", "/hello", nil)
	if err != nil {
		t.Fatal(err)
	}

	// 创建一个新的 HTTP 响应
	w := httptest.NewRecorder()

	// 发送 HTTP 请求到路由处理函数
	r.ServeHTTP(w, req)

	// 检查 HTTP 响应的状态码和内容
	if w.Code != http.StatusOK {
		t.Errorf("Unexpected status code %d", w.Code)
	}
	exp := "Hello, world! from world"
	//msg := fmt.Sprintf("\nUnexpected body: %s", exp)
	assert.Equal(t, w.Body.String(), exp)

}

func TestMain(t *testing.T) {
	main()
}
