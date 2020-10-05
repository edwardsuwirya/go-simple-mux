package manager

import (
	"gosimplemux/infra"
	"gosimplemux/useCases"
)

type ServiceManager interface {
	UserUseCase() useCases.IUserUseCase
	UserAuthUseCase() useCases.IUserAuthUseCase
}

type serviceManager struct {
	repo RepoManager
}

func (sm *serviceManager) UserUseCase() useCases.IUserUseCase {
	return useCases.NewUserUseCase(sm.repo.UserRepo())
}

func (sm *serviceManager) UserAuthUseCase() useCases.IUserAuthUseCase {
	return useCases.NewUserAuthUseCase(sm.repo.UserAuthRepo(), sm.repo.UserRepo())
}
func NewServiceManger(infra infra.Infra) ServiceManager {
	return &serviceManager{repo: NewRepoManager(infra)}
}
