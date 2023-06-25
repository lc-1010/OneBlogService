// tag 标签管理
package v1

import "github.com/gin-gonic/gin"

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
// @Success 200 {object} model.BlogTag "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "请求错误"
// @Router /api/v1/tags [get]
func (t Tag) List(c *gin.Context) {}

// @Summary 创建标签
// @Produce json
// @Param name body string false "标签名" maxlength(100)
// @Param state body int false "状态" Enums(0,1) default(1)
// @Param created_by body string false "创建者"  minlength(3) maxlength(100)
// @Success 200 {object} model.BlogTag "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "请求错误"
// @Router /api/v1/tags [post]
func (t Tag) Create(c *gin.Context) {}

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
