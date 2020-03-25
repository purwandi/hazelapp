package domain

import (
	"time"

	"github.com/purwandi/hazelapp/issue/types"
)

const (
	// IssueOpen is when issue state is open
	IssueOpen = "OPEN"
	// IssueClosed is when issue state is closed
	IssueClosed = "CLOSED"
)

// Issue is to wrap issue domain
type Issue struct {
	ID          int
	ProjectID   int
	AuthorID    int
	Number      int
	MilestoneID int
	Title       string
	Body        string
	State       string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// IssueService is an interface to get last issue number from given project
type IssueService interface {
	GetLastIssueNumberFromGivenProject(int) (int, error)
}

// CreateIssue is to create new issue instance
func CreateIssue(service IssueService, args types.CreateIssueInput) (*Issue, error) {
	if args.Title == "" {
		return nil, IssueError{IssueErrorTitleIsBlank}
	}

	iid, err := service.GetLastIssueNumberFromGivenProject(args.ProjectID)
	if err != nil {
		return nil, err
	}

	return &Issue{
		ProjectID:   args.ProjectID,
		AuthorID:    args.AuthorID,
		MilestoneID: args.MilestoneID,
		Number:      iid + 1,
		Title:       args.Title,
		Body:        args.Body,
		State:       IssueOpen,
		CreatedAt:   time.Now(),
	}, nil
}

// ChangeTitle to change issue title
func (i *Issue) ChangeTitle(title string) error {
	if title == "" {
		return IssueError{IssueErrorTitleIsBlank}
	}

	i.Title = title
	i.UpdatedAt = time.Now()
	return nil
}

// ChangeBody is to change issue body
func (i *Issue) ChangeBody(body string) error {
	i.Body = body
	i.UpdatedAt = time.Now()
	return nil
}

// ReOpen is to re open issue
func (i *Issue) ReOpen() error {
	i.State = IssueOpen
	i.UpdatedAt = time.Now()
	return nil
}

// Closed is to close issue
func (i *Issue) Closed() error {
	i.State = IssueClosed
	i.UpdatedAt = time.Now()
	return nil
}
