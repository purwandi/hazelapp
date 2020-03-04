package inmemory

import (
	"github.com/purwandi/hazelapp/user/domain"
	"github.com/purwandi/hazelapp/user/repository"
	"github.com/purwandi/hazelapp/user/storage"
)

type UserRepositoryInMemory struct {
	Storage *storage.UserStorage
}

func NewUserRepositoryInMemory(s *storage.UserStorage) repository.UserRepository {
	return &UserRepositoryInMemory{Storage: s}
}

func (repo *UserRepositoryInMemory) Save(user *domain.User) <-chan error {
	result := make(chan error)

	go func() {
		lastUserID := 1
		if len(repo.Storage.UserMap) > 0 {
			lastUserID = repo.Storage.UserMap[len(repo.Storage.UserMap)-1].ID + 1
		}

		user.ID = lastUserID
		repo.Storage.UserMap = append(repo.Storage.UserMap, *user)

		result <- nil
		close(result)
	}()

	return result
}
