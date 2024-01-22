package service

import (
	"trafilea/pkg/repository"
)

type ReadCacheRepositoryMock struct {
	times    int
	id       string
	response *repository.Collection
	boolean  bool
}

func (m *ReadCacheRepositoryMock) Read(id string) (*repository.Collection, bool) {
	m.times++
	m.id = id

	return m.response, m.boolean
}

type UpdateCacheRepositoryMock struct {
	times      int
	id         string
	collection repository.Collection
}

func (m *UpdateCacheRepositoryMock) Update(id string, collection repository.Collection) {
	m.times++
	m.id = id
	m.collection = collection

	return
}
