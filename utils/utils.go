package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type serviceError struct {
	Error  string `json:"error"`
	Detail string `json:"detail"`
}

func InternalServiceError(ctx *gin.Context, err error) {
	errStruct := serviceError{
		Error:  "internal service error",
		Detail: err.Error(),
	}
	ctx.IndentedJSON(http.StatusInternalServerError, errStruct)
}

func ValidationError(ctx *gin.Context, err error) {
	errStruct := serviceError{
		Error:  "validation error",
		Detail: err.Error(),
	}
	ctx.IndentedJSON(http.StatusBadRequest, errStruct)
}

func BindObjToContext(ctx *gin.Context, obj any) {
	ctx.IndentedJSON(http.StatusOK, obj)
}

func AcceptAllHosts(ctx *gin.Context) {
	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.Header("Access-Control-Allow-Credentials", "true")
	ctx.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, X-Session-Key, Set-Cookie")
	ctx.Header("Access-Control-Expose-Headers", "Set-Cookie")
}

func CountParamLength(params ...any) int {
	for i := range params {
		if params[i] == nil {
			return 0
		}
	}
	return len(params)
}

func CountParamLengthUInt(params ...any) int {
	for i := range params {
		if params[i].([]uint8) == nil {
			return 0
		}
	}
	return len(params)
}
