package utils

import (
	"net/http"
	"fmt"
	httpContext "app/src/shared/http-context"
	"app/src/shared/constants"
	"app/src/shared/exception"
)

// Error: method must have no type parameters
// So we cannot implement this method in the struct HttpContext
func GetValidation[T comparable](ctx *httpContext.CustomContext, valType constants.ValidationType) (*T){
	data, exists := ctx.Get(fmt.Sprintf("%d", valType))
	if(!exists){
		ctx.AbortWithStatusJSON(
			http.StatusInternalServerError,
			exception.HttpError{
				Message: "Internal Server Error",
				RequestId: ctx.GetRequestId(),
			},
		)
	}

	validData, ok := data.(T)
	if(!ok){
		ctx.AbortWithStatusJSON(
			http.StatusInternalServerError,
			exception.HttpError{
				Message: "Internal Server Error",
				RequestId: ctx.GetRequestId(),
			},
		)
	}

	return &validData
}
