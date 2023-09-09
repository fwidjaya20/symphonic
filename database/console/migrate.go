package console

import (
	"fmt"
	"github.com/fwidjaya20/go-framework/constant"
	"github.com/fwidjaya20/go-framework/contracts/config"
	"github.com/fwidjaya20/go-framework/database/driver"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"os"
)

func getMigrate(config config.Config) (*migrate.Migrate, error) {
	rootDir, _ := os.Getwd()
	dbDriver := driver.GetDatabaseDriver(config)

	artifactsDir := fmt.Sprintf("%s/%s/%s", rootDir, config.Env("database.dir", constant.DefaultDatabasePath), constant.DefaultMigrationDir)
	databaseName := config.Env(fmt.Sprintf("database.connections.%s.database", config.Env("database.default")))

	instance, err := dbDriver.GetInstance()
	if nil != err {
		return nil, err
	}

	return migrate.NewWithDatabaseInstance(fmt.Sprintf("file://%s", artifactsDir), databaseName.(string), instance)
}
