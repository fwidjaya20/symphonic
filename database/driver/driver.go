package driver

import (
	"database/sql"
	"log"

	"github.com/fwidjaya20/symphonic/contracts/config"
	"github.com/golang-migrate/migrate/v4/database"
)

type DatabaseDriver interface {
	GetDSN() string
	Open() (*sql.DB, error)
	GetInstance(table string) (database.Driver, error)
}

func GetDatabaseDriver(config config.Config) DatabaseDriver {
	switch config.Get("database.default") {
	case "postgresql":
		return NewPostgreSqlDriver(config)
	default:
		log.Fatalln("database driver only support postgresql")
		return nil
	}
}
