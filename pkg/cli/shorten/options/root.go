package options

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/pkg/errors"
)

type Config struct {
	Server struct {
		Port int `yaml:"port" env:"PORT" env-default:"8080"`
	} `yaml:"server"`
	Storage struct {
		Engine string `yaml:"engine" env:"STORAGE" env-default:"postgresql"`
		DBUrl  string `yaml:"dburl" env:"DATABASE_URL" env-default:"postgres://shorten:shorten@localhost:16541/shorten_dev?sslmode=disable"`
	} `yaml:"storage"`
	AuthServices []struct {
		Name         string `yaml:"name"`
		ClientId     string `yaml:"clientid"`
		ClientSecret string `yaml:"clientsecret"`
		Prefix       string `yaml:"prefix"`
	} `yaml:"authservices"`
}

// Current is the de-facto Config used.
var Current = ConfigWithDefaults()

const (
	BackendInMemory   = "memory"
	BackendPostgreSQL = "postgresql"
	BackendMySQL      = "mysql"
	BackendFile       = "file"
)

// ConfigWithDefaults sets default values and returns a new Config.
func ConfigWithDefaults() *Config {
	var cfg Config
	err := cleanenv.ReadConfig("config.yml", &cfg)
	if err != nil {
		return nil
	}

	return &cfg
}

func (c *Config) Validate() error {
	if c.Storage.Engine != BackendInMemory && c.Storage.Engine != BackendPostgreSQL && c.Storage.Engine != BackendMySQL && c.Storage.Engine != BackendFile {
		return errors.Errorf("storage backend must be set")
	}
	if c.Storage.Engine == BackendPostgreSQL || c.Storage.Engine == BackendMySQL && len(c.Storage.DBUrl) == 0 {
		return errors.Errorf("db connection string must be set for backend type %s", c.Storage.Engine)
	}
	return nil
}

// Get returns always a pointer to a copy of Current.
func Get() *Config {
	c := Copy()
	return &c
}

// Copy returns an immutable Config.
func Copy() Config {
	return *Current
}
