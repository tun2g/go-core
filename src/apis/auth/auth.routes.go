package auth

import (
	"app/src/shared/auth"
	httpContext "app/src/shared/http-context"
	"app/src/shared/jwt"

	"github.com/gin-gonic/gin"
)

func (authController *AuthController) InitRoute(
	routes *gin.RouterGroup,
	jwtAccessTokenManager *jwt.JWTManager,
	jwtRefreshTokenManager *jwt.JWTManager,
) {
	routes.POST(
		"/sign-in",
		httpContext.CustomContextHandler(authController.Login),
	)

	routes.POST(
		"/sign-up",
		httpContext.CustomContextHandler(authController.Register),
	)

	routes.GET(
		"/refresh-token",
		httpContext.CustomContextHandler(auth.TokenAuthMiddleware(jwtRefreshTokenManager)),
		httpContext.CustomContextHandler(authController.RefreshToken),
	)
}
