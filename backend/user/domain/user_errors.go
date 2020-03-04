package domain

type UserError struct {
	Code int
}

const (
	UserErrorFullNameIsBlank = iota
	UserErrorEmailIsBlank
	UserErrorEmailIsExists
	UserErrorUsernameIsBlank
	UserErrorUsernameIsExists
	UserErrorPasswordIsBlank
	UserErrorPasswordConfirmationIsNotMatch
)

func (e UserError) Error() string {
	switch e.Code {
	default:
		return "Unrecognized user error code"
	case UserErrorFullNameIsBlank:
		return "Fullname can't be blank"
	case UserErrorEmailIsBlank:
		return "Email can't be blank"
	case UserErrorEmailIsExists:
		return "Email already exist"
	case UserErrorUsernameIsBlank:
		return "Username can't be blank"
	case UserErrorUsernameIsExists:
		return "Username already exists"
	case UserErrorPasswordIsBlank:
		return "Password can't be blank"
	case UserErrorPasswordConfirmationIsNotMatch:
		return "Password confirmation is not match"
	}
}
