package resolver

import (
	"github.com/graph-gophers/graphql-go"
	"github.com/purwandi/hazelapp/helpers"
	"github.com/purwandi/hazelapp/issue/domain"
)

// MilestoneResolver resolve ...
type MilestoneResolver struct {
	Field    domain.Milestone
	Resolver *Resolver
}

// ID resolve ...
func (m *MilestoneResolver) ID() graphql.ID {
	return graphql.ID(helpers.EncryptID("milestone", m.Field.ID))
}

// Project resolve ...
func (m *MilestoneResolver) Project() (*ProjectResolver, error) {
	project, err := m.Resolver.ProjectService.FindProjectByID(m.Field.ProjectID)
	if err != nil {
		return nil, err
	}

	return &ProjectResolver{
		Field:    project,
		Resolver: m.Resolver,
	}, nil
}

// Name resolve ...
func (m *MilestoneResolver) Name() string {
	return m.Field.Name
}

// Description resolve ...
func (m *MilestoneResolver) Description() *string {
	return helpers.String(m.Field.Description)
}

// StartDate resolve ...
func (m *MilestoneResolver) StartDate() *string {
	return helpers.String("")
}

// EndDate resolve ...
func (m *MilestoneResolver) EndDate() *string {
	return helpers.String("")
}
