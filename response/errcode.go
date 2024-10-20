package response

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code int
	Msg  string
	Data interface{}
}

type Error struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

var codes = map[int]string{}

func NewError(code int, msg string) *Error {
	if _, ok := codes[code]; ok {
		panic(fmt.Sprintf("错误码 %d 已存在，请更换一个", code))
	}
	codes[code] = msg
	return &Error{Code: code, Msg: msg}
}

func Result(code int, data interface{}, context *gin.Context) {
	context.JSON(http.StatusOK, Response{Code: code, Msg: codes[code], Data: data})
}
