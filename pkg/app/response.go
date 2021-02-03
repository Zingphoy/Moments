package app

import (
	"Moments/pkg/hint"

	"github.com/gin-gonic/gin"
)

type GinCtx struct {
	C *gin.Context
}

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func (C *GinCtx) MakeJsonRes(httpCode int, statusCode int, data interface{}) {
	C.C.JSON(httpCode, Response{
		Code: statusCode,
		Msg:  hint.GetHintMsg(statusCode),
		Data: data,
	})
}
