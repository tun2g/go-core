package middlewares

import (
	"app/src/shared/exception"
	"app/src/shared/constants"
	httpContext "app/src/shared/http-context"
)

func ValidateMiddleware[T comparable](valType constants.ValidationType) func(ctx *httpContext.CustomContext) {
	return func(ctx *httpContext.CustomContext) {
		var r T

		switch valType {
		case constants.BODY:
			if err := ctx.ShouldBindJSON(&r); err != nil {
				exception.ThrowUnprocessableEntityException(ctx, err)
				return
			}
		case constants.QUERY:
			if err := ctx.ShouldBindQuery(&r); err != nil {
				exception.ThrowUnprocessableEntityException(ctx, err)
				return
			}
		case constants.PARAM:
			if err := ctx.ShouldBindUri(&r); err != nil {
				exception.ThrowUnprocessableEntityException(ctx, err)
				return
			}
		}
		ctx.SetValidation(r, valType)
		ctx.Next()
	}
}
