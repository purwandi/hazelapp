package resolver

import (
	"github.com/graph-gophers/graphql-go"
	"github.com/purwandi/hazelapp/issue/domain"
	"github.com/purwandi/hazelapp/types"
)

type MilestoneResolver struct {
	Field    domain.Milestone
	Resolver *Resolver
}

func (m *MilestoneResolver) ID() graphql.ID {
	return graphql.ID(m.Field.ID.String())
}

func (m *MilestoneResolver) Name() string {
	return m.Field.Name
}

func (m *MilestoneResolver) Description() *string {
	return types.String(m.Field.Description)
}

func (m *MilestoneResolver) ProjectId() string {
	return m.Field.ProjectID.String()
}
