package dataservice

import (
	"github.com/leachim2k/go-shorten/pkg/cli/shorten/options"
	"github.com/mrcrgl/pflog/log"
)

type Backend interface {
	Create(request CreateRequest) (*Entity, error)
	CreateStat(shortenerId int, clientIp string, userAgent string, referer string) (*StatEntity, error)
	Read(code string) (*Entity, error)
	Update(entity *Entity) (*Entity, error)
	Delete(owner string, code string) error
	All(owner string) (*[]*Entity, error)
	AllStats(code string) (*[]*StatEntity, error)
}

var backendMap = map[string]func() Backend{
	options.BackendFile:       func() Backend { return NewFileBackend() },
	options.BackendInMemory:   func() Backend { return NewInmemoryBackend() },
	options.BackendPostgreSQL: func() Backend { return NewDBBackend() },
}

func GetDataService(key string) Backend {
	if val, ok := backendMap[key]; ok {
		log.Infof("using data service backend: %s", key)
		return val()
	}
	log.Fatalf("no data service backend found for: %s", key)
	return nil
}

func GetDataServiceByConfig() Backend {
	return GetDataService(options.Current.StorageBackend)
}
