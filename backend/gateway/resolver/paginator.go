package resolver

// PageInfoResolver is information about pagination in a connection
type PageInfoResolver struct {
	startCursor     *string
	endCursor       *string
	hasNextPage     bool
	hasPreviousPage bool
}

// StartCursor resolves ...
func (p PageInfoResolver) StartCursor() *string {
	return p.startCursor
}

// EndCursor resolves ...
func (p PageInfoResolver) EndCursor() *string {
	return p.endCursor
}

// HasNextPage resolves ...
func (p PageInfoResolver) HasNextPage() bool {
	return p.hasNextPage
}

// HasPreviousPage resolves ...
func (p PageInfoResolver) HasPreviousPage() bool {
	return p.hasPreviousPage
}
