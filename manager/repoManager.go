package manager

import "gosimplemux/repositories"

type RepoManager interface {
	UserRepo() repositories.UserRepository
	UserAuthRepo() repositories.UserAuthRepository
}
type repoManager struct {
}

func (rm *repoManager) UserRepo() repositories.UserRepository {
	return repositories.NewUserRepository()
}
func (rm *repoManager) UserAuthRepo() repositories.UserAuthRepository {
	return repositories.NewUserAuthRepository()
}

func NewRepoManager() RepoManager {
	return &repoManager{}
}
