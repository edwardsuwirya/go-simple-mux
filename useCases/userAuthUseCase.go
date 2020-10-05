package useCases

import (
	"gosimplemux/models"
	"gosimplemux/repositories"
)

type IUserAuthUseCase interface {
	UserNamePasswordValidation(userName string, password string) *models.User
}

type UserAuthUseCase struct {
	userAuthRepo repositories.UserAuthRepository
	userRepo     repositories.UserRepository
}

func NewUserAuthUseCase(userAuthRepo repositories.UserAuthRepository, userRepo repositories.UserRepository) IUserAuthUseCase {
	return &UserAuthUseCase{
		userAuthRepo, userRepo,
	}
}

func (uc *UserAuthUseCase) UserNamePasswordValidation(userName string, password string) *models.User {
	userAuth := uc.userAuthRepo.FindOneByUserNameAndPassword(userName, password)
	if userAuth != nil {
		userInfo := uc.userRepo.FindOneById(userAuth.UserRegId)
		return userInfo
	}
	return nil
}
