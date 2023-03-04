package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-pay/gopher/ecode"
)

type CommonRsp struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func JSON(c *gin.Context, data interface{}, err error) {
	e := ecode.AnalyseError(err)
	if e == ecode.Success && ecode.Success != nil {
		e = ecode.Success
	}
	rsp := &CommonRsp{
		Code:    e.Code(),
		Message: e.Message(),
		Data:    data,
	}
	c.JSON(http.StatusOK, rsp)
}
