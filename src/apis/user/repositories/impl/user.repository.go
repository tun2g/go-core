package user

import (
	userModel "app/src/apis/user/models"
	userRepository "app/src/apis/user/repositories"
	commonRepository "app/src/shared/repository"

	"gorm.io/gorm"
)

type UserRepository struct {
	commonRepository.BaseRepository[userModel.User]
}

func NewUserRepository(storage *gorm.DB) userRepository.IUserRepository {
	return &UserRepository{
		BaseRepository: commonRepository.BaseRepository[userModel.User]{DB: storage},
	}
}
