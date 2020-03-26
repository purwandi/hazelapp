package inmemory

import (
	"sort"

	"github.com/purwandi/hazelapp/project/domain"
)

type by func(p1, p2 *domain.Project) bool

func (b by) Sort(projects []domain.Project) {
	s := &projectSorter{
		projects: projects,
		by:       b,
	}

	sort.Sort(s)
}

type projectSorter struct {
	projects []domain.Project
	by       func(p1, p2 *domain.Project) bool
}

// Len is part of sort.Interface.
func (p *projectSorter) Len() int {
	return len(p.projects)
}

// Swap is part of sort.Interface.
func (p *projectSorter) Swap(i, j int) {
	p.projects[i], p.projects[j] = p.projects[j], p.projects[i]
}

// Less is part of sort.Interface. It is implemented by calling the "by" closure in the sorter.
func (p *projectSorter) Less(i, j int) bool {
	return p.by(&p.projects[i], &p.projects[j])
}
