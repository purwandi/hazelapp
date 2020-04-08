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

const (
	// StoryType is story
	StoryType = "story"
	// BugType is bug
	BugType = "bug"
	// TaskType is task
	TaskType = "task"
	// EpicType is epic
	EpicType = "epic"
)

// IssueType is issue type
type IssueType struct {
	Code  string
	Label string
}

// issueTypes is to get available issue types
func issueTypes() []IssueType {
	return []IssueType{
		IssueType{Code: StoryType, Label: "Story"},
		IssueType{Code: BugType, Label: "Bug"},
		IssueType{Code: TaskType, Label: "Task"},
		IssueType{Code: EpicType, Label: "Epic"},
	}
}

func validateIssueType(issueType string) error {
	for _, item := range issueTypes() {
		if item.Code == issueType {
			return nil
		}
	}
	return IssueError{IssueErrorUndefinedIssueType}
}

// Issue is to wrap issue domain
type Issue struct {
	ID          int
	ProjectID   int
	AuthorID    int
	Number      int
	MilestoneID *int
	IssueType   string
	Title       string
	Body        *string
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

	if err := validateIssueType(args.IssueType); err != nil {
		return nil, err
	}

	iid, err := service.GetLastIssueNumberFromGivenProject(args.ProjectID)
	if err != nil {
		return nil, err
	}

	issue := &Issue{
		ProjectID: args.ProjectID,
		AuthorID:  args.AuthorID,
		Number:    iid + 1,
		Title:     args.Title,
		Body:      args.Body,
		State:     IssueOpen,
		IssueType: args.IssueType,
		CreatedAt: time.Now(),
	}

	if args.MilestoneID != nil {
		issue.MilestoneID = args.MilestoneID
	}

	return issue, nil
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
	i.Body = &body
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

// ChangeMilestone is to change milestone
func (i *Issue) ChangeMilestone(milestone int) error {
	i.MilestoneID = &milestone
	i.UpdatedAt = time.Now()
	return nil
}
