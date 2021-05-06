package dataservice

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"os"
	"path"
	"strconv"
	"time"
)

const baseDir = "./.data"

type fileBackend struct{}

func (m *fileBackend) All(owner string) (*[]*Entity, error) {
	entities := make([]*Entity, 0)

	dirEntries, err := os.ReadDir(path.Join(baseDir, owner[:2], owner))
	if err != nil {
		return &entities, nil
	}

	for _, entry := range dirEntries {
		if entry.IsDir() {
			entity, err := m.Read(entry.Name())
			if err != nil {
				continue
			}
			entities = append(entities, entity)
		}
	}

	return &entities, nil
}

func NewFileBackend() Backend {
	rand.Seed(time.Now().UnixNano())
	return &fileBackend{}
}

func buildPath(owner string, code string) string {
	return path.Join(baseDir, owner[:2], owner, code)
}

func (m *fileBackend) Create(request CreateRequest) (*Entity, error) {
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

	marshal, err := json.Marshal(entity)
	if err != nil {
		return nil, err
	}

	filePath := buildPath(entity.Owner, entity.Code)

	err = os.MkdirAll(filePath, 0777)
	if err != nil {
		return nil, err
	}

	err = ioutil.WriteFile(path.Join(filePath, "shorten.json"), marshal, 0644)
	if err != nil {
		return nil, err
	}

	// write owner file
	err = os.MkdirAll(path.Join(baseDir, "codeowner", entity.Code[:2]), 0777)
	if err != nil {
		return nil, err
	}
	err = ioutil.WriteFile(path.Join(baseDir, "codeowner", entity.Code[:2], entity.Code), []byte(entity.Owner), 0644)
	if err != nil {
		return nil, err
	}

	return &entity, nil
}

func (m *fileBackend) CreateStat(shortenerId int, clientIp string, userAgent string, referer string) (*StatEntity, error) {
	entity := StatEntity{
		ShortenerID: shortenerId,
		ClientIP:    clientIp,
		UserAgent:   userAgent,
		Referer:     referer,
		CreatedAt:   time.Now(),
	}

	dirPath := path.Join(baseDir, "stats", strconv.Itoa(shortenerId)[:2], strconv.Itoa(shortenerId))
	err := os.MkdirAll(dirPath, 0777)
	if err != nil {
		return nil, err
	}

	marshal, err := json.Marshal(entity)
	if err != nil {
		return nil, err
	}

	err = ioutil.WriteFile(path.Join(dirPath, strconv.Itoa(int(time.Now().UnixNano()))+".json"), marshal, 0644)
	if err != nil {
		return nil, err
	}

	return &entity, nil
}

func (m *fileBackend) AllStats(code string) (*[]*StatEntity, error) {
	entity, err := m.Read(code)
	if err != nil {
		return nil, err
	}

	dirPath := path.Join(baseDir, "stats", strconv.Itoa(entity.ID)[:2], strconv.Itoa(entity.ID))
	dirEntries, err := os.ReadDir(dirPath)
	if err != nil {
		return nil, err
	}

	entities := make([]*StatEntity, 0)

	for _, entry := range dirEntries {
		bytes, err := ioutil.ReadFile(path.Join(dirPath, entry.Name()))
		if err != nil {
			return nil, nil
		}

		entity := StatEntity{}

		err = json.Unmarshal(bytes, &entity)
		if err != nil {
			return nil, err
		}

		entities = append(entities, &entity)
	}

	return &entities, nil
}

func (m *fileBackend) Read(code string) (*Entity, error) {
	ownerBytes, err := ioutil.ReadFile(path.Join(baseDir, "codeowner", code[:2], code))
	if err != nil {
		return nil, nil
	}

	filePath := path.Join(buildPath(string(ownerBytes), code), "shorten.json")

	bytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, nil
	}

	entity := Entity{}

	err = json.Unmarshal(bytes, &entity)
	if err != nil {
		return nil, err
	}

	return &entity, nil
}

func (m *fileBackend) Update(entity *Entity) (*Entity, error) {
	marshal, err := json.Marshal(entity)
	if err != nil {
		return nil, err
	}

	filePath := buildPath(entity.Owner, entity.Code)

	err = os.MkdirAll(filePath, 0777)
	if err != nil {
		return nil, err
	}

	err = ioutil.WriteFile(path.Join(filePath, "shorten.json"), marshal, 0644)
	if err != nil {
		return nil, err
	}

	return entity, nil
}

func (m *fileBackend) Delete(owner string, code string) error {
	filePath := buildPath(owner, code)
	err := os.RemoveAll(filePath)
	if err != nil {
		return err
	}

	err = os.RemoveAll(path.Join(baseDir, "codeowner", code[:2], code))
	if err != nil {
		return err
	}

	return nil
}
