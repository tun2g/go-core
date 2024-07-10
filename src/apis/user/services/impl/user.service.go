package user

import (
	userModel "app/src/apis/user/models"
	userRepository "app/src/apis/user/repositories"
	userService "app/src/apis/user/services"
	commonDto "app/src/shared/dto"
	httpContext "app/src/shared/http-context"
	"app/src/shared/utils"
)

type UserService struct {
	userRepository userRepository.IUserRepository
}

func NewUserService(userRepository *userRepository.IUserRepository) userService.IUserService {
	return &UserService{
		userRepository: *userRepository,
	}
}

func (srv *UserService) GetMe(ctx *httpContext.CustomContext) *commonDto.CurrentUser {

	user := ctx.GetUser()

	return user
}

func (srv *UserService) GetUsers(
	ctx *httpContext.CustomContext, 
	dto *commonDto.PageOptionsDto,
	) (*commonDto.PageDto[userModel.User], error) {
	entities, total, err := srv.userRepository.Paging(dto)
	if err != nil {
		return nil, err
	}

	pageRes := utils.GeneratePaginationResult[userModel.User](total, *entities, dto)

	return pageRes, nil
}
