package ginHelper

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type BaseResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewBaseResponse(code int, message string, data interface{}) BaseResponse {
	return BaseResponse{
		Code:    code,
		Message: message,
		Data:    data,
	}
}

func Success(context *gin.Context, data interface{}) {
	context.JSON(http.StatusOK, NewBaseResponse(0, "success", data))
}

func Fail(context *gin.Context, statusCode, code int, data interface{}) {
	context.JSON(statusCode, NewBaseResponse(code, "fail", data))
}
