package httpContext

import (
	"app/src/shared/dto"
	httpContextConstants "app/src/shared/http-context/constants"

	"github.com/gin-gonic/gin"
)
type CustomContext struct {

	*gin.Context
}

func (ctx *CustomContext) GetUser() *dto.CurrentUser {
	user, exists := ctx.Get(httpContextConstants.AUTH_USER)
	if !exists {
		return nil
	}
	if userStruct, ok := user.(*dto.CurrentUser); ok {
		return userStruct
	}
	return nil
}

func (ctx *CustomContext) SetUser(user *dto.CurrentUser){
	ctx.Set(httpContextConstants.AUTH_USER, user)
}

func (ctx *CustomContext) SetRequestId(requestId string){
	ctx.Set(httpContextConstants.REQUEST_ID, requestId)
}

func (ctx * CustomContext) GetRequestId() string {
	requestId:= ctx.GetString(httpContextConstants.REQUEST_ID)
	return requestId
}

func (ctx *CustomContext) SetIsPublicRoute(){
	ctx.Set(httpContextConstants.IS_PUBLIC, true)
}

func (ctx *CustomContext) GetIsPublicRoute() bool{
	isPublic := ctx.GetBool(httpContextConstants.IS_PUBLIC)
	return isPublic
}

