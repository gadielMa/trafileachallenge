package repository

import (
	gocache "github.com/patrickmn/go-cache"
)

type UpdateCacheRepository struct {
	cacheData CacheData
}

func NewUpdateCacheRepository(cache CacheData) UpdateCacheRepository {
	return UpdateCacheRepository{
		cacheData: cache,
	}
}

func (c *UpdateCacheRepository) Update(id string, product Collection) {
	c.cacheData.Collection.Set(id, product, gocache.DefaultExpiration)
}
