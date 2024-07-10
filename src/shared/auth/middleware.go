package auth

import (
	authConstants "app/src/shared/auth/constants"
	httpContext "app/src/shared/http-context"
	"app/src/shared/dto"
	"app/src/shared/exception"
	"app/src/shared/jwt"
	"app/src/shared/utils"
	"strings"
)

func IsPublicRouteMiddleware() func(ctx *httpContext.CustomContext){
	return func(ctx *httpContext.CustomContext){
		ctx.SetIsPublicRoute()
	}
}

func TokenAuthMiddleware(jwtManager *jwt.JWTManager) func(ctx *httpContext.CustomContext) {
	return func(ctx *httpContext.CustomContext) {

		isPublic := ctx.GetIsPublicRoute()
		var accessToken string
		authorizationHeader := ctx.GetHeader(authConstants.AUTHORIZATION_HEADER_KEY)
		
		if len(authorizationHeader) == 0 && isPublic == false {
			exception.ThrowUnauthorizedException(ctx)
			return
		}
		
		fields := strings.Fields(authorizationHeader)
		if len(fields) < 2 && isPublic == false{
			exception.ThrowUnauthorizedException(ctx)
			return
		}

		authorizationType := strings.ToLower(fields[0])
		if authorizationType != authConstants.AUTHORIZATION_BEARER_KEY && isPublic == false {
			exception.ThrowUnauthorizedException(ctx)
			return
		}
		accessToken = fields[1]

		payload, err := jwtManager.VerifyToken(accessToken)
		if err != nil && isPublic == false{
			exception.ThrowUnauthorizedException(ctx)
			return
		}

		currentUser := dto.NewCurrentUser(payload)

		ctx.SetUser(currentUser)
		ctx.Next()
	}
}

func RoleAuthMiddleware(roles []authConstants.Role) func(ctx *httpContext.CustomContext){
	return func(ctx *httpContext.CustomContext) {
		user := ctx.GetUser()
		if(user == nil){
			exception.ThrowForbiddenException(ctx)
			return
		}

		if(!utils.IsContains(roles, user.Role)){
			exception.ThrowForbiddenException(ctx)
			return
		}
		ctx.Next()
	}
}