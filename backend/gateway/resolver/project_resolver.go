package resolver

import (
	"context"
	"time"

	"github.com/graph-gophers/graphql-go"
	"github.com/purwandi/hazelapp/helpers"
	"github.com/purwandi/hazelapp/project/domain"
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
func (p *ProjectResolver) Owner() (*UserResolver, error) {
	return p.Resolver.User(
		context.Background(),
		UserLoginInput{Login: helpers.EncryptID("user", p.Field.OwnerID)},
	)
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
	Projects []*ProjectResolver
	Count    int
	From     string
	To       string
}

// TotalCount to get project total
func (r *ProjectConnectionResolver) TotalCount() int32 {
	return int32(r.Count)
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
