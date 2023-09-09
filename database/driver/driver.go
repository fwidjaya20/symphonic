package driver

import (
	"database/sql"
	"github.com/fwidjaya20/go-framework/contracts/config"
	"github.com/golang-migrate/migrate/v4/database"
	"log"
)

type DatabaseDriver interface {
	GetDSN() string
	Open() (*sql.DB, error)
	GetInstance() (database.Driver, error)
}

func GetDatabaseDriver(config config.Config) DatabaseDriver {
	switch config.Env("database.default") {
	case "postgresql":
		return NewPostgreSqlDriver(config)
	default:
		log.Fatalln("database driver only support postgresql")
		return nil
	}
}
