package app

import (
	"net/http"

	"github.com/consult98/simple-blog-go/pkg/errcode"
	"github.com/gin-gonic/gin"
)

//app是对返回的相应进行封装

type Response struct {
	Ctx *gin.Context //gin.Context中实现了对request和response的封装，包括请求的解析和相应的写入等
}

type Pager struct {
	Page      int `json:"page"`
	PageSize  int `json:"page_size"`
	TotalRows int `json:"total_rows"`
}

//实例化一个Response
func NewResponse(ctx *gin.Context) *Response {
	return &Response{Ctx: ctx}
}

//ToResponse 请求成功，返回响应的数据
func (r *Response) ToResponse(data interface{}) {
	if data == nil {
		data = gin.H{}
	}
	r.Ctx.JSON(http.StatusOK, data)
}

//请求成功，使用JSON返回gin前端请求中的参数（核心字段）
func (r *Response) ToResponseList(list interface{}, totalRows int) {
	r.Ctx.JSON(http.StatusOK, gin.H{
		"list": list,
		"pager": Pager{
			Page:      GetPage(r.Ctx),
			PageSize:  GetPageSize(r.Ctx),
			TotalRows: totalRows,
		},
	})
}

//请求失败，返回错误信息
func (r *Response) ToErrorResponse(err *errcode.Error) {
	response := gin.H{"code": err.Code(), "msg": err.Msg()}
	details := err.Details()
	if len(details) > 0 {
		response["details"] = details
	}

	r.Ctx.JSON(err.StatusCode(), response)
}
