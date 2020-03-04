package domain

import (
	"time"

	uuid "github.com/satori/go.uuid"

	"github.com/purwandi/hazelapp/user/types"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	FindUserByEmail(string) (User, error)
	FindUserByUsername(string) (User, error)
}

// User is user domain
type User struct {
	ID          int       `json:"id"`
	Fullname    string    `json:"fullname"`
	Username    string    `json:"username"`
	Email       string    `json:"email"`
	Password    []byte    `json:"-"`
	AccessToken string    `json:"accessToken"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

// RegisterUser it to handle register user creation
func RegisterUser(service UserService, args types.RegisterUserInput) (*User, error) {
	if args.Fullname == "" {
		return nil, UserError{UserErrorFullNameIsBlank}
	}

	if args.Email == "" {
		return nil, UserError{UserErrorEmailIsBlank}
	}

	if args.Username == "" {
		return nil, UserError{UserErrorUsernameIsBlank}
	}

	if args.Password == "" {
		return nil, UserError{UserErrorPasswordIsBlank}
	}

	if args.PasswordConfirmation != args.Password {
		return nil, UserError{UserErrorPasswordConfirmationIsNotMatch}
	}

	if _, err := service.FindUserByEmail(args.Email); err == nil {
		return nil, UserError{UserErrorEmailIsExists}
	}

	if _, err := service.FindUserByUsername(args.Username); err == nil {
		return nil, UserError{UserErrorUsernameIsExists}
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(args.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	token, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}

	return &User{
		Fullname:    args.Fullname,
		Username:    args.Username,
		Email:       args.Email,
		Password:    hash,
		AccessToken: token.String(),
		CreatedAt:   time.Now(),
	}, nil
}
