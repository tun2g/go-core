package logger

import (
	httpContext "app/src/shared/http-context"
	"fmt"

	"github.com/sirupsen/logrus"
)

func RequestLoggerMiddleware(ctx *httpContext.CustomContext) {
	logger := NewLogger("LoggerMiddleware")

	clientIP := ctx.ClientIP()
	method := ctx.Request.Method
	path := ctx.Request.URL.Path
	userAgent := ctx.Request.UserAgent()
	dataLength := ctx.Writer.Size()

	entry := logger.WithFields(logrus.Fields{
		"request_id": ctx.GetRequestId(),
		"client_ip":   clientIP,
		"method":      method,
		"path":        path,
		"user_agent":  userAgent,
		"data_length": dataLength,
	})

	if len(ctx.Errors) > 0 {
		entry.Error(ctx.Errors.String())
	}
	entry.Info(fmt.Sprintf("----------Request received: %s", ctx.GetRequestId()))

	user := ctx.GetUser()

	if(user != nil){
		context := logger.WithFields(logrus.Fields{
			"user_id": user.Id,
			"email": user.Email,
			"role": user.Role,
			"full_name": user.FullName,
		})
		context.Info("----------Context before is")
	} else{
		context := logger.WithField("request_id", ctx.GetRequestId())
		context.Info("----------Context before is")
	}

}

func ResponseLoggerMiddleware(ctx *httpContext.CustomContext) {
	logger := NewLogger("LoggerMiddleware")

	ctx.Next()

	statusCode := ctx.Writer.Status()
	clientIP := ctx.ClientIP()
	method := ctx.Request.Method
	path := ctx.Request.URL.Path
	userAgent := ctx.Request.UserAgent()
	dataLength := ctx.Writer.Size()

	entry := logger.WithFields(logrus.Fields{
		"request_id": ctx.GetRequestId(),
		"status_code": statusCode,
		"client_ip":   clientIP,
		"method":      method,
		"path":        path,
		"user_agent":  userAgent,
		"data_length": dataLength,
	})

	user := ctx.GetUser()

	if(user != nil){
		context := logger.WithFields(logrus.Fields{
			"user_id": user.Id,
			"email": user.Email,
			"role": user.Role,
			"full_name": user.FullName,
		})
		context.Info("----------Context after is")
	} else{
		context := logger.WithField("request_id", ctx.GetRequestId())
		context.Info("----------Context after is")
	}

	if len(ctx.Errors) > 0 {
		entry.Errorf("----------Request completed: due to %s in %s", ctx.Errors.String(), ctx.HandlerName())
	} else{
		entry.Info(fmt.Sprintf("----------Request completed: %s", ctx.GetRequestId()))
	}
}
