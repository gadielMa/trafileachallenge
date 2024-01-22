package service

import (
	"context"
	"trafilea/cmd/logger"
	"trafilea/pkg/repository"
)

type CreateNumberService struct {
	ReadCacheRepository   repository.ReadCacheRepository
	UpdateCacheRepository repository.UpdateCacheRepository
}

func NewCreateNumberService(readCacheRepository repository.ReadCacheRepository,
	updateCacheRepository repository.UpdateCacheRepository) CreateNumberService {
	return CreateNumberService{
		ReadCacheRepository:   readCacheRepository,
		UpdateCacheRepository: updateCacheRepository,
	}
}

func (s *CreateNumberService) Create(ctx context.Context, collection string, number int) {
	logger.Info(ctx, "starting...")

	collectionP, ok := s.ReadCacheRepository.Read(collection)
	if ok {
		collectionP.Numbers = append(collectionP.Numbers, number)
		s.UpdateCacheRepository.Update(collection, *collectionP)
		return
	}

	s.UpdateCacheRepository.Update(collection, repository.Collection{ID: collection, Numbers: []int{number}})
	return
}
