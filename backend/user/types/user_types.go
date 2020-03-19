package types

type RegisterUserInput struct {
	Fullname             string
	Email                string
	Username             string
	Password             string
	PasswordConfirmation string
}

type LoginUserInput struct {
	Username string
	Password string
}

type FindUserInput struct {
	ID          *int
	Email       *string
	Username    *string
	AccessToken *string
}
