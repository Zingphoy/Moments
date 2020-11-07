package app

import (
	"Moments/pkg/hints"

	"github.com/gin-gonic/gin"
)

type GinCtx struct {
	C *gin.Context
}

type Respons struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func (C *GinCtx) MakeJsonRes(httpCode int, statusCode int, data interface{}) {
	C.C.JSON(httpCode, Respons{
		Code: statusCode,
		Msg:  hints.GetHintMsg(statusCode),
		Data: data,
	})
}
