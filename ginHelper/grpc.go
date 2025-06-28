package ginHelper

import (
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
)

// 处理从grpc的错误到http的错误的相关函数
// 不过这样的话日志信息就无法打印了 --- 前置打印就行

func ErrorFromGrpcToHttp(context *gin.Context, s *status.Status) {
	//// 1. 从err中提取grpc的错误信息
	//s, ok := status.FromError(err)
	//if !ok {
	//	// 没有提取到错误
	//	Fail(context, http.StatusInternalServerError, -1, "系统错误")
	//}
	// 根据对应错误码返回对应信息
	switch s.Code() {
	case codes.InvalidArgument:
		Fail(context, http.StatusBadRequest, -1, s.Message())
	case codes.NotFound:
		Fail(context, http.StatusNotFound, -1, s.Message())
	case codes.Internal:
		Fail(context, http.StatusInternalServerError, -1, "系统错误")
	default:
		Fail(context, http.StatusInternalServerError, -1, "其它错误")
	}
}
