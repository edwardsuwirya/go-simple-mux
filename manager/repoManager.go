package manager

import (
	"gosimplemux/infra"
	"gosimplemux/repositories"
)

type RepoManager interface {
	UserRepo() repositories.UserRepository
	UserAuthRepo() repositories.UserAuthRepository
}
type repoManager struct {
	infra infra.Infra
}

func (rm *repoManager) UserRepo() repositories.UserRepository {
	return repositories.NewUserRepository(rm.infra.SqlDb())
}
func (rm *repoManager) UserAuthRepo() repositories.UserAuthRepository {
	return repositories.NewUserAuthRepository(rm.infra.SqlDb())
}

func NewRepoManager(infra infra.Infra) RepoManager {
	return &repoManager{infra}
}
