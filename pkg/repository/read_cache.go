package repository

import gocache "github.com/patrickmn/go-cache"

type ReadCacheRepository struct {
	cacheData CacheData
}

func NewReadCacheRepository(cache CacheData) ReadCacheRepository {
	return ReadCacheRepository{
		cacheData: cache,
	}
}

type CacheData struct {
	Collection *gocache.Cache
}

type Collection struct {
	ID      string `json:"id"`
	Numbers []int  `json:"numbers"`
}

func (c *ReadCacheRepository) Read(id string) (*Collection, bool) {
	collection, ok := c.cacheData.Collection.Get(id)
	if ok {
		collect := collection.(Collection)

		return &collect, true
	}

	return nil, false
}
