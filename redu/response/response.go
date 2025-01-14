package response

import (
	"github.com/gin-gonic/gin"
	"redu/global"
	"time"

	"net/http"
)

type response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Cost string      `json:"cost,omitempty"`
	Data interface{} `json:"data,omitempty"`
}

type listResponse[T any] struct {
	List  T     `json:"list"`
	Count int64 `json:"count"`
}

func FailWithMsg(c *gin.Context, code int, msg string) {
	message := (*global.ResMap)[code]
	if msg != "" {
		message = msg
	}
	c.JSON(http.StatusOK, response{
		Code: code,
		Msg:  message,
	})
}

func OkWithMsg(c *gin.Context, msg string) {
	code := SUCCESS
	message := (*global.ResMap)[code]
	if msg != "" {
		message = msg
	}
	c.JSON(http.StatusOK, response{
		Code: code,
		Msg:  message,
		Cost: time.Since(c.GetTime("requestStartTime")).String(),
	})
}

func OkWithData(c *gin.Context, data interface{}) {
	code := SUCCESS
	msg := (*global.ResMap)[code]
	c.JSON(http.StatusOK, response{
		Code: code,
		Msg:  msg,
		Data: data,
		Cost: time.Since(c.GetTime("requestStartTime")).String(),
	})
}

func OkWithList(c *gin.Context, list any, count int64) {
	OkWithData(c, listResponse[any]{
		List:  list,
		Count: count,
	})
}
