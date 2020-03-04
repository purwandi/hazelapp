package inmemory

import (
	"github.com/purwandi/hazelapp/user/domain"
	"github.com/purwandi/hazelapp/user/repository"
	"github.com/purwandi/hazelapp/user/storage"
	"github.com/purwandi/hazelapp/user/types"
)

type UserQueryInMemory struct {
	Storage *storage.UserStorage
}

func NewUserQueryInMemory(s *storage.UserStorage) repository.UserQuery {
	return &UserQueryInMemory{Storage: s}
}

func (query *UserQueryInMemory) FindUser(args *types.FindUserInput) <-chan repository.QueryResult {
	result := make(chan repository.QueryResult)

	go func() {
		user := domain.User{}
		for _, item := range query.Storage.UserMap {
			if args.ID != nil && *args.ID == item.ID {
				user = item
				break
			}

			if args.Email != nil && *args.Email == item.Email {
				user = item
				break
			}

			if args.Username != nil && *args.Username == item.Username {
				user = item
				break
			}
		}

		result <- repository.QueryResult{Result: user}
		close(result)
	}()

	return result
}
