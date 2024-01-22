package service

import (
	"context"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
	"trafilea/cmd/logger"
	"trafilea/pkg/repository"
)

func TestGetCollectionService_Ok(t *testing.T) {
	ctx := context.WithValue(context.Background(), logger.RequestID, uuid.New().String())

	readCacheRepositoryMock := ReadCacheRepositoryMock{response: &repository.Collection{ID: "collection-1"}, boolean: true}

	service := NewGetCollectionService(&readCacheRepositoryMock)

	collection, err := service.Get(ctx, "collection-1")

	assert.Equal(t, 1, readCacheRepositoryMock.times)
	assert.Empty(t, err)
	assert.Equal(t, "collection-1", collection.ID)
}

func TestGetCollectionService_Fail(t *testing.T) {
	ctx := context.WithValue(context.Background(), logger.RequestID, uuid.New().String())

	readCacheRepositoryMock := ReadCacheRepositoryMock{response: &repository.Collection{ID: "collection-1"}, boolean: false}

	service := NewGetCollectionService(&readCacheRepositoryMock)

	collection, err := service.Get(ctx, "collection-1")

	assert.Equal(t, 1, readCacheRepositoryMock.times)
	assert.NotEmpty(t, err)
	assert.Empty(t, collection)
}
