package utils

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"go-project-frame/pkg/globalError"
	"net/http"
)

type ResponseCode int

type Response struct {
	Code    ResponseCode `json:"code"`
	Msg     string       `json:"msg"`
	RealErr string       `json:"real_err"`
	Data    interface{}  `json:"data"`
}

// ResponseSuccess 响应成功的公共处理，内部会封装好响应体，设置给context
func ResponseSuccess(ctx *gin.Context, data interface{}) {
	resp := &Response{Code: http.StatusOK, Msg: "", Data: data}
	msg, ok := data.(string)
	if ok && msg == "" {
		resp.Msg = "操作成功"
	}

	// 即使发生错误，也只是在请求的response中展示，请求本身的响应状态码还应该设置为200
	ctx.JSON(http.StatusOK, resp)
	// 将resp对象，序列化成json格式
	response, _ := json.Marshal(resp)
	// 将封装好的响应体response，设置给 上下文对象Context，后面就可以直接返回给前端了
	ctx.Set("response", response)
}

// ResponseError 响应失败的公共处理，内部会封装好响应体，设置给context
func ResponseError(ctx *gin.Context, err error) {
	// 将错误转成我们自定义的错误类型，获取code的具体值
	var code ResponseCode
	myError := new(globalError.GlobalError)
	if errors.As(err, &myError) {
		code = ResponseCode(myError.Code)
	}
	resp := &Response{Code: code, Msg: err.Error(), RealErr: myError.RealErrorMessage, Data: ""}

	// 即使发生错误，也只是在请求的response中展示，请求本身的响应状态码还应该设置为200
	ctx.JSON(http.StatusOK, resp)
	// 将resp对象，序列化成json格式
	response, _ := json.Marshal(resp)
	// 将封装好的响应体response，设置给 上下文对象Context，后面就可以直接返回给前端了
	ctx.Set("response", response)
}
