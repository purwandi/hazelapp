package services

import (
	"github.com/purwandi/hazelapp/helpers"
	"github.com/purwandi/hazelapp/user/domain"
	"github.com/purwandi/hazelapp/user/repository"
	"github.com/purwandi/hazelapp/user/types"
	"golang.org/x/crypto/bcrypt"
)

// UserService is a service to expose user service
type UserService struct {
	query repository.UserQuery
	repo  repository.UserRepository
}

// NewUserService to create new UserService instance
func NewUserService(q repository.UserQuery, r repository.UserRepository) *UserService {
	return &UserService{query: q, repo: r}
}

// FindUserByID is to find user by id
func (s *UserService) FindUserByID(id int) (domain.User, error) {
	result := <-s.query.FindUser(&types.FindUserInput{ID: helpers.Int(id)})
	if result.Error != nil {
		return domain.User{}, result.Error
	}

	user := result.Result.(domain.User)
	if user.ID == 0 {
		return domain.User{}, UserServiceError{UserServiceErrorResourceNotFound}
	}

	return user, nil
}

// FindUserByUsername is to find user by username
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

// FindUserByEmail is to find a user by email
func (s *UserService) FindUserByEmail(email string) (domain.User, error) {
	result := <-s.query.FindUser(&types.FindUserInput{Email: helpers.String(email)})
	if result.Error != nil {
		return domain.User{}, result.Error
	}

	user := result.Result.(domain.User)
	if user.ID == 0 {
		return domain.User{}, UserServiceError{UserServiceErrorResourceNotFound}
	}

	return user, nil
}

// FindUserByAccessToken is to find a user by access token
func (s *UserService) FindUserByAccessToken(token string) (domain.User, error) {
	result := <-s.query.FindUser(&types.FindUserInput{AccessToken: helpers.String(token)})
	if result.Error != nil {
		return domain.User{}, result.Error
	}

	user := result.Result.(domain.User)
	if user.ID == 0 {
		return domain.User{}, UserServiceError{UserServiceErrorResourceNotFound}
	}

	return user, nil
}

// RegisterUser to create new user
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

// LoginUser is to validate login user
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
