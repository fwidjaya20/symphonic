package driver

import (
	"github.com/fwidjaya20/go-framework/contracts/config"
	"log"
)

type DatabaseDriver interface {
	GetDSN() string
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
