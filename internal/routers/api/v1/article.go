// 文章管理
package v1

import (
	"blog-server/global"
	"fmt"
	"net/http"

	"github.com/lc-1010/OneBlogService/internal/model"

	"github.com/gin-gonic/gin"
)

type Article struct {
	model.BlogArticle
}

func NewArticle() Article {
	return Article{}
}

// @Summary 获取单独文章
// @Produce json
// @Param id path int true "文章id"
// @Success 200 {object} model.BlogArticle "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "请求错误"
// @Router /api/v1/articles/{id} [get]
func (t Article) Get(c *gin.Context) {}

// @Summary 获取多篇文章
// @Produce json
// @Param name query string false "文章名"
// @Param tag_id query int false "标签ID"
// @Param page query int false "页码"
// @Param paage_size query int false "每页数量"
// @Success 200 {object} model.BlogArticle "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "请求错误"
// @Router /api/v1/articles [get]
func (t Article) List(c *gin.Context) {}

// @Summary 创建文章
// @Produce json
// @Param title body string false "文章名" maxlength(100)
// @Param tag_id body int false "标签"
// @Param desc body string fasle "简述"
// @Param cover_image_url  body string false "图片地址"
// @Param content body string fasle "内容"
// @Param state body int fasle "状态"
// @Success 200 {object} model.BlogArticle "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "请求错误"
// @Router /api/v1/articles [post]
func (t Article) Create(c *gin.Context) {
	db := global.DBEngine
	a := NewArticle()
	db.Limit(1).Find(&a)
	fmt.Println(a)
	c.JSON(http.StatusOK,
		map[string]string{"res": " ok"})
}

// @Summary 更新文章
// @Produce json
// @Param id path int true "文章id"
// @Param title body string false "文章名" maxlength(100)
// @Param tag_id body int false "标签"
// @Param desc body string fasle "简述"
// @Param cover_image_url  body string false "图片地址"
// @Param content body string fasle "内容"
// @Param modified_by body string true "修改人"
// @Success 200 {string} string "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "请求错误"
// @Router /api/v1/articles [put]
func (t Article) Update(c *gin.Context) {}

// @Summary 删除文章
// @Produce json
// @Param id path int true "文章id"
// @Param state body int fasle "状态"
// @Success 200 {string} string "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "请求错误"
// @Router /api/v1/articles [delete]
func (t Article) Delete(c *gin.Context) {}

/*************/
/***   model.CreateArticle 避免mutiple body 的问题***/
/*************/
/*

// @Summary 获取单独文章
// @Produce json
// @Param id path int true "文章id"
// @Success 200 {object} model.BlogArticle "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "请求错误"
// @Router /api/v1/articles/{id} [get]
func (t Article) Get(c *gin.Context) {}

// @Summary 获取多篇文章
// @Produce json
// @Param name query string false "文章名"
// @Param tag_id query int false "标签ID"
// @Param page query int false "页码"
// @Param paage_size query int false "每页数量"
// @Success 200 {object} model.BlogArticle "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "请求错误"
// @Router /api/v1/articles [get]
func (t Article) List(c *gin.Context) {}

// @Summary 创建文章
// @Produce json
// @Param requset body model.CreateArticle true "创建文章"
// @Success 200 {object} model.BlogArticle "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/articles [post]
func (t Article) Create(c *gin.Context) {
	db := global.DBEngine
	a := NewArticle()
	db.Limit(1).Find(&a)
	fmt.Println(a)
	c.JSON(http.StatusOK,
		map[string]string{"res": " ok"})
}

// @Summary 更新文章
// @Produce json
// @Param request body model.UpdateArticle true "更新文章"
// @Success 200 {string} string "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "请求错误"
// @Router /api/v1/articles [put]
func (t Article) Update(c *gin.Context) {}

// @Summary 删除文章
// @Produce json
// @Param id path int true "文章id"
// @Param state body int fasle "状态"
// @Success 200 {string} string "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "请求错误"
// @Router /api/v1/articles/{id} [delete]
func (t Article) Delete(c *gin.Context) {}


*/
