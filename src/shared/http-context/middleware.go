package httpContext

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	httpContextConstant "app/src/shared/http-context/constants"
)

func HttpContextMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		customContext := &CustomContext{Context: ctx}
		requestId := ctx.GetHeader(httpContextConstant.X_REQUEST_ID)

		if requestId == "" {
			requestId = uuid.New().String()
		}

		customContext.SetRequestId(requestId)

		ctx.Next()
	}
}
