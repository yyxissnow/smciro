package http

import (
	"github.com/gin-gonic/gin"
	"github.com/yyxissnow/smicro/ecode"
	"net/http"
)

type SuccessResponse struct {
	Code int         `json:"code"`
	Data interface{} `json:"data,omitempty"`
}

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"msg"`
}

type PageSearch struct {
	Keyword string `json:"keyword"`
	Page    int    `json:"page"`
	Size    int    `json:"size"`
}

type ListResponse struct {
	Total int64       `json:"total"`
	Count int64       `json:"count"`
	List  interface{} `json:"list"`
}

func ResSuccess(c *gin.Context) {
	c.JSON(http.StatusOK, SuccessResponse{Code: ecode.Success.Code()})
}

func ResSuccessData(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, SuccessResponse{Code: ecode.Success.Code(), Data: data})
}

func ResError(c *gin.Context, err error) {
	e := ecode.AnalyseError(err)
	c.JSON(http.StatusOK, ErrorResponse{Code: e.Code(), Message: e.Message()})
}
