package user

import (
	httpContext "app/src/shared/http-context"
	commonDto "app/src/shared/dto"
	userModel "app/src/apis/user/models"
)

type IUserService interface {
	GetMe(ctx *httpContext.CustomContext) *commonDto.CurrentUser
	GetUsers(ctx *httpContext.CustomContext, dto *commonDto.PageOptionsDto) (*commonDto.PageDto[userModel.User], error)
}