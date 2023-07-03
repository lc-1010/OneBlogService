package routers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	"github.com/go-playground/locales/zh_Hant_TW"
	ut "github.com/go-playground/universal-translator"
	validator "github.com/go-playground/validator/v10"

	_ "github.com/lc-1010/OneBlogService/docs" // 必须引入不然找不到文件
	"github.com/lc-1010/OneBlogService/global"
	"github.com/lc-1010/OneBlogService/internal/middleware"
	"github.com/lc-1010/OneBlogService/internal/routers/api"
	v1 "github.com/lc-1010/OneBlogService/internal/routers/api/v1"
	"github.com/lc-1010/OneBlogService/internal/routers/ping"
	"github.com/lc-1010/OneBlogService/pkg/limiter"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var v *validator.Validate
var uni *ut.UniversalTranslator

// methodLimiters
var methodLimiters = limiter.NewMethodLimiter().AddBucket(
	limiter.LimiterBucketRule{
		Key:          "/auth",
		FillInterval: time.Second,
		Capacity:     10,
		Quantum:      10,
	},
)

func init() {
	// 初始化语言设置
	uni = ut.New(en.New(), zh.New(), zh_Hant_TW.New())
	//trans, _ = uni.GetTranslator("zh")
	v, _ = binding.Validator.Engine().(*validator.Validate)
}

// NewRouter tags articles curd
func NewRouter() *gin.Engine {
	r := gin.New()
	if global.ServerSetting.RunMode == "debug" {
		// middleware 中间件
		r.Use(gin.Logger())
		r.Use(gin.Recovery())

	} else {
		r.Use(middleware.Recovery())
		r.Use(middleware.AccessLog())
	}
	// Translate
	r.Use(middleware.TranslateSet(uni, v))
	// limiter
	r.Use(middleware.Limiter(methodLimiters))
	// context with timeout
	r.Use(middleware.ContextTimeout(global.AppSetting.DefaultContextTimeout * time.Second))
	//tracing
	r.Use(middleware.Tracing())
	// 业务路由
	article := v1.NewArticle()
	tags := v1.NewTag()

	// test
	ping := ping.NewPing()
	//upload image
	upload := api.NewUpload()
	r.POST("/upload/file", upload.UploadFile)
	r.StaticFS("/static", http.Dir(global.AppSetting.UploadSavePath))

	// auth
	r.POST("/auth", api.GetAuth)
	//ping test
	p := r.Group("/test")
	{
		p.GET("/ping", ping.Pong)
	}

	// api router
	apiv1 := r.Group("/api/v1")
	//apiv1.Use(middleware.JWT())
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

	// swag router
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
