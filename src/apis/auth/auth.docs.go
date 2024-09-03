package auth

import (
	authDto "app/src/apis/auth/dtos"
	"app/src/shared/exception"
	httpContext "app/src/shared/http-context"
)

func initAuthDto(_ authDto.AuthResDto)    {}
func initException(_ exception.HttpError) {}

// @Summary Login
// @Description User login
// @Tags Auth
// @Accept  json
// @Produce  json
// @Param   loginReq  body  authDto.LoginReqDto  true  "Login request"
// @Success 200 {object} authDto.AuthResDto
// @Failure 422 {object} exception.HttpError
// @Failure 400 {object} exception.HttpError
// @Router /api/v1/auth/sign-in [post]
func login(ctx *httpContext.CustomContext) {}

// @Summary Register
// @Description User Register
// @Tags Auth
// @Accept  json
// @Produce  json
// @Param   registerReq  body  authDto.RegisterReqDto  true  "Register request"
// @Success 201 {object} authDto.AuthResDto
// @Failure 422 {object} exception.HttpError
// @Failure 400 {object} exception.HttpError
// @Router /api/v1/auth/sign-up [post]
func register(ctx *httpContext.CustomContext) {}

// @Summary Refresh Token
// @Description Refresh Token
// @Tags Auth
// @Produce json
// @Security BearerAuth
// @Success 200 {object} authDto.TokenResDto
// @Failure 401 {object} exception.HttpError
// @Router /api/v1/auth/refresh-token [get]
func refreshToken(ctx *httpContext.CustomContext) {}
