package shorten

import (
	"github.com/jonboulle/clockwork"
	"github.com/leachim2k/go-shorten/pkg/dataservice"
	"github.com/mrcrgl/pflog/log"
	"time"
)

type Handler struct {
	Clock   clockwork.Clock
	Backend dataservice.Backend
}

func NewHandler(clock clockwork.Clock, backend dataservice.Backend) *Handler {
	return &Handler{
		Clock:   clock,
		Backend: backend,
	}
}

func (m *Handler) Get(code string) (*dataservice.Entity, error) {
	entity, err := m.Backend.Read(code)
	if err != nil {
		log.Infof("error during read from backend: %s", err)
		return nil, err
	}

	if entity == nil {
		log.Infof("could not find entity for code %s", code)
	}

	return entity, nil
}

func (m *Handler) ConvertEntityToLink(entity *dataservice.Entity) (string, error) {
	if entity == nil {
		return "", nil
	}

	if entity.ExpiresAt != nil && entity.ExpiresAt.Before(time.Now()) {
		return "", nil
	}

	if entity.StartTime != nil && entity.StartTime.After(time.Now()) {
		return "", nil
	}
	/*
		m.Mutex.Lock()
		entity.Count += 1
		m.EntityCache[entity.Code] = *entity
		m.Mutex.Unlock()
	*/
	return entity.Link, nil
}

func (m *Handler) Add(request dataservice.CreateRequest) (*dataservice.Entity, error) {
	entity, err := m.Backend.Create(request)
	if err != nil {
		return nil, err
	}
	return entity, nil
}

func (m *Handler) Delete(code string) error {
	return m.Backend.Delete(code)
}

func (m *Handler) Update(code string, request dataservice.UpdateRequest) (*dataservice.Entity, error) {
	entity, err := m.Backend.Read(code)
	if entity == nil || err != nil {
		return nil, err
	}

	entity.Link = request.Link
	entity.MaxCount = request.MaxCount
	entity.StartTime = request.StartTime
	entity.ExpiresAt = request.ExpiresAt
	entity.UpdatedAt = time.Now()

	if request.Attributes != nil {
		if entity.Attributes == nil {
			entity.Attributes = new(map[string]interface{})
		}
		for k, v := range request.Attributes {
			(*entity.Attributes)[k] = v
		}
	}

	_, err = m.Backend.Update(entity)
	if err != nil {
		return nil, err
	}

	return entity, nil
}
