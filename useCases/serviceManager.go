package useCases

type ServiceManager interface {
	UserUseCase() IUserUseCase
}

type serviceManager struct {
}

func (sm *serviceManager) UserUseCase() IUserUseCase {
	return NewUserUseCase()
}
func NewServiceManger() ServiceManager {
	return &serviceManager{}
}
