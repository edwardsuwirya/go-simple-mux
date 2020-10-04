package manager

import "gosimplemux/repositories"

type RepoManager interface {
	UserRepo() repositories.UserRepository
}
type repoManager struct {
}

func (rm *repoManager) UserRepo() repositories.UserRepository {
	return repositories.NewUserRepository()
}
func NewRepoManager() RepoManager {
	return &repoManager{}
}
