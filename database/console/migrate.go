package console

import (
	"database/sql"
	"fmt"
	"github.com/fwidjaya20/go-framework/contracts/config"
	"github.com/fwidjaya20/go-framework/database/driver"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"log"
)

func getMigrate(config config.Config) (*migrate.Migrate, error) {
	dbDriver := driver.GetDatabaseDriver(config)

	artifactsDir := config.Env("database.dir", "./database")
	databaseName := config.Env(fmt.Sprintf("database.connections.%s.database", config.Env("database.default")))

	conn, err := sql.Open("postgres", dbDriver.GetDSN())
	if nil != err {
		log.Fatalln(err.Error())
	}

	instance, err := postgres.WithInstance(conn, &postgres.Config{
		MigrationsTable: fmt.Sprintf("%s/migrations", config.Env("database.dir")),
	})
	if nil != instance {
		log.Fatalln(err.Error())
	}

	return migrate.NewWithDatabaseInstance(fmt.Sprintf("file://%s", artifactsDir), databaseName.(string), instance)
}
