// tag 标签管理
package v1

import (
	"github.com/lc-1010/OneBlogService/pkg/app"

	"github.com/gin-gonic/gin"
	"github.com/lc-1010/OneBlogService/global"
	"github.com/lc-1010/OneBlogService/internal/service"
	"github.com/lc-1010/OneBlogService/pkg/errcode"
)

type Tag struct{}

func NewTag() Tag {
	return Tag{}
}

// @Summary 获取多个标签
// @Produce json
// @Param name query string false "标签名" maxlength(100)
// @Param state query int false "状态" Enums(0,1) default(1)
// @Param page query int false "页码"
// @Param paage_size query int false "每页数量"
// @Success 200 {object} model.TagSwagger "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "请求错误"
// @Router /api/v1/tags [get]
func (t Tag) List(c *gin.Context) {
	// 校验参数
	// param := struct {
	// 	Name  string `form:"name,omitempty" binding:"max=100"`
	// 	State uint8  `form:"state,default=1" binding:"oneof=0 1" `
	// }{}
	param := service.GetTagListRequest{}

	response := app.NewResponse(c)
	// 绑定参数进行校验
	valid, errs := app.BindAndValid(c, &param)

	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		e := errcode.InvalidParams.WithDetails(errs.Errors()...)
		response.ToErrorResponse(e)
		return
	}

	svc := service.New(c.Request.Context())

	pager := app.Pager{Page: app.GetPage(c), PageSize: app.GetPageSize(c)}
	totalRows, err := svc.CountTag(&service.ConutTagRequest{Name: param.Name, State: param.State})
	if err != nil {
		global.Logger.Errorf(c, "svc.CountTag err: %v", err)
		response.ToErrorResponse(errcode.ErrorCountTagFail)
		return
	}
	tags, err := svc.GetTagList(&param, &pager)
	if err != nil {
		global.Logger.Errorf(c, "svc.GetTagList err: %v", err)
		response.ToErrorResponse(errcode.ErrorGetTagListFail)
		return
	}

	response.ToResponseList(tags, totalRows)

}

// @Summary 创建标签
// @Produce json
// @Param request body model.UpdateTagRequest true "标签信息"
// @Success 200 {object} model.BlogTag "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "请求错误"
// @Router /api/v1/tags [post]
func (t Tag) Create(c *gin.Context) {
	param := service.CrateTagRquest{}
	respone := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		respone.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.CrateTag(&param)
	if err != nil {
		global.Logger.Errorf(c, "svc.CrateTag err: %v", err)
		respone.ToErrorResponse(errcode.ErrorCreateTagFail)
	}
	respone.ToResponse(gin.H{})
}

// @Summary 更新标签
// @Produce json
// @Param request body model.UpdateTagRequest true "标签信息"
// @Success 200 {object} model.BlogTag "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "请求错误"
// @Router /api/v1/tags [put]
func (t Tag) Update(c *gin.Context) {
	param := service.UpdateTagRequest{}
	response := app.NewResponse(c)

	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid err :%v", errs)
		response.ToErrorResponse(errcode.ErrorUpdateTagFail.WithDetails(errs.Errors()...))
	}
	svc := service.New(c.Request.Context())
	err := svc.UpdateTag(&param)
	if err != nil {
		global.Logger.Errorf(c, "svc.UpdateTag err: %v", err)
		response.ToErrorResponse(errcode.ErrorUpdateTagFail)
		return
	}
	response.ToResponse(gin.H{})

}

// @Summary 删除标签
// @Produce json
// @Param id path int true "标签id"
// @Success 200 {string} string "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "请求错误"
// @Router /api/v1/tags/{id} [delete]
func (t Tag) Delete(c *gin.Context) {
	param := service.DeleteTagRequest{}
	response := app.NewResponse(c)

	valid, errs := app.BindAndValid(c, &param)

	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid err:%v", errs)
		response.ToErrorResponse(errcode.ErrorDeleteTagFail.WithDetails(errs.Errors()...))
		return
	}
	svc := service.New(c.Request.Context())
	err := svc.DeleteTag(&param)
	if err != nil {
		global.Logger.Errorf(c, "svc.DeleteTag err:%v", err)
		response.ToErrorResponse(errcode.ErrorDeleteTagFail)
		return
	}
	response.ToResponse(gin.H{})
}
