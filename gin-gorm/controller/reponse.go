/**
 * @Author: yangyi
 * @Description:
 * @File:  reponse
 * @Version: 1.0.0
 * @Date: 2022/6/14 下午5:03
 */

package controller

import (
    "github.com/gin-gonic/gin"
    //"gin-gorm/controller"
    "net/http"
)

type ResponseData struct {
    Code    MyCode      `json:"code"`
    Message string      `json:"message"`
    Data    interface{} `json:"data,omitempty"`	// omitempty当data为空时,不展示这个字段
}

func ResponseError(ctx *gin.Context, c MyCode) {
    rd := &ResponseData{
        Code:    c,
        Message: c.Msg(),
        Data:    nil,
    }
    ctx.JSON(http.StatusOK, rd)
}

func ResponseErrorWithMsg(ctx *gin.Context, code MyCode, data interface{}) {
    rd := &ResponseData{
        Code:    code,
        Message: code.Msg(),
        Data:    nil,
    }
    ctx.JSON(http.StatusOK, rd)
}

func ResponseSuccess(ctx *gin.Context, data interface{}) {
    rd := &ResponseData{
        Code:    CodeSuccess,
        Message: CodeSuccess.Msg(),
        Data:    data,
    }
    ctx.JSON(http.StatusOK, rd)
}
