package resolver

import (
	"time"

	"github.com/graph-gophers/graphql-go"
	"github.com/purwandi/hazelapp/issue/domain"
)

type IssueResolver struct {
	Field    domain.Issue
	Resolver *Resolver
}

func (i *IssueResolver) ID() graphql.ID {
	return graphql.ID(i.Field.ID.String())
}

func (i *IssueResolver) IID() int32 {
	return int32(i.Field.AutoNumber)
}

func (i *IssueResolver) Title() string {
	return i.Field.Title
}

func (i *IssueResolver) Body() string {
	return i.Field.Body
}

func (i *IssueResolver) CreatedAt() string {
	return i.Field.CreatedAt.Format(time.RFC3339)
}
