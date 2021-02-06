package app

import (
	"Moments/pkg/hint"
	"Moments/pkg/log"
	"fmt"
	"net/http"

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

func (C *GinCtx) MakeSuccessJsonRes(httpCode int, data interface{}) {
	code := hint.SUCCESS
	C.C.JSON(httpCode, Response{
		Code: code,
		Msg:  hint.GetHintMsg(code),
		Data: data,
	})
}

func (C *GinCtx) MakeFailedJsonRes(httpCode int, data interface{}) {
	if err, ok := data.(hint.CustomError); ok {
		// return 5xx http code if it's database error, no matter what httpCode param is
		if hint.DATABASE_ERROR <= err.Code && err.Code <= 2999 {
			log.Error(C.C, "database error:", err.Err.Error())
			C.C.JSON(http.StatusInternalServerError, Response{
				Code: err.Code,
				Msg:  hint.GetHintMsg(err.Code),
				Data: err.Err.Error(),
			})
			return
		}

		if err.Code == hint.INVALID_PARAM {
			log.Info(C.C, fmt.Sprintf("data parse json error, params are: %v", C.C.Params))
		}

		C.C.JSON(httpCode, Response{
			Code: err.Code,
			Msg:  hint.GetHintMsg(err.Code),
			Data: err.Err.Error(),
		})
	}
}

func (C *GinCtx) MakeJsonRes(httpCode int, statuCode int, data interface{}) {
	C.C.JSON(httpCode, Response{
		Code: statuCode,
		Msg:  hint.GetHintMsg(statuCode),
		Data: data,
	})
}
