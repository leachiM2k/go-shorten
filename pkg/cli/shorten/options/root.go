package options

import (
	"github.com/pkg/errors"
	"os"
	"strconv"
)

// Current is the de-facto Config used.
var Current = ConfigWithDefaults()

const (
	BackendInMemory   = "memory"
	BackendPostgreSQL = "postgresql"
	BackendMySQL      = "mysql"
	BackendFile       = "file"
)

func getOr(this, or string) string {
	if len(this) == 0 {
		return or
	}
	return this
}

// ConfigWithDefaults sets default values and returns a new Config.
func ConfigWithDefaults() *Config {
	defaultPort := 80
	if i, err := strconv.ParseInt(os.Getenv("PORT"), 10, 32); err == nil {
		defaultPort = int(i)
	}
	return &Config{
		RESTListenPort: defaultPort,
		DBConnection:   getOr(os.Getenv("DB_CONNECTION"), "postgres://shorten:shorten@localhost:16541/shorten_dev?sslmode=disable"),
		StorageBackend: getOr(os.Getenv("STORAGE"), BackendPostgreSQL),
	}
}

// Config holds information that can be modified by config or command line flags.
type Config struct {
	StorageBackend string
	RESTListenPort int
	DBConnection   string
}

func (c *Config) Validate() error {
	if c.StorageBackend != BackendInMemory && c.StorageBackend != BackendPostgreSQL && c.StorageBackend != BackendMySQL && c.StorageBackend != BackendFile {
		return errors.Errorf("storage backend must be set")
	}
	if c.StorageBackend == BackendPostgreSQL || c.StorageBackend == BackendMySQL && len(c.DBConnection) == 0 {
		return errors.Errorf("db connection string must be set for backend type %s", c.StorageBackend)
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
