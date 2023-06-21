// 分页处理方法
// 页面数
package app

import (
	"github.com/gin-gonic/gin"
	"github.com/lc-1010/OneBlogService/convert"
	"github.com/lc-1010/OneBlogService/global"
)

func GetPage(c *gin.Context) int {
	page := convert.StrTo(c.Query("page")).MustInt()
	if page <= 0 {
		return 1
	}
	return page
}

func GetPageSize(c *gin.Context) int {
	pageSize := convert.StrTo(c.Query("page_size")).MustInt()
	if pageSize <= 0 {
		return global.AppSetting.DefaultPageSize
	}
	if pageSize > global.AppSetting.MaxPageSize {
		return global.AppSetting.MaxPageSize
	}

	return pageSize
}

func GetPageOffset(page, pageSize int) int {
	reuslt := 0
	if page > 0 {
		reuslt = (page - 1) * pageSize
	}

	return reuslt
}

/************************************/
/***********  ***********/
/************************************/
