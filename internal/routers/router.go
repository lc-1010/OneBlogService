package routers

import (
	"github.com/gin-gonic/gin"
	_ "github.com/lc-1010/OneBlogService/docs"
	v1 "github.com/lc-1010/OneBlogService/internal/routers/api/v1"
	"github.com/lc-1010/OneBlogService/internal/routers/ping"
)

// NewRouter tags articles curd
func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	article := v1.NewArticle()
	tags := v1.NewTag()

	ping := ping.NewPing()

	p := r.Group("/test")
	{
		p.GET("/ping", ping.Pong)
	}

	apiv1 := r.Group("/api/v1")
	{
		//tags
		apiv1.POST("/tags", tags.Create)
		apiv1.DELETE("/tags/:id", tags.Delete)
		apiv1.PUT("/tags/:id", tags.Update)
		apiv1.PATCH("/tags/:id/state", tags.Update)
		apiv1.GET("/tags", tags.List)
		//articles
		apiv1.POST("/articles", article.Create)
		apiv1.GET("/articles", article.List)
		apiv1.GET("/articles/:id", article.Get)
		apiv1.PATCH("/articles/:id/state", article.Update)
		apiv1.PUT("/articles/:id", article.Update)
		apiv1.DELETE("/articles/:id", article.Delete)
	}
	return r
}
