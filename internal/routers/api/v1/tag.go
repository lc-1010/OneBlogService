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
// @Success 200 {object} model.TagSwagger  "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "请求错误"
// @Router /api/v1/tags [get]
func (t Tag) List(c *gin.Context) {
	response := app.NewResponse(c)
	// 绑定参数进行校验
	valid, errs := app.BindAndValid(c, &service.TagListRequest{})
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		e := errcode.InvalidParams.WithDetails(errs.Errors()...)
		response.ToErrorResponse(e)
		return
	}
	response.ToResponse(gin.H{})
}

// @Summary 创建标签
// @Produce json
// @Param name body string false "标签名" maxlength(100)
// @Param state body int false "状态" Enums(0,1) default(1)
// @Param created_by body string false "创建者"  minlength(3) maxlength(100)
// @Success 200 {object} model.BlogTag "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "请求错误"
// @Router /api/v1/tags [post]
func (t Tag) Create(c *gin.Context) {

}

// @Summary 更新标签
// @Produce json
// @Param id path int true "标签id"
// @Param name body string false "标签名"  minlength(3) maxlength(100)
// @Param state body int false "状态" Enums(0,1) default(1)
// @Param created_by body string false "创建者"  minlength(3) maxlength(100)
// @Success 200 {object} model.BlogTag "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "请求错误"
// @Router /api/v1/tags [put]
func (t Tag) Update(c *gin.Context) {}

// @Summary 删除标签
// @Produce json
// @Param id path int true "标签id"
// @Success 200 {string} string "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "请求错误"
// @Router /api/v1/tags [delete]
func (t Tag) Delete(c *gin.Context) {}

/********* mutilpe body *******************/
/*

// @Summary 获取多个标签
// @Produce json
// @Param name query string false "标签名" maxlength(100)
// @Param state query int false "状态" Enums(0,1) default(1)
// @Param page query int false "页码"
// @Param paage_size query int false "每页数量"
// @Success 200 {object} model.BlogTag "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "请求错误"
// @Router /api/v1/tags [get]
func (t Tag) List(c *gin.Context) {}

// @Summary 创建标签
// @Produce json
// @Param request body model.UpdateTagRequest true "标签信息"
// @Success 200 {object} model.BlogTag "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "请求错误"
// @Router /api/v1/tags [post]
func (t Tag) Create(c *gin.Context) {}

// @Summary 更新标签
// @Produce json
// @Param request body model.UpdateTagRequest true "标签信息"
// @Success 200 {object} model.BlogTag "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "请求错误"
// @Router /api/v1/tags [put]
func (t Tag) Update(c *gin.Context) {}

// @Summary 删除标签
// @Produce json
// @Param id path int true "标签id"
// @Success 200 {string} string "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "请求错误"
// @Router /api/v1/tags/{id} [delete]
func (t Tag) Delete(c *gin.Context) {}

*/
