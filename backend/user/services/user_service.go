package services

import (
	"github.com/purwandi/hazelapp/helpers"
	"github.com/purwandi/hazelapp/user/domain"
	"github.com/purwandi/hazelapp/user/repository"
	"github.com/purwandi/hazelapp/user/types"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	query repository.UserQuery
	repo  repository.UserRepository
}

func NewUserService(q repository.UserQuery, r repository.UserRepository) *UserService {
	return &UserService{query: q, repo: r}
}

func (s *UserService) FindUserByUsername(username string) (domain.User, error) {
	result := <-s.query.FindUser(&types.FindUserInput{Username: helpers.String(username)})
	if result.Error != nil {
		return domain.User{}, result.Error
	}

	user := result.Result.(domain.User)
	if user.ID == 0 {
		return domain.User{}, UserServiceError{UserServiceErrorResourceNotFound}
	}

	return user, nil
}

func (s *UserService) FindUserByEmail(email string) (domain.User, error) {
	result := <-s.query.FindUser(&types.FindUserInput{Email: helpers.String(email)})
	logrus.Info(result.Result.(domain.User))
	if result.Error != nil {
		return domain.User{}, result.Error
	}

	user := result.Result.(domain.User)
	if user.ID == 0 {
		return domain.User{}, UserServiceError{UserServiceErrorResourceNotFound}
	}

	return user, nil
}

func (s *UserService) RegisterUser(args types.RegisterUserInput) (domain.User, error) {
	user, err := domain.RegisterUser(s, args)
	if err != nil {
		return domain.User{}, err
	}

	// Persist
	if err := <-s.repo.Save(user); err != nil {
		return domain.User{}, err
	}

	// Return
	return *user, nil
}

func (s *UserService) LoginUser(args types.LoginUserInput) (domain.User, error) {
	result := <-s.query.FindUser(&types.FindUserInput{Username: helpers.String(args.Username)})
	if result.Error != nil {
		return domain.User{}, result.Error
	}

	user := result.Result.(domain.User)

	// Checking password
	err := bcrypt.CompareHashAndPassword(user.Password, []byte(args.Password))
	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}
