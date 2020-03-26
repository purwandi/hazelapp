package inmemory

import (
	"sort"

	"github.com/purwandi/hazelapp/issue/domain"
)

type by func(p1, p2 *domain.Issue) bool

func (b by) Sort(issues []domain.Issue) {
	s := &issueSorter{
		issues: issues,
		by:     b,
	}

	sort.Sort(s)
}

type issueSorter struct {
	issues []domain.Issue
	by     func(d1, d2 *domain.Issue) bool
}

// Len is part of sort.Interface.
func (d *issueSorter) Len() int {
	return len(d.issues)
}

// Swap is part of sort.Interface.
func (d *issueSorter) Swap(i, j int) {
	d.issues[i], d.issues[j] = d.issues[j], d.issues[i]
}

// Less is part of sort.Interface. It is implemented by calling the "by" closure in the sorter.
func (d *issueSorter) Less(i, j int) bool {
	return d.by(&d.issues[i], &d.issues[j])
}
