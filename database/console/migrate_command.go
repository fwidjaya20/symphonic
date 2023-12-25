//nolint:dupl // ignore due to database migrate logic
package console

import (
	"errors"

	ContractConfig "github.com/fwidjaya20/symphonic/contracts/config"
	ContractConsole "github.com/fwidjaya20/symphonic/contracts/console"
	"github.com/golang-migrate/migrate/v4"
	"github.com/gookit/color"
	"github.com/urfave/cli/v2"
)

type MigrateCommand struct {
	config ContractConfig.Config
}

func NewMigrateCommand(config ContractConfig.Config) ContractConsole.Command {
	return &MigrateCommand{
		config: config,
	}
}

func (cmd *MigrateCommand) Setup() *cli.Command {
	return &cli.Command{ //nolint:exhaustruct // ignore due to cli configuration
		Name:        "migrate",
		Description: "Run all database migrations",
		Action:      cmd.Handle,
	}
}

func (cmd *MigrateCommand) Handle(*cli.Context) error {
	instance, err := getMigrate(cmd.config)
	if err != nil {
		if errors.Is(err, ErrEmptyMigrationDir) {
			color.Yellowln("There is no migration files yet.")
			return nil
		}

		return err
	}

	if instance == nil {
		color.Yellowln("Database configuration was invalid!")
		return nil
	}

	if err := instance.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			color.Greenln("There is no new migration files.")
			return nil
		}

		return err
	}

	color.Greenln("Migration has been completed.")

	return nil
}
