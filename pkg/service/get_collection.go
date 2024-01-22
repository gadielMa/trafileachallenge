package service

import (
	"context"
	"errors"
	"trafilea/cmd/logger"
	"trafilea/pkg/repository"
)

type GetCollectionService struct {
	ReadCacheRepository repository.ReadCacheRepository
}

func NewGetCollectionService(readCacheRepository repository.ReadCacheRepository) GetCollectionService {
	return GetCollectionService{
		ReadCacheRepository: readCacheRepository,
	}
}

func (s *GetCollectionService) Get(ctx context.Context, collectionID string) (*repository.Collection, error) {
	logger.Info(ctx, "starting...")

	collection, ok := s.ReadCacheRepository.Read(collectionID)
	if !ok {
		return nil, errors.New("the collection does not exist")
	}

	return collection, nil
}
