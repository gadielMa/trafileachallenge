package service

import (
	"context"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
	"trafilea/cmd/logger"
)

func TestCreateNumberService_Create(t *testing.T) {
	ctx := context.WithValue(context.Background(), logger.RequestID, uuid.New().String())

	readCacheRepositoryMock := ReadCacheRepositoryMock{}
	updateCacheRepositoryMock := UpdateCacheRepositoryMock{}

	service := NewCreateNumberService(&readCacheRepositoryMock, &updateCacheRepositoryMock)

	service.Create(ctx, "collection-1", 2)

	assert.Equal(t, 1, readCacheRepositoryMock.times)
	assert.Equal(t, 1, updateCacheRepositoryMock.times)
}
