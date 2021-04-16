package dataservice

import (
	"sync"
	"time"
)

type backend struct {
	mutex       sync.RWMutex
	entityCache map[string]Entity
}

func (m *backend) All(owner string) (*[]*Entity, error) {
	// TODO: implement me
	panic("implement me")
}

func NewInmemoryBackend() Backend {
	return &backend{
		entityCache: map[string]Entity{},
	}
}

func (m *backend) Create(request CreateRequest) (*Entity, error) {
	entity := Entity{
		Owner:      *request.Owner,
		Link:       *request.Link,
		Code:       request.Code,
		Count:      0,
		MaxCount:   request.MaxCount,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
		StartTime:  request.StartTime,
		ExpiresAt:  request.ExpiresAt,
		Attributes: request.Attributes,
	}
	m.mutex.Lock()
	m.entityCache[request.Code] = entity
	m.mutex.Unlock()

	return &entity, nil
}

func (m *backend) Read(code string) (*Entity, error) {
	m.mutex.RLock()
	entity, ok := m.entityCache[code]
	m.mutex.RUnlock()

	if !ok {
		return nil, nil
	}

	return &entity, nil
}

func (m *backend) Update(entity *Entity) (*Entity, error) {
	m.mutex.Lock()
	m.entityCache[entity.Code] = *entity
	m.mutex.Unlock()

	return entity, nil
}

func (m *backend) Delete(code string) error {
	m.mutex.Lock()
	delete(m.entityCache, code)
	m.mutex.Unlock()

	return nil
}
