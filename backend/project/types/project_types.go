package types

type CreateProjectInput struct {
	OwnerID     int
	Name        string
	Code        *string
	Description *string
}

type ProjectOrder struct {
	Field     string
	Direction string
}

type GetProjectsInput struct {
	OwnerID int
	After   *int
	Before  *int
	First   *int
	Last    *int
	OrderBy *ProjectOrder
}

type FindProjectInput struct {
	OwnerID int
	Name    string
}
