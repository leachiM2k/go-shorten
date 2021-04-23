package dataservice

import (
	"math/rand"
	"sync"
	"time"
)

type backend struct {
	mutex       sync.RWMutex
	entityCache map[string]Entity
	statCache   map[int]StatEntity
}

func (m *backend) All(owner string) (*[]*Entity, error) {
	// TODO: implement me
	panic("implement me")
}

func NewInmemoryBackend() Backend {
	rand.Seed(time.Now().UnixNano())
	return &backend{
		entityCache: map[string]Entity{},
		statCache:   map[int]StatEntity{},
	}
}

func (m *backend) Create(request CreateRequest) (*Entity, error) {
	entity := Entity{
		ID:         rand.Int(),
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

func (m *backend) CreateStat(shortenerId int, clientIp string, userAgent string, referer string) (*StatEntity, error) {
	entity := StatEntity{
		ShortenerID: shortenerId,
		ClientIP:    clientIp,
		UserAgent:   userAgent,
		Referer:     referer,
		CreatedAt:   time.Now(),
	}
	m.mutex.Lock()
	m.statCache[shortenerId] = entity
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

func (m *backend) Delete(owner string, code string) error {
	m.mutex.Lock()
	delete(m.entityCache, code)
	m.mutex.Unlock()

	return nil
}
