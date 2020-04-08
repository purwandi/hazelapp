package storage

import "github.com/purwandi/hazelapp/issue/domain"

// IssueStorage is an implemantion issue in memory storage data
type IssueStorage struct {
	IssueMap []domain.Issue
}

// NewIssueStorage to create IssueStorage instance
func NewIssueStorage() *IssueStorage {
	return &IssueStorage{
		IssueMap: []domain.Issue{},
	}
}

// Demo to feed data demo for IssueStorage
func (i *IssueStorage) Demo() error {
	i.IssueMap = []domain.Issue{
		domain.Issue{ID: 1, Number: 1, ProjectID: 1, AuthorID: 1, State: domain.IssueOpen, IssueType: "bug", Title: "Bug: React Lazy Suspense with Errorboundary fails on older devices"},
		domain.Issue{ID: 2, Number: 2, ProjectID: 1, AuthorID: 2, State: domain.IssueClosed, IssueType: "story", Title: "Bug: High-pri setState causes primary tree to get unhidden"},
		domain.Issue{ID: 3, Number: 3, ProjectID: 1, AuthorID: 3, State: domain.IssueOpen, IssueType: "story", Title: "Bug: Updates in the primary tree only unsuspend once"},
		domain.Issue{ID: 4, Number: 4, ProjectID: 1, AuthorID: 4, State: domain.IssueClosed, IssueType: "bug", Title: "Bug: React-test-renderer error when updating state in act "},
		domain.Issue{ID: 5, Number: 5, ProjectID: 1, AuthorID: 5, State: domain.IssueOpen, IssueType: "story", Title: "Bug: react-refresh silently fails if version does not match react version"},
		domain.Issue{ID: 6, Number: 6, ProjectID: 1, AuthorID: 1, State: domain.IssueClosed, IssueType: "task", Title: "Custom file picker validation message overlaps any following content"},
		domain.Issue{ID: 7, Number: 7, ProjectID: 1, AuthorID: 2, State: domain.IssueOpen, IssueType: "story", Title: "Using calc() for columns width"},
		domain.Issue{ID: 8, Number: 8, ProjectID: 1, AuthorID: 3, State: domain.IssueClosed, IssueType: "task", Title: "Create mixin for anchor tag (link)"},
		domain.Issue{ID: 9, Number: 9, ProjectID: 1, AuthorID: 5, State: domain.IssueClosed, IssueType: "story", Title: "Segoe UI clipping with overflow:hidden"},
		domain.Issue{ID: 10, Number: 10, ProjectID: 1, AuthorID: 2, State: domain.IssueOpen, IssueType: "story", Title: "Split variable $table-cell-padding into separate variables"},
		domain.Issue{ID: 11, Number: 11, ProjectID: 1, AuthorID: 6, State: domain.IssueClosed, IssueType: "task", Title: "HTML style reboot issue for placeholder link"},
		domain.Issue{ID: 12, Number: 12, ProjectID: 1, AuthorID: 9, State: domain.IssueOpen, IssueType: "story", Title: "tabs - accessibility issue when using ul/li semantic"},
		domain.Issue{ID: 13, Number: 13, ProjectID: 1, AuthorID: 1, State: domain.IssueClosed, IssueType: "bug", Title: "Click on Static Modal Backdrop Scaling up the Modal Content and Adding Scrollbar to Backdrop"},
		domain.Issue{ID: 14, Number: 14, ProjectID: 1, AuthorID: 2, State: domain.IssueOpen, IssueType: "task", Title: "Update image previews on the Bootstrap examples page"},
		domain.Issue{ID: 15, Number: 15, ProjectID: 1, AuthorID: 3, State: domain.IssueClosed, IssueType: "story", Title: "Carousel - Back To Top anchor link does not always work."},
		domain.Issue{ID: 16, Number: 16, ProjectID: 1, AuthorID: 4, State: domain.IssueClosed, IssueType: "story", Title: "Sass: Add Option like $enable-grid-classes for All Utility Classes"},
		domain.Issue{ID: 17, Number: 17, ProjectID: 1, AuthorID: 5, State: domain.IssueOpen, IssueType: "story", Title: "Dropdown data-toggle dropdown element click events not bubble"},

		domain.Issue{ID: 18, Number: 1, ProjectID: 2, AuthorID: 1, State: domain.IssueOpen, IssueType: "story", Title: "Use `<article>`s for cards when appropriate "},
		domain.Issue{ID: 19, Number: 2, ProjectID: 2, AuthorID: 2, State: domain.IssueOpen, IssueType: "story", Title: "`bs-custom-file-input` doesn't work anymore since the form updates "},
		domain.Issue{ID: 20, Number: 3, ProjectID: 2, AuthorID: 3, State: domain.IssueClosed, IssueType: "task", Title: "Add support for block buttons based on specific screen widths, eg. btn-xs-block"},
		domain.Issue{ID: 21, Number: 4, ProjectID: 2, AuthorID: 4, State: domain.IssueOpen, IssueType: "story", Title: "Double increase carousel interval after click 4.4.1"},
		domain.Issue{ID: 22, Number: 5, ProjectID: 2, AuthorID: 5, State: domain.IssueOpen, IssueType: "task", Title: "Revisit browser and device docs"},
		domain.Issue{ID: 23, Number: 6, ProjectID: 2, AuthorID: 2, State: domain.IssueClosed, IssueType: "story", Title: "Explore a .content class for broad margin changes and more"},
		domain.Issue{ID: 24, Number: 7, ProjectID: 2, AuthorID: 2, State: domain.IssueOpen, IssueType: "task", Title: "input-group-lg and input-group-sm override validation icon padding"},
		domain.Issue{ID: 25, Number: 8, ProjectID: 2, AuthorID: 1, State: domain.IssueOpen, IssueType: "story", Title: "dropdown-dark"},

		domain.Issue{ID: 26, Number: 1, ProjectID: 3, AuthorID: 1, State: domain.IssueOpen, IssueType: "story", Title: "Downloadable Bootstrap Examples"},
		domain.Issue{ID: 27, Number: 2, ProjectID: 3, AuthorID: 2, State: domain.IssueClosed, IssueType: "story", Title: "<blockquote> tag being specified incorrectly"},
		domain.Issue{ID: 28, Number: 3, ProjectID: 3, AuthorID: 3, State: domain.IssueOpen, IssueType: "bug", Title: "Using SVGs inline"},
		domain.Issue{ID: 29, Number: 4, ProjectID: 3, AuthorID: 2, State: domain.IssueOpen, IssueType: "bug", Title: "Create standalone accordion component"},
		domain.Issue{ID: 30, Number: 5, ProjectID: 3, AuthorID: 2, State: domain.IssueOpen, IssueType: "bug", Title: "Confusing classes `rounded-sm` and `rounded-lg`"},
	}

	return nil
}
