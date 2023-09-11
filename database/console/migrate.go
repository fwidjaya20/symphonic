package console

import (
	"fmt"
	"os"

	"github.com/fwidjaya20/go-framework/constant"
	"github.com/fwidjaya20/go-framework/contracts/config"
	"github.com/fwidjaya20/go-framework/database/driver"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func getMigrate(config config.Config) (*migrate.Migrate, error) {
	rootDir, _ := os.Getwd()
	dbDriver := driver.GetDatabaseDriver(config)

	artifactsDir := fmt.Sprintf("%s/%s/%s", rootDir, config.Get("database.dir", constant.DefaultDatabasePath), constant.DefaultMigrationDir)
	databaseName := config.Get(fmt.Sprintf("database.connections.%s.database", config.Get("database.default")))

	instance, err := dbDriver.GetInstance("schema_migrations")
	if nil != err {
		return nil, err
	}

	return migrate.NewWithDatabaseInstance(fmt.Sprintf("file://%s", artifactsDir), databaseName.(string), instance)
}
