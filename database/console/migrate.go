package console

import (
	"fmt"
	"os"

	"github.com/fwidjaya20/symphonic/constant"
	"github.com/fwidjaya20/symphonic/contracts/config"
	"github.com/fwidjaya20/symphonic/database/migration"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func getMigrate(config config.Config) (*migrate.Migrate, error) {
	rootDir, _ := os.Getwd()
	dbDriver := migration.GetDatabaseDriver(config)

	artifactsDir := fmt.Sprintf("%s/%s/%s", rootDir, config.Get("database.dir", constant.DefaultDatabasePath), constant.DefaultMigrationDir)
	databaseName := config.Get(fmt.Sprintf("database.connections.%s.database", config.Get("database.default")))

	entries, err := os.ReadDir(artifactsDir)
	if nil != err {
		return nil, err
	}

	if len(entries) == 0 {
		return nil, ErrEmptyMigrationDir
	}

	instance, err := dbDriver.GetInstance("schema_migrations")
	if nil != err {
		return nil, err
	}

	return migrate.NewWithDatabaseInstance(fmt.Sprintf("file://%s", artifactsDir), databaseName.(string), instance)
}
