package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Success struct {
	Code int         `json:"code"`
	Data interface{} `json:"data,omitempty"`
}

type Error struct {
	Code    ErrorCode `json:"code"`
	Message string    `json:"msg"`
}

type PageSearch struct {
	Keyword string
	Page    int
	Size    int
}

type ListResponse struct {
	Total int64       `json:"total"`
	Count int64       `json:"count"`
	List  interface{} `json:"list"`
}

func ResSuccess(c *gin.Context) {
	c.JSON(http.StatusOK, Success{Code: successCode})
}

func ResSuccessData(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Success{Code: successCode, Data: data})
}

func ResError(c *gin.Context, code ErrorCode) {
	c.JSON(http.StatusOK, Error{Code: code, Message: getErrMessage(code)})
}
