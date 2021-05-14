package backend

import (
	"github.com/leachim2k/go-shorten/pkg/cli/shorten/options"
	"github.com/leachim2k/go-shorten/pkg/dataservice/interfaces"
	"github.com/leachim2k/go-shorten/pkg/models"
	"github.com/mrcrgl/pflog/log"
	logOriginal "log"
	"os"
	"time"
)

type dbBackend struct {
}

func NewDBBackend() interfaces.Backend {
	models.InitDB("postgres", options.Current.Storage.DBUrl, logOriginal.New(os.Stdout, "[sql] ", logOriginal.LstdFlags))

	return &dbBackend{}
}

func ConvertDbItemToEntity(dbItem *models.ShortenerItem) *interfaces.Entity {
	return &interfaces.Entity{
		ID:          dbItem.ID,
		Owner:       dbItem.Owner,
		Link:        dbItem.Link,
		Code:        dbItem.Code,
		Description: dbItem.Description,
		Count:       dbItem.Count,
		MaxCount:    dbItem.MaxCount,
		CreatedAt:   dbItem.CreatedAt,
		UpdatedAt:   dbItem.UpdatedAt,
		StartTime:   dbItem.StartTime,
		ExpiresAt:   dbItem.ExpiresAt,
		Attributes:  (*map[string]interface{})(dbItem.Attributes),
	}
}

func ConvertStatDbItemToEntity(dbItem *models.ShortStatItem) *interfaces.StatEntity {
	return &interfaces.StatEntity{
		ClientIP:  dbItem.ClientIP,
		UserAgent: dbItem.UserAgent,
		Referer:   dbItem.Referer,
		CreatedAt: dbItem.CreatedAt,
	}
}

func (m *dbBackend) CreateStat(shortenerId int, clientIp string, userAgent string, referer string) (*interfaces.StatEntity, error) {
	dbItem := models.ShortStatItem{
		ShortenerID: shortenerId,
		ClientIP:    clientIp,
		UserAgent:   userAgent,
		Referer:     referer,
		CreatedAt:   time.Now(),
	}
	err := models.AddShortStat(dbItem)
	if err != nil {
		log.Infof("create stat failed with error: %#v", err)
		return nil, err
	}
	return &interfaces.StatEntity{
		ShortenerID: dbItem.ShortenerID,
		ClientIP:    dbItem.ClientIP,
		UserAgent:   dbItem.UserAgent,
		Referer:     dbItem.Referer,
		CreatedAt:   dbItem.CreatedAt,
	}, nil
}

func (m *dbBackend) Create(request interfaces.CreateRequest) (*interfaces.Entity, error) {
	dbItem := models.ShortenerItem{
		Link:        *request.Link,
		Owner:       *request.Owner,
		Code:        request.Code,
		Description: request.Description,
		Count:       0,
		MaxCount:    0,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		StartTime:   request.StartTime,
		ExpiresAt:   request.ExpiresAt,
		Attributes:  (*models.Attributes)(request.Attributes),
	}
	dbResult, err := models.AddShort(dbItem)
	if err != nil {
		return nil, err
	}

	return ConvertDbItemToEntity(dbResult), nil
}

func (m *dbBackend) All(owner string) (*[]*interfaces.Entity, error) {
	items, err := models.AllShortsByOwner(owner)
	if err != nil {
		return nil, err
	}

	entities := make([]*interfaces.Entity, len(items))
	for i, item := range items {
		entities[i] = ConvertDbItemToEntity(item)
	}

	return &entities, nil
}

func (m *dbBackend) AllStats(code string) (*[]*interfaces.StatEntity, error) {
	items, err := models.AllShortStats(code)
	if err != nil {
		return nil, err
	}

	entities := make([]*interfaces.StatEntity, len(items))
	for i, item := range items {
		entities[i] = ConvertStatDbItemToEntity(item)
	}

	return &entities, nil
}

func (m *dbBackend) Read(code string) (*interfaces.Entity, error) {
	entity, err := models.GetShortByCode(code)
	if entity == nil || err != nil {
		return nil, err
	}

	return ConvertDbItemToEntity(entity), nil
}

func (m *dbBackend) Update(entity *interfaces.Entity) (*interfaces.Entity, error) {
	dbEntity := models.ShortenerItem{
		Owner:       entity.Owner,
		Code:        entity.Code,
		Link:        entity.Link,
		Description: entity.Description,
		Count:       entity.Count,
		MaxCount:    entity.MaxCount,
		StartTime:   entity.StartTime,
		ExpiresAt:   entity.ExpiresAt,
		Attributes:  (*models.Attributes)(entity.Attributes),
	}
	_, err := models.UpdateShort(&dbEntity)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (m *dbBackend) Delete(owner string, code string) error {
	err := models.DeleteShortByCode(owner, code)
	if err != nil {
		return err
	}
	return nil
}
