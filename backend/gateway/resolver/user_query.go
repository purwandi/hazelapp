package resolver

import (
	"context"

	"github.com/purwandi/hazelapp/user/domain"
)

func (r *Resolver) Viewer(ctx context.Context) (*UserResolver, error) {
	ViewerUser := ctx.Value("viewer")

	return &UserResolver{
		Field:    ViewerUser.(domain.User),
		Resolver: r,
	}, nil
}
