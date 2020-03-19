package storage

import (
	"github.com/purwandi/hazelapp/user/domain"
	"golang.org/x/crypto/bcrypt"
)

type UserStorage struct {
	UserMap []domain.User
}

func NewUserStorage() *UserStorage {
	return &UserStorage{UserMap: []domain.User{}}
}

func (store *UserStorage) Demo() error {
	hash, err := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	store.UserMap = []domain.User{
		domain.User{ID: 1, Email: "foobar@kaboor.com", Username: "foobar", Fullname: "Foobar Keba", Password: hash, AccessToken: "a8340b1d-1293-4efa-8da1-e2fa26dfa740"},
		domain.User{ID: 2, Email: "baban@kaboo.com", Username: "baban", Fullname: "Baban Indatu", Password: hash, AccessToken: "a8340b1d-1293-4efa-8da1-e2fa26dfa741"},
		domain.User{ID: 3, Email: "kazom@kaboo.com", Username: "kazoom", Fullname: "Kazoom Kaxe", Password: hash, AccessToken: "a8340b1d-1293-4efa-8da1-e2fa26dfa742"},
	}

	return nil
}
