package resolver

import (
	"context"

	"github.com/purwandi/hazelapp/helpers"
	"github.com/purwandi/hazelapp/user/domain"
)

// UserLoginInput is graphql user input
type UserLoginInput struct {
	Login string
}

// Viewer is to get current viewer
func (r *Resolver) Viewer(ctx context.Context) (*UserResolver, error) {
	ViewerUser := ctx.Value("viewer")

	return &UserResolver{
		Field:    ViewerUser.(domain.User),
		Resolver: r,
	}, nil
}

// User is to get user by login
func (r *Resolver) User(ctx context.Context, args UserLoginInput) (*UserResolver, error) {
	userID, err := helpers.DecryptID("user-", args.Login)
	if err != nil {
		return nil, err
	}

	user, err := r.UserService.FindUserByID(userID)
	if err != nil {
		return nil, err
	}

	return &UserResolver{
		Field:    user,
		Resolver: r,
	}, nil
}
