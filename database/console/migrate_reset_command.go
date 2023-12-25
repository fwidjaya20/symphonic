package console

import (
	"errors"

	ContractConfig "github.com/fwidjaya20/symphonic/contracts/config"
	ContractConsole "github.com/fwidjaya20/symphonic/contracts/console"
	"github.com/golang-migrate/migrate/v4"
	"github.com/gookit/color"
	"github.com/urfave/cli/v2"
)

type MigrateResetCommand struct {
	config ContractConfig.Config
}

func NewMigrateResetCommand(config ContractConfig.Config) ContractConsole.Command {
	return &MigrateResetCommand{
		config: config,
	}
}

func (cmd *MigrateResetCommand) Setup() *cli.Command {
	return &cli.Command{ //nolint:exhaustruct // ignore due to cli configuration
		Name:        "migrate:reset",
		Category:    "migrate",
		Description: "Rollback all database migrations",
		Action:      cmd.Handle,
	}
}

func (cmd *MigrateResetCommand) Handle(*cli.Context) error {
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

	if err = instance.Down(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return err
	}

	instance, err = getSeeder(cmd.config)
	if err != nil {
		if errors.Is(err, ErrEmptyMigrationDir) {
			color.Yellowln("There is no seeder files yet.")
			return nil
		}

		return err
	}

	if instance == nil {
		color.Yellowln("Database configuration was invalid!")
		return nil
	}

	if err := instance.Down(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return err
	}

	color.Greenln("Migration reset has been completed.")

	return nil
}
