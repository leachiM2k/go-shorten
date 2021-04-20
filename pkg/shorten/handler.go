package shorten

import (
	"github.com/jonboulle/clockwork"
	"github.com/leachim2k/go-shorten/pkg/dataservice"
	"github.com/mrcrgl/pflog/log"
	"math/rand"
	"strings"
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
	entity, err := m.Backend.All(owner)
	if err != nil {
		log.Infof("error during read from backend: %s", err)
		return nil, err
	}

	if entity == nil {
		log.Infof("could not find entities for owner %s", owner)
	}

	return entity, nil
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
	_, err := m.Backend.Update(entity)
	if err != nil {
		log.Warningf("Could not update short count: %s", err)
	}

	return entity.Link, nil
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

const (
	letterBytes   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890-_" // 52 possibilities
	letterIdxBits = 6                                                                  // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1                                               // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits                                                 // # of letter indices fitting in 63 bits
)

var src = rand.NewSource(time.Now().UnixNano())

func GenerateRandomString(n int) string {
	sb := strings.Builder{}
	sb.Grow(n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			sb.WriteByte(letterBytes[idx])
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return sb.String()
}
