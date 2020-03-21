package resolver

import (
	"context"

	"github.com/graph-gophers/graphql-go"
	"github.com/purwandi/hazelapp/helpers"
	"github.com/purwandi/hazelapp/user/domain"
)

// UserResolver is to resolve user resolver
type UserResolver struct {
	Field    domain.User
	Resolver *Resolver
}

// ID is to get user id
func (u *UserResolver) ID() graphql.ID {
	return graphql.ID(helpers.EncryptID("user", u.Field.ID))
}

// Fullname is to get user fullname
func (u *UserResolver) Fullname() string {
	return u.Field.Fullname
}

// Email is to get user email
func (u *UserResolver) Email() string {
	return u.Field.Email
}

// Username is to get username
func (u *UserResolver) Username() string {
	return u.Field.Username
}

// Projects it to get project belongs to
func (u *UserResolver) Projects(ctx context.Context, args struct{ Input *GetProjectsInput }) (*ProjectConnectionResolver, error) {
	return u.Resolver.Projects(ctx, u.Field.ID, args)
}
