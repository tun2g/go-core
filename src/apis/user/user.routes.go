package user

import (
	"app/src/shared/auth"
	authConstants "app/src/shared/auth/constants"
	httpContext "app/src/shared/http-context"
	"app/src/shared/jwt"

	"github.com/gin-gonic/gin"
)

func (userController *UserController) InitRoute(
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
		httpContext.CustomContextHandler(userController.GetUsers),
	)
}
