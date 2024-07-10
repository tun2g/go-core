package user

import (
	"net/http"
	"app/src/shared/exception"
	userService "app/src/apis/user/services"
	httpContext "app/src/shared/http-context"
	pageDto "app/src/shared/dto"
)

type UserController struct {
	userService userService.IUserService
}

func NewUserController(userService userService.IUserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (ctl *UserController) GetMe(ctx *httpContext.CustomContext) {
	user := ctl.userService.GetMe(ctx)

	if user == nil {
		ctx.Error(exception.NewBadRequestException(
			ctx.GetRequestId(),
			[]exception.ErrorDetail{{}},
		))
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func (ctl *UserController) GetUsers(ctx *httpContext.CustomContext) {
	var queryDto pageDto.PageOptionsDto
	if err := ctx.ShouldBindQuery(&queryDto); err != nil {
		ctx.Error(exception.NewUnprocessableEntityException(ctx.GetRequestId(), err))
		return
	}

	queryDto = *pageDto.NewPageOptionsDto(&queryDto);
	data, err := ctl.userService.GetUsers(ctx, &queryDto);

	if(err!=nil){
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, data)
}
