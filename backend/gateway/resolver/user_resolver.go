package resolver

import (
	"github.com/graph-gophers/graphql-go"
	"github.com/purwandi/hazelapp/helpers"
	"github.com/purwandi/hazelapp/user/domain"
)

type UserResolver struct {
	Field    domain.User
	Resolver *Resolver
}

func (u *UserResolver) ID() graphql.ID {
	return graphql.ID(helpers.EncryptID(string(u.Field.ID)))
}

func (u *UserResolver) Fullname() string {
	return u.Field.Fullname
}

func (u *UserResolver) Email() string {
	return u.Field.Email
}

func (u *UserResolver) Username() string {
	return u.Field.Username
}
