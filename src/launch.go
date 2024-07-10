package src

import (
	"app/src/shared/jwt"
	"app/src/shared/utils"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"

	authController "app/src/apis/auth"
	userController "app/src/apis/user"

	authService "app/src/apis/auth/services/impl"
	userService "app/src/apis/user/services/impl"

	userRepository "app/src/apis/user/repositories/impl"
)

func (server *Server) launchingServer(route *gin.Engine) {
	var jwtAccessTokenManager = jwt.NewJWTManager(
		server.config.JwtAccessTokenSecret,
		time.Duration(server.config.JwtAccessTokenExpirationTime),
	)

	var jwtRefreshTokenManager = jwt.NewJWTManager(
		server.config.JwtRefreshTokenSecret,
		time.Duration(server.config.JwtRefreshTokenExpirationTime),
	)

	var bcrypt = utils.NewBcryptEncoder(bcrypt.DefaultCost)

	// Repository
	var userRepository = userRepository.NewUserRepository(server.db)

	// Service
	var _authService = authService.NewAuthService(
		&userRepository,
		&jwtAccessTokenManager,
		&jwtRefreshTokenManager,
		&bcrypt,
	)

	var _userService = userService.NewUserService(
		&userRepository,
	)

	// Controller
	var _authController = authController.NewAuthController(_authService)
	authRoutes := route.Group("/auth")
	_authController.InitRoute(authRoutes, &jwtAccessTokenManager, &jwtRefreshTokenManager)

	var _userController = userController.NewUserController(_userService)
	userRoutes := route.Group("/users")
	_userController.InitRoute(userRoutes, &jwtAccessTokenManager)
}
