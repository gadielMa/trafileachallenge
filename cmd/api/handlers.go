package api

import (
	"context"
	"trafilea/pkg/repository"
)

type DivisibleService interface {
	Prints(ctx context.Context) []string
}

type ValueNumberService interface {
	Process(ctx context.Context, number int) string
}

type CreateNumberService interface {
	Create(ctx context.Context, collection string, number int)
}

type GetCollectionService interface {
	Get(ctx context.Context, collectionID string) (*repository.Collection, error)
}
