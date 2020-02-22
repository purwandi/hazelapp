package resolver

import "context"

type ProjectCreateInput struct {
	Name        string
	Description *string
}

func (r *Resolver) ProjectCreate(ctx context.Context, args ProjectCreateInput) (*ProjectResolver, error) {
	return &ProjectResolver{}, nil
}
