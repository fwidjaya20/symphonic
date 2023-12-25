package driver

import (
	"database/sql"
	"log"

	ContractConfig "github.com/fwidjaya20/symphonic/contracts/config"
	"github.com/golang-migrate/migrate/v4/database"
)

type DatabaseDriver interface {
	GetDSN() string
	Open() (*sql.DB, error)
	GetInstance(table string) (database.Driver, error)
}

func GetDatabaseDriver(config ContractConfig.Config) DatabaseDriver {
	switch config.Get("database.default") {
	case "postgresql":
		return NewPostgreSQLDriver(config)
	default:
		log.Fatalln("database driver only support postgresql")
		return nil
	}
}
