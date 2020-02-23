package inmemory

import (
	"github.com/purwandi/hazelapp/issue/domain"
	"github.com/purwandi/hazelapp/issue/repository"
	"github.com/purwandi/hazelapp/issue/storage"
)

type MilestoneRepositoryInMemory struct {
	Storage *storage.MilestoneStorage
}

func NewMilestoneRepositoryInMemory(s *storage.MilestoneStorage) repository.MilestoneRepository {
	return &MilestoneRepositoryInMemory{Storage: s}
}

func (repo *MilestoneRepositoryInMemory) Save(milestone *domain.Milestone) <-chan error {
	result := make(chan error)

	go func() {
		repo.Storage.MilestoneMap = append(repo.Storage.MilestoneMap, *milestone)

		result <- nil
		close(result)
	}()

	return result
}
