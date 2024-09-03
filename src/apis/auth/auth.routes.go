package auth

import (
	sharedAuth "app/src/shared/auth"
	httpContext "app/src/shared/http-context"
	"app/src/shared/jwt"
	"app/src/shared/middlewares"
	sharedConstants "app/src/shared/constants"

	dto "app/src/apis/auth/dtos"

	"github.com/gin-gonic/gin"
)

func (authController *AuthController) InitRouteV1(
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
		httpContext.CustomContextHandler(middlewares.ValidateMiddleware[dto.RegisterReqDto](sharedConstants.BODY)),
		httpContext.CustomContextHandler(authController.Register),
	)

	routes.GET(
		"/refresh-token",
		httpContext.CustomContextHandler(sharedAuth.TokenAuthMiddleware(jwtRefreshTokenManager)),
		httpContext.CustomContextHandler(authController.RefreshToken),
	)
}
