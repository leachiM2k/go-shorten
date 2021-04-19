package dataservice

import (
	"github.com/leachim2k/go-shorten/pkg/cli/shorten/options"
	"github.com/leachim2k/go-shorten/pkg/models"
	logOriginal "log"
	"os"
	"time"
)

type dbBackend struct {
}

func NewDBBackend() Backend {
	models.InitDB("postgres", options.Current.DBConnection, logOriginal.New(os.Stdout, "[sql] ", logOriginal.LstdFlags))

	return &dbBackend{}
}

func ConvertDbItemToEntity(dbItem *models.ShortenerItem) *Entity {
	return &Entity{
		Owner:      dbItem.Owner,
		Link:       dbItem.Link,
		Code:       dbItem.Code,
		Count:      dbItem.Count,
		MaxCount:   dbItem.MaxCount,
		CreatedAt:  dbItem.CreatedAt,
		UpdatedAt:  dbItem.UpdatedAt,
		StartTime:  dbItem.StartTime,
		ExpiresAt:  dbItem.ExpiresAt,
		Attributes: (*map[string]interface{})(dbItem.Attributes),
	}
}

func (m *dbBackend) Create(request CreateRequest) (*Entity, error) {
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
	err := models.AddShort(dbItem)
	if err != nil {
		return nil, err
	}

	return ConvertDbItemToEntity(&dbItem), nil
}

func (m *dbBackend) All(owner string) (*[]*Entity, error) {
	items, err := models.AllShortsByOwner(owner)
	if err != nil {
		return nil, err
	}

	entities := make([]*Entity, len(items))
	for i, item := range items {
		entities[i] = ConvertDbItemToEntity(item)
	}

	return &entities, nil
}

func (m *dbBackend) Read(code string) (*Entity, error) {
	entity, err := models.GetShortByCode(code)
	if entity == nil || err != nil {
		return nil, err
	}

	return ConvertDbItemToEntity(entity), nil
}

func (m *dbBackend) Update(entity *Entity) (*Entity, error) {
	dbEntity := models.ShortenerItem{
		Code:       entity.Code,
		Link:       entity.Link,
		MaxCount:   entity.MaxCount,
		StartTime:  entity.StartTime,
		ExpiresAt:  entity.ExpiresAt,
		Attributes: (*models.Attributes)(entity.Attributes),
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
