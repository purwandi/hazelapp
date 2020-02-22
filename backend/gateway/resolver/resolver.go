package resolver

import (
	"os"

	"github.com/purwandi/hazelapp/project/repository/inmemory"
	"github.com/purwandi/hazelapp/project/services"
	"github.com/purwandi/hazelapp/project/storage"
	"github.com/sirupsen/logrus"
)

type Resolver struct {
	ProjectService *services.ProjectService
}

func NewResolver() *Resolver {
	resolver := &Resolver{}

	dbDriver := os.Getenv("DB_DRIVER")
	switch dbDriver {
	default:
		logrus.Fatal("DB driver is not configurable yet")
	case "inmemory":
		storage := storage.NewProjectStorage()

		resolver.ProjectService = services.NewProjectService(
			inmemory.NewProjectQueryInMemory(storage),
			inmemory.NewProjectRepositoryInMemory(storage),
		)
	}

	return resolver
}
