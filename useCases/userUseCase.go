package useCases

import (
	"gosimplemux/models"
	"gosimplemux/repositories"
)

type IUserUseCase interface {
	Register(newUser *models.User) error
	GetUserInfo(id string) *models.User
	Unregister(id string) error
	UpdateInfo(id string, newUser *models.User) error
}

type UserUseCase struct {
	userRepo repositories.UserRepository
}

func NewUserUseCase(userRepo repositories.UserRepository) IUserUseCase {
	return &UserUseCase{
		userRepo,
	}
}

func (uc *UserUseCase) Register(newUser *models.User) error {
	return uc.userRepo.Create(newUser)
}

func (uc *UserUseCase) GetUserInfo(id string) *models.User {
	return uc.userRepo.FindOneById(id)
}

func (uc *UserUseCase) UpdateInfo(id string, newUser *models.User) error {
	return uc.userRepo.Update(id, newUser)
}

func (uc *UserUseCase) Unregister(id string) error {
	return uc.userRepo.Delete(id)
}
