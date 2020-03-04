package storage

import "github.com/purwandi/hazelapp/user/domain"

type UserStorage struct {
	UserMap []domain.User
}

func NewUserStorage() *UserStorage {
	return &UserStorage{UserMap: []domain.User{}}
}

func (store *UserStorage) Demo() error {
	store.UserMap = []domain.User{
		domain.User{ID: 1, Email: "foobar@kaboor.com", Username: "foobar"},
		domain.User{ID: 2, Email: "baban@kaboo.com", Username: "baban"},
		domain.User{ID: 3, Email: "kazom@kaboo.com", Username: "kazoom"},
	}

	return nil
}
