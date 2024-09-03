package user

import (
	"app/src/shared/auth"
	authConstants "app/src/shared/auth/constants"
	httpContext "app/src/shared/http-context"
	"app/src/shared/jwt"
	"app/src/shared/middlewares"

	"github.com/gin-gonic/gin"

	pageDto "app/src/shared/dto"
	shareConstants "app/src/shared/constants"

)

func (userController *UserController) InitRouteV1(
	routes *gin.RouterGroup,
	jwtAccessTokenManager *jwt.JWTManager,
) {
	routes.GET(
		"/me",
		httpContext.CustomContextHandler(auth.TokenAuthMiddleware(jwtAccessTokenManager)),
		httpContext.CustomContextHandler(userController.GetMe),
	)

	routes.GET(
		"",
		httpContext.CustomContextHandler(auth.TokenAuthMiddleware(jwtAccessTokenManager)),
		httpContext.CustomContextHandler(auth.RoleAuthMiddleware([]authConstants.Role{authConstants.RoleAdmin})),
		httpContext.CustomContextHandler(middlewares.ValidateMiddleware[pageDto.PageOptionsDto](shareConstants.QUERY)),
		httpContext.CustomContextHandler(userController.GetUsers),
	)
}
