package resolver

import (
	"os"

	RepositoryInMemoryIssue "github.com/purwandi/hazelapp/issue/repository/inmemory"
	ServiceIssue "github.com/purwandi/hazelapp/issue/services"
	StorageIssue "github.com/purwandi/hazelapp/issue/storage"
	RepositoryInMemoryProject "github.com/purwandi/hazelapp/project/repository/inmemory"
	ServiceProject "github.com/purwandi/hazelapp/project/services"
	StorageProject "github.com/purwandi/hazelapp/project/storage"
	RepositoryInMemoryUser "github.com/purwandi/hazelapp/user/repository/inmemory"
	ServiceUser "github.com/purwandi/hazelapp/user/services"
	StorageUser "github.com/purwandi/hazelapp/user/storage"
	"github.com/sirupsen/logrus"
)

// Resolver is root resolver
type Resolver struct {
	UserService      *ServiceUser.UserService
	ProjectService   *ServiceProject.ProjectService
	IssueService     *ServiceIssue.IssueService
	MilestoneService *ServiceIssue.MilestoneService
}

// NewResolver is to create new Resolver instance
func NewResolver() *Resolver {
	resolver := &Resolver{}

	dbDriver := os.Getenv("DB_DRIVER")
	switch dbDriver {
	default:
		logrus.Fatal("DB driver is not configurable yet")
	case "inmemory":
		UserStorage := StorageUser.NewUserStorage()
		ProjectStorage := StorageProject.NewProjectStorage()
		MilestoneStorage := StorageIssue.NewMilestoneStorage()
		IssueStorage := StorageIssue.NewIssueStorage()

		UserStorage.Demo()
		ProjectStorage.Demo()
		IssueStorage.Demo()

		resolver.ProjectService = ServiceProject.NewProjectService(
			RepositoryInMemoryProject.NewProjectQueryInMemory(ProjectStorage),
			RepositoryInMemoryProject.NewProjectRepositoryInMemory(ProjectStorage),
		)

		resolver.UserService = ServiceUser.NewUserService(
			RepositoryInMemoryUser.NewUserQueryInMemory(UserStorage),
			RepositoryInMemoryUser.NewUserRepositoryInMemory(UserStorage),
		)

		resolver.MilestoneService = ServiceIssue.NewMilestoneService(
			RepositoryInMemoryIssue.NewMilestoneQueryInMemory(MilestoneStorage),
			RepositoryInMemoryIssue.NewMilestoneRepositoryInMemory(MilestoneStorage),
		)

		resolver.IssueService = ServiceIssue.NewIssueService(
			RepositoryInMemoryIssue.NewIssueQueryInMemory(IssueStorage),
			RepositoryInMemoryIssue.NewIssueRepositoryInMemory(IssueStorage),
		)

	}

	return resolver
}
