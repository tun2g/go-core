package auth

import (
	"net/http"

	authService "app/src/apis/auth/services"
	userModel "app/src/apis/user/models"
	httpContext "app/src/shared/http-context"
	sharedConstants "app/src/shared/constants"
	"app/src/shared/utils"
	dto "app/src/apis/auth/dtos"
)

type AuthController struct {
	authService authService.IAuthService
}

func NewAuthController(authService authService.IAuthService) *AuthController {
	return &AuthController{
		authService: authService,
	}
}

func (ctl *AuthController) Login(ctx *httpContext.CustomContext) {
	var err error
	var user *userModel.User
	var tokens *dto.TokenResDto

	reqDto := utils.GetValidation[dto.LoginReqDto](ctx, sharedConstants.BODY)

	user, tokens, err = ctl.authService.Login(reqDto, ctx)
	if err != nil {
		ctx.Error(err)
		return
	}

	userRes := dto.UserResDto{
		ID:       user.Id,
		Email:    user.Email,
		FullName: user.FullName,
	}

	authRes := dto.AuthResDto{
		User:   userRes,
		Tokens: *tokens,
	}

	ctx.JSON(http.StatusOK, authRes)
}

func (ctl *AuthController) Register(ctx *httpContext.CustomContext) {
	var err error
	var user *userModel.User
	var tokens *dto.TokenResDto

	reqDto := utils.GetValidation[dto.RegisterReqDto](ctx, sharedConstants.BODY)

	user, tokens, err = ctl.authService.Register(reqDto, ctx)

	if err != nil {
		ctx.Error(err)
		return
	}

	userRes := dto.UserResDto{
		ID:       user.Id,
		Email:    user.Email,
		FullName: user.FullName,
	}

	authRes := dto.AuthResDto{
		User:   userRes,
		Tokens: *tokens,
	}

	ctx.JSON(http.StatusCreated, authRes)
}

func (ctl *AuthController) RefreshToken(ctx *httpContext.CustomContext) {
	tokens, err := ctl.authService.RefreshToken(ctx)

	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, *tokens)
}
