package user

import (
	authDto "app/src/apis/auth/dtos"
	userModel "app/src/apis/user/models"
	pageDto "app/src/shared/dto"
	"app/src/shared/exception"
	httpContext "app/src/shared/http-context"
)

func initUserModel(_ userModel.User)       {}
func initException(_ exception.HttpError)  {}
func initPageDto(_ pageDto.PageOptionsDto) {}
func initAuthDto(_ authDto.UserResDto)     {}

// @Summary Get me
// @Description Get me
// @Tags User
// @Produce json
// @Security BearerAuth
// @Success 200 {object} authDto.UserResDto
// @Failure 401 {object} exception.HttpError
// @Router /users/me [get]
func getMe(ctx *httpContext.CustomContext) {}

// @Summary Get all users
// @Description Get all users
// @Tags User
// @Param pageOptions query pageDto.PageOptionsDto true "Pagination and ordering options"
// @Produce json
// @Security BearerAuth
// @Success 200 {array} authDto.UserResDto
// @Failure 422 {object} exception.HttpError
// @Failure 401 {object} exception.HttpError
// @Failure 403 {object} exception.HttpError
// @Router /users [get]
func getUsers(ctx *httpContext.CustomContext) {}
