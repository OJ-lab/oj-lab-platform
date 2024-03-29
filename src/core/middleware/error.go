package middleware

import (
	"fmt"

	"github.com/OJ-lab/oj-lab-services/src/core"
	"github.com/gin-gonic/gin"
)

func GetServiceError(ginErr gin.Error) *core.SeviceError {
	if core.IsServiceError(ginErr.Meta) {
		return ginErr.Meta.(*core.SeviceError)
	} else {
		serviceErr := core.NewInternalError(fmt.Sprintf("%v", ginErr.Err))
		serviceErr.CaptureStackTrace()
		return serviceErr
	}
}

func HandleError(ginCtx *gin.Context) {
	ginCtx.Next()

	errCount := len(ginCtx.Errors)
	if errCount > 0 {
		core.AppLogger().Errorf("Last error from GIN middleware: %+v", ginCtx.Errors[errCount-1].Err)
		err := GetServiceError(*ginCtx.Errors[errCount-1])
		ginCtx.JSON(err.Code, gin.H{
			"code": err.Code,
			"msg":  err.Msg,
		})
	}
}
