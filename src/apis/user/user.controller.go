package user

import (
	userService "app/src/apis/user/services"
	pageDto "app/src/shared/dto"
	"app/src/shared/exception"
	httpContext "app/src/shared/http-context"
	sharedConstants "app/src/shared/constants"
	"app/src/shared/utils"
	"net/http"
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
	queryDto := utils.GetValidation[pageDto.PageOptionsDto](ctx, sharedConstants.QUERY)
	
	queryDto = pageDto.NewPageOptionsDto(queryDto);
	data, err := ctl.userService.GetUsers(ctx, queryDto);

	if(err!=nil){
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, data)
}
