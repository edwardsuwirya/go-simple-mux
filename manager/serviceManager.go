package manager

import "gosimplemux/useCases"

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
func NewServiceManger() ServiceManager {
	return &serviceManager{repo: NewRepoManager()}
}
