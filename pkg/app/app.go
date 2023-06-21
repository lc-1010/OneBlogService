// 返回处理 c.JSON respone
// 内容列表整合 pagination
package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lc-1010/OneBlogService/pkg/errcode"
)

type Response struct {
	Ctx *gin.Context
}

type Pager struct {
	Page      int `json:"page,omitempty"`
	PageSize  int `json:"page_size,omitempty"`
	TotalRows int `json:"total_rows,omitempty"`
}

func NewResponse(ctx *gin.Context) *Response {
	return &Response{Ctx: ctx}
}

func (r *Response) ToResponse(data any) {
	if data == nil {
		data = gin.H{}
	}
	r.Ctx.JSON(http.StatusOK, data)
}

func (r *Response) ToResponseList(list any, totalRows int) {
	r.Ctx.JSON(http.StatusOK, gin.H{
		"list": list,
		"pager": Pager{
			Page:      GetPage(r.Ctx),
			PageSize:  GetPageSize(r.Ctx),
			TotalRows: totalRows,
		},
	})
}

func (r *Response) ToErrorResponse(err *errcode.Error) {
	response := gin.H{"code": err.Code(), "msg": err.Msg()}
	details := err.Details()
	if len(details) > 0 {
		response["details"] = details
	}
	r.Ctx.JSON(err.StatusCode(), response)
}
