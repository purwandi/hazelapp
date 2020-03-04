package repository

import (
	"github.com/purwandi/hazelapp/user/domain"
	"github.com/purwandi/hazelapp/user/types"
)

type QueryResult struct {
	Result interface{}
	Error  error
}

type UserQuery interface {
	FindUser(types *types.FindUserInput) <-chan QueryResult
}

type UserRepository interface {
	Save(*domain.User) <-chan error
}
