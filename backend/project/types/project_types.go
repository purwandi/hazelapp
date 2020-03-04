package types

type CreateProjectInput struct {
	OwnerID     int
	Name        string
	Code        *string
	Description *string
}

type GetProjectsInput struct{}

type FindProjectInput struct {
	OwnerID int
	Name    string
}
