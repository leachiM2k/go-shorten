package dataservice

import (
	"github.com/leachim2k/go-shorten/pkg/cli/shorten/options"
	"github.com/mrcrgl/pflog/log"
)

type Backend interface {
	Create(request CreateRequest) (*Entity, error)
	Read(code string) (*Entity, error)
	Update(entity *Entity) (*Entity, error)
	Delete(code string) error
}

var backendMap = map[string]Backend{
	options.BackendInMemory:   NewInmemoryBackend(),
	options.BackendPostgreSQL: NewDBBackend(),
}

func GetDataService(key string) Backend {
	if val, ok := backendMap[key]; ok {
		log.Infof("using data service backend: %s", key)
		return val
	}
	log.Warningf("no data service backend found for: %s", key)
	return nil
}

func GetDataServiceByConfig() Backend {
	return GetDataService(options.Current.StorageBackend)
}
