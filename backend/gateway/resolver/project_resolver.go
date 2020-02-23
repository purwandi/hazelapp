package resolver

import (
	"time"

	"github.com/graph-gophers/graphql-go"
	"github.com/purwandi/hazelapp/project/domain"
	"github.com/purwandi/hazelapp/types"
)

// ProjectResolver is a struct to resolve project domain
type ProjectResolver struct {
	Field    domain.Project
	Resolver *Resolver
}

func (p *ProjectResolver) ID() graphql.ID {
	return graphql.ID(p.Field.ID.String())
}

func (p *ProjectResolver) Name() string {
	return p.Field.Name
}

func (p *ProjectResolver) Description() *string {
	return types.String(p.Field.Description)
}

func (p *ProjectResolver) CreatedAt() string {
	return p.Field.CreatedAt.Format(time.RFC3339)
}

type ProjectEdgeResolver struct {
	cursor graphql.ID
	Field  *ProjectResolver
}

func (r *ProjectEdgeResolver) Cursor() graphql.ID {
	return graphql.ID(r.Field.ID())
}

func (r *ProjectEdgeResolver) Node() *ProjectResolver {
	return r.Field
}

type ProjectConnectionResolver struct {
	Projects []*ProjectResolver
	Count    int
	From     string
	To       string
}

func (r *ProjectConnectionResolver) TotalCount() int32 {
	return int32(r.Count)
}

func (r *ProjectConnectionResolver) Edges() *[]*ProjectEdgeResolver {
	projects := make([]*ProjectEdgeResolver, len(r.Projects))

	for i, project := range r.Projects {
		projects[i] = &ProjectEdgeResolver{
			cursor: project.ID(),
			Field:  project,
		}
	}

	return &projects
}
