package backend

import (
	"github.com/leachim2k/go-shorten/pkg/dataservice/interfaces"
	"math/rand"
	"sync"
	"time"
)

type backend struct {
	mutex       sync.RWMutex
	entityCache map[string]interfaces.Entity
	statCache   map[int][]*interfaces.StatEntity
}

func (m *backend) All(owner string) (*[]*interfaces.Entity, error) {
	// TODO: implement me
	panic("implement me")
}

func NewInmemoryBackend() interfaces.Backend {
	rand.Seed(time.Now().UnixNano())
	return &backend{
		entityCache: map[string]interfaces.Entity{},
		statCache:   map[int][]*interfaces.StatEntity{},
	}
}

func (m *backend) Create(request interfaces.CreateRequest) (*interfaces.Entity, error) {
	entity := interfaces.Entity{
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

func (m *backend) CreateStat(shortenerId int, clientIp string, userAgent string, referer string) (*interfaces.StatEntity, error) {
	entity := interfaces.StatEntity{
		ShortenerID: shortenerId,
		ClientIP:    clientIp,
		UserAgent:   userAgent,
		Referer:     referer,
		CreatedAt:   time.Now(),
	}
	m.mutex.Lock()
	stats, ok := m.statCache[shortenerId]
	if !ok {
		stats = make([]*interfaces.StatEntity, 0)
	}
	m.statCache[shortenerId] = append(stats, &entity)
	m.mutex.Unlock()

	return &entity, nil
}

func (m *backend) AllStats(code string) (*[]*interfaces.StatEntity, error) {
	entity, err := m.Read(code)
	if err != nil {
		return nil, err
	}

	m.mutex.RLock()
	stats, ok := m.statCache[entity.ID]
	m.mutex.RUnlock()

	if !ok {
		return nil, nil
	}

	return &stats, nil
}

func (m *backend) Read(code string) (*interfaces.Entity, error) {
	m.mutex.RLock()
	entity, ok := m.entityCache[code]
	m.mutex.RUnlock()

	if !ok {
		return nil, nil
	}

	return &entity, nil
}

func (m *backend) Update(entity *interfaces.Entity) (*interfaces.Entity, error) {
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
