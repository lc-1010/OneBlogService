package api

import (
	"github.com/gin-gonic/gin"
	"github.com/lc-1010/OneBlogService/global"
	"github.com/lc-1010/OneBlogService/internal/service"
	"github.com/lc-1010/OneBlogService/pkg/app"
	"github.com/lc-1010/OneBlogService/pkg/errcode"
)

// GetAuth returns the authentication model for the given app key and app secret.
func GetAuth(c *gin.Context) {
	param := service.AuthRequest{}

	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs :%v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.CheckAuth(&param)
	if err != nil {
		global.Logger.Errorf(c, "svc.CheckAuth errs :%v", errs)
		response.ToErrorResponse(errcode.UnauthorizedAuthNotExist)
		return
	}
	token, err := app.GenerateToken(param.Appkey, param.AppSecret)
	if err != nil {
		global.Logger.Errorf(c, "svc.GenerateToken errs :%v", errs)
		response.ToErrorResponse(errcode.UnauthoerizedTokenGenerate)
		return
	}
	response.ToResponse(gin.H{
		"token": token,
	})
}
