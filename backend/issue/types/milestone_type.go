package types

type GetMilestoneQueryInput struct {
	ProjectID string
}

type CreateMilestoneInput struct {
	ProjectID   string
	Name        string
	Description *string
}
