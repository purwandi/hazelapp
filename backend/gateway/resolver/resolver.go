package resolver

import (
	"os"

	RepositoryInMemoryIssue "github.com/purwandi/hazelapp/issue/repository/inmemory"
	ServiceIssue "github.com/purwandi/hazelapp/issue/services"
	StorageIssue "github.com/purwandi/hazelapp/issue/storage"
	RepositoryInMemoryProject "github.com/purwandi/hazelapp/project/repository/inmemory"
	ServiceProject "github.com/purwandi/hazelapp/project/services"
	StorageProject "github.com/purwandi/hazelapp/project/storage"
	"github.com/sirupsen/logrus"
)

type Resolver struct {
	ProjectService   *ServiceProject.ProjectService
	IssueService     *ServiceIssue.IssueService
	MilestoneService *ServiceIssue.MilestoneService
}

func NewResolver() *Resolver {
	resolver := &Resolver{}

	dbDriver := os.Getenv("DB_DRIVER")
	switch dbDriver {
	default:
		logrus.Fatal("DB driver is not configurable yet")
	case "inmemory":
		ProjectStorage := StorageProject.NewProjectStorage()
		IssueStorage := StorageIssue.NewIssueStorage()
		MilestoneStorage := StorageIssue.NewMilestoneStorage()

		resolver.ProjectService = ServiceProject.NewProjectService(
			RepositoryInMemoryProject.NewProjectQueryInMemory(ProjectStorage),
			RepositoryInMemoryProject.NewProjectRepositoryInMemory(ProjectStorage),
		)

		resolver.IssueService = ServiceIssue.NewIssueService(
			RepositoryInMemoryIssue.NewIssueQueryInMemory(IssueStorage),
			RepositoryInMemoryIssue.NewIssueRepositoryInMemory(IssueStorage),
		)

		resolver.MilestoneService = ServiceIssue.NewMilestoneService(
			RepositoryInMemoryIssue.NewMilestoneQueryInMemory(MilestoneStorage),
			RepositoryInMemoryIssue.NewMilestoneRepositoryInMemory(MilestoneStorage),
		)

	}

	return resolver
}
