package console

//nolint:revive // ignore due to golang-mgirate requirement
import (
	"fmt"
	"os"

	"github.com/fwidjaya20/symphonic/constant"
	ContractConfig "github.com/fwidjaya20/symphonic/contracts/config"
	"github.com/fwidjaya20/symphonic/database/driver"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func getMigrate(config ContractConfig.Config) (*migrate.Migrate, error) {
	rootDir, _ := os.Getwd()
	dbDriver := driver.GetDatabaseDriver(config)

	artifactsDir := fmt.Sprintf(
		"%s/%s/%s",
		rootDir,
		config.Get("database.dir", constant.DefaultDatabasePath),
		constant.DefaultMigrationDir,
	)
	databaseName := config.Get(
		fmt.Sprintf("database.connections.%s.database", config.Get("database.default")),
	)

	entries, err := os.ReadDir(artifactsDir)
	if err != nil {
		return nil, err
	}

	if len(entries) == 0 {
		return nil, ErrEmptyMigrationDir
	}

	instance, err := dbDriver.GetInstance("schema_migrations")
	if err != nil {
		return nil, err
	}

	return migrate.NewWithDatabaseInstance(
		fmt.Sprintf("file://%s", artifactsDir),
		databaseName.(string),
		instance,
	)
}
