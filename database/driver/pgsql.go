package driver

import (
	"database/sql"
	"fmt"

	"github.com/fwidjaya20/go-framework/contracts/config"
	"github.com/golang-migrate/migrate/v4/database"
	"github.com/golang-migrate/migrate/v4/database/postgres"
)

type Pgsql struct {
	config config.Config
}

func NewPostgreSqlDriver(config config.Config) DatabaseDriver {
	return &Pgsql{
		config: config,
	}
}

func (driver *Pgsql) GetDSN() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=%s",
		driver.config.Env("database.connections.postgresql.host"),
		driver.config.Env("database.connections.postgresql.port"),
		driver.config.Env("database.connections.postgresql.username"),
		driver.config.Env("database.connections.postgresql.password"),
		driver.config.Env("database.connections.postgresql.database"),
		driver.config.Env("database.timezone"),
	)
}

func (driver *Pgsql) GetInstance(table string) (database.Driver, error) {
	conn, err := driver.Open()
	if nil != err {
		return nil, err
	}

	return postgres.WithInstance(conn, &postgres.Config{
		MigrationsTable: table,
	})
}

func (driver *Pgsql) Open() (*sql.DB, error) {
	return sql.Open("postgres", driver.GetDSN())
}
