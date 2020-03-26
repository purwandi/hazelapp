package inmemory

import (
	"github.com/purwandi/hazelapp/issue/domain"
	"github.com/purwandi/hazelapp/issue/repository"
	"github.com/purwandi/hazelapp/issue/storage"
	"github.com/purwandi/hazelapp/issue/types"
)

// IssueQueryInMemory is to IssueQuery implementation in memory
type IssueQueryInMemory struct {
	Storage *storage.IssueStorage
}

// NewIssueQueryInMemory is to create IssueQueryInMemory instance
func NewIssueQueryInMemory(s *storage.IssueStorage) repository.IssueQuery {
	return &IssueQueryInMemory{Storage: s}
}

// GetLastIssueNumberFromGivenProject is to get last issue number from given project
func (query *IssueQueryInMemory) GetLastIssueNumberFromGivenProject(id int) <-chan repository.QueryResult {
	result := make(chan repository.QueryResult)

	go func() {
		issue := domain.Issue{}
		for _, item := range query.Storage.IssueMap {
			if item.ProjectID == id && item.Number > issue.Number {
				issue = item
			}
		}

		if (domain.Issue{}) == issue {
			result <- repository.QueryResult{Result: 0}
		}

		result <- repository.QueryResult{Result: issue.Number}
		close(result)
	}()

	return result
}

// GetIssues to get all issues
func (query *IssueQueryInMemory) GetIssues(args types.GetIssuesInput) <-chan repository.QueryResult {
	result := make(chan repository.QueryResult)
	total := 0

	go func() {
		issues := []domain.Issue{}

		if args.First == nil && args.Last == nil {
			result <- repository.QueryResult{
				Error: repository.IssueQueryError{Code: repository.IssueQueryErrorMissingPaginationBoundaries},
			}
		}

		if args.First != nil && args.Last != nil {
			result <- repository.QueryResult{
				Error: repository.IssueQueryError{Code: repository.IssueQueryErrorPagination},
			}
		}

		for _, issue := range query.Storage.IssueMap {
			if issue.ProjectID == args.ProjectID {
				issues = append(issues, issue)
			}
		}

		// Filter

		// Sort
		by(func(p1, p2 *domain.Issue) bool {
			return p1.ID < p2.ID
		}).Sort(issues)

		// Calculate total query
		total = len(issues)
		if total == 0 {
			result <- repository.QueryResult{
				Result: issues,
			}
		}

		// Cursor
		if args.After != nil {
			var offset int
			for index, item := range issues {
				if item.ID > *args.After {
					offset = index
					break
				}
			}
			issues = issues[offset:]
		}

		if args.Before != nil {
			var offset int
			for index, item := range issues {
				if item.ID == *args.Before {
					offset = index
				}
			}
			issues = issues[:offset]
		}

		// Get n first items
		if args.First != nil {
			limit := *args.First
			if limit > len(issues) {
				limit = len(issues)
			}
			issues = issues[:limit]
		}

		if args.Last != nil {
			limit := *args.Last
			if limit > len(issues) {
				limit = len(issues)
			}
			offset := len(issues) - limit
			issues = issues[offset:]
		}

		// Result
		res := repository.QueryResult{
			Result: issues,
			Total:  total,
		}

		if len(issues) > 0 {
			res.PageInfo = repository.PageInfo{
				StartCursor: issues[0].ID,
				EndCursor:   issues[len(issues)-1].ID,
				// Todo implement hasNextPage and endNextPage
			}
		}

		result <- res
		close(result)
	}()

	return result
}

// GetIssueByNumber to get an issue
func (query *IssueQueryInMemory) GetIssueByNumber(args types.GetIssueInput) <-chan repository.QueryResult {
	result := make(chan repository.QueryResult)

	go func() {
		issue := domain.Issue{}
		for _, item := range query.Storage.IssueMap {
			if item.Number == args.Number && item.ProjectID == args.ProjectID {
				issue = item
			}
		}

		result <- repository.QueryResult{Result: issue}
		close(result)
	}()

	return result
}
