package resolver

import (
	"context"
	"time"

	"github.com/graph-gophers/graphql-go"
	"github.com/purwandi/hazelapp/helpers"
	"github.com/purwandi/hazelapp/project/domain"
	"github.com/purwandi/hazelapp/project/repository"
)

// ProjectResolver is a struct to resolve project domain
type ProjectResolver struct {
	Field    domain.Project
	Resolver *Resolver
}

// ID get project ID
func (p *ProjectResolver) ID() graphql.ID {
	return graphql.ID(helpers.EncryptID("project", p.Field.ID))
}

// Owner get project owner
func (p *ProjectResolver) Owner(ctx context.Context) (*UserResolver, error) {
	return p.Resolver.UserByID(ctx, p.Field.OwnerID)
}

// Name get project name
func (p *ProjectResolver) Name() string {
	return p.Field.Name
}

// Description get project description
func (p *ProjectResolver) Description() *string {
	return helpers.String(p.Field.Description)
}

// CreatedAt is to get project creation date
func (p *ProjectResolver) CreatedAt() string {
	return p.Field.CreatedAt.Format(time.RFC3339)
}

// ProjectConnectionResolver is to get project connection
type ProjectConnectionResolver struct {
	Projects  []*ProjectResolver
	Paginator repository.PageInfo
	Count     int
}

// Edges to get project edges
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

// Nodes to get available nodes
func (r *ProjectConnectionResolver) Nodes() *[]*ProjectResolver {
	return &r.Projects
}

// TotalCount to get project total
func (r *ProjectConnectionResolver) TotalCount() int32 {
	return int32(r.Count)
}

// PageInfo is resolves ...
func (r *ProjectConnectionResolver) PageInfo() PageInfoResolver {
	return PageInfoResolver{
		startCursor: helpers.String(helpers.EncryptID("project", r.Paginator.StartCursor)),
		endCursor:   helpers.String(helpers.EncryptID("project", r.Paginator.EndCursor)),
	}
}

// ProjectEdgeResolver is to wrap project edges
type ProjectEdgeResolver struct {
	cursor graphql.ID
	Field  *ProjectResolver
}

// Cursor to get project pagination cursor
func (r *ProjectEdgeResolver) Cursor() graphql.ID {
	return graphql.ID(r.Field.ID())
}

// Node get project pagination node
func (r *ProjectEdgeResolver) Node() *ProjectResolver {
	return r.Field
}
