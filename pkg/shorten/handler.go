package shorten

import (
	"github.com/jonboulle/clockwork"
	"github.com/leachim2k/go-shorten/pkg/dataservice"
	"github.com/mrcrgl/pflog/log"
	"math/rand"
	"time"
)

type Handler struct {
	Clock   clockwork.Clock
	Backend dataservice.Backend
}

func NewHandler(clock clockwork.Clock, backend dataservice.Backend) *Handler {
	rand.Seed(time.Now().UnixNano())
	return &Handler{
		Clock:   clock,
		Backend: backend,
	}
}

func (m *Handler) GetAll(owner string) (*[]*dataservice.Entity, error) {
	entities, err := m.Backend.All(owner)
	if err != nil {
		log.Infof("error during read from backend: %s", err)
		return nil, err
	}

	if entities == nil {
		log.Infof("could not find entities for owner %s", owner)
	}

	return entities, nil
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

	if entity.MaxCount != 0 && entity.MaxCount >= entity.Count {
		return "", nil
	}

	entity.Count += 1

	go func(updateEntity *dataservice.Entity) {
		_, err := m.Backend.Update(updateEntity)
		if err != nil {
			log.Warningf("Could not update short count: %s", err)
		}
	}(entity)

	return entity.Link, nil
}

func (m *Handler) AddStat(shortenerId int, clientIp string, userAgent string, referer string) (*dataservice.StatEntity, error) {
	entity, err := m.Backend.CreateStat(shortenerId, clientIp, userAgent, referer)
	if err != nil {
		return nil, err
	}
	return entity, nil
}

func (m *Handler) Add(request dataservice.CreateRequest) (*dataservice.Entity, error) {
	if request.Code == "" {
		request.Code = GenerateRandomString(8)
	}
	entity, err := m.Backend.Create(request)
	if err != nil {
		return nil, err
	}
	return entity, nil
}

func (m *Handler) Delete(owner string, code string) error {
	return m.Backend.Delete(owner, code)
}

func (m *Handler) Update(code string, request dataservice.UpdateRequest) (*dataservice.Entity, error) {
	entity, err := m.Backend.Read(code)
	if entity == nil || err != nil {
		return nil, err
	}

	entity.Description = request.Description
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

func (m *Handler) GetStats(code string) (*[]*dataservice.StatEntity, error) {
	entities, err := m.Backend.AllStats(code)
	if err != nil {
		log.Infof("error during read stats from backend: %s", err)
		return nil, err
	}

	if entities == nil {
		log.Infof("could not find stats for code %s", code)
	}

	return entities, nil
}
