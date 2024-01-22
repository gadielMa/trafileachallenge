package service

import "trafilea/pkg/repository"

type UpdateCacheRepository interface {
	Update(id string, collection repository.Collection)
}

type ReadCacheRepository interface {
	Read(id string) (*repository.Collection, bool)
}
