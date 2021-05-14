package dataservice

import (
	"github.com/leachim2k/go-shorten/pkg/cli/shorten/options"
	b "github.com/leachim2k/go-shorten/pkg/dataservice/backend"
	i "github.com/leachim2k/go-shorten/pkg/dataservice/interfaces"
	"github.com/mrcrgl/pflog/log"
)

var backendMap = map[string]func() i.Backend{
	options.BackendFile:       func() i.Backend { return b.NewFileBackend() },
	options.BackendInMemory:   func() i.Backend { return b.NewInmemoryBackend() },
	options.BackendPostgreSQL: func() i.Backend { return b.NewDBBackend() },
}

func GetDataService(key string) i.Backend {
	if val, ok := backendMap[key]; ok {
		log.Infof("using data service backend: %s", key)
		return val()
	}
	log.Fatalf("no data service backend found for: %s", key)
	return nil
}

func GetDataServiceByConfig() i.Backend {
	return GetDataService(options.Current.Storage.Engine)
}
